# gobank
Learning Go by creating a banking API which handles accounts, transfering money with auth.
Following a guide by Anthony GG.

### What do I want to get out of this
A good understanding of Go, how to create stable APIs which can be handled by a seperate frontend and implement authentication with JWT.

### env
This product uses an .env file to hold the database connections details

Currently using a Neon.tech connection string using YAML format

eg: `DB_CON: postgres://daniel:<password>@ep-mute-rain-952417.us-east-2.aws.neon.tech/neondb`

### Makefile
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
