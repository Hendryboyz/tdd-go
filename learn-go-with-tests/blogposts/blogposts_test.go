package blogposts

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("always fail file system")
}

func TestBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts, err := NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("blog content with correct order", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
			secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
		)
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := NewPostsFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}

		got := posts[0]
		expected := Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World`,
		}

		assertPost(t, got, expected)
	})

	t.Run("blog content with random order", func(t *testing.T) {
		const (
			firstBody = `Tags: tdd, go
Title: Post 1
Description: Description 1
---
Hello
World`
		)
		fs := fstest.MapFS{
			"hello world.md": {Data: []byte(firstBody)},
		}

		posts, err := NewPostsFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}

		got := posts[0]
		expected := Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World`,
		}

		assertPost(t, got, expected)
	})
}

func assertPost(t *testing.T, got, expected Post) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %+v, want %+v", got, expected)
	}
}
