# gobank
Learning Go by creating a banking API which handles accounts, transfering money with auth.

Following a guide by Anthony GG.

## What do I want to get out of this
A good understanding of Go, how to create stable APIs which can be handled by a seperate frontend and implement authentication with JWT.

## Database
This branch was changed from Postgres to sqlite3 database named "gobank.db".

The package glebarez/go-sqlite is used which is CGo free.

## Makefile
This project uses a basic makefile for now.

Run the project with `make run`

```
build:
  @go build -o bin/gobank

run: build
  @./bin/gobank

test:
  @go test -v ./..
```
