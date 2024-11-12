package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostsFromFS(filesSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(filesSystem, ".")

	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(filesSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagSeparator         = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func() (separator string, meta string) {
		scanner.Scan()
		content := scanner.Text()
		if strings.HasPrefix(content, titleSeparator) {
			separator = titleSeparator
		} else if strings.HasPrefix(content, descriptionSeparator) {
			separator = descriptionSeparator
		} else {
			separator = tagSeparator
		}
		meta = strings.TrimPrefix(scanner.Text(), separator)
		return
	}

	post := Post{}
	for range 3 {
		separator, meta := readMetaLine()
		switch separator {
		case titleSeparator:
			post.Title = meta
		case descriptionSeparator:
			post.Description = meta
		case tagSeparator:
			post.Tags = strings.Split(meta, ", ")
		}
	}
	post.Body = readBody(scanner)
	return post, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
