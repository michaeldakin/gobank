# gobank
Bank API in Go

## Makefile
This project uses a basic makefile for now.
```
build:
	@go build -o bin/gobank

run: build
	@./bin/gobank

test:
	@go test -v ./..
```
