# tdd-go

A repository to practice TDD(Test Driven Development) with GoLang. This practice follow the steps from [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/) website.

## Command

```bash
# run test
$ go test

# run test and examples
$ go test -v # show the testing details

# run benchmark
go test -bench="."

# get test coverage
go test -cover

# check race condition
go test -race

# view the document in http://localhost:6060/pkg/
$ godoc -http=:6060

# format current directory
$ gofmt -w .

# examines Go source code and reports suspicious constructs
$ go vet
```

## Dependencies Modules

> _errcheck_: `go install github.com/kisielk/errcheck@latest`

```
errcheck .
```
