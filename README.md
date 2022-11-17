# tdd-go

A repository to practice TDD(Test Driven Development) with GoLang

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

# view the document in http://localhost:6060/pkg/
$ godoc -http=:6060
```

## Dependencies Modules

> _errcheck_: `go install github.com/kisielk/errcheck@latest`

```
errcheck .
```
