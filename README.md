# `developerhappiness.app`

> A web application that syncs users from slack workspace and displays them as a list

Project layout is based on [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

### Project Layout

```
{root}

├── client   # React FE Client
├── internal # Private application and library code
├── cmd      # Main driver for Golang app (our entry point)
├── pkg      # Public library code to be used by external applications
    └── api  # Slack event parsing business logic
    └── db   # Database connection and models (lightweight schemas provided by gorm)
...
```

### Requirements

- [Node + NPM](https://nodejs.org/en/) 
- [golang](https://golang.org/)
- [docker](https://docs.docker.com/get-docker/)
- [docker-compose](https://docs.docker.com/compose/install/)
- [goose for migrations](https://github.com/pressly/goose)
- (optional) [google cloud sdk for deployment](https://cloud.google.com/sdk/docs/downloads-versioned-archives)

### Setup

- `cd client && npm ci`
- `cp sample.config.yaml config.yaml`

### Development

- `make server`
- `make client`

### Tests

- `make test` # runs unit tests
- `make test-ci` # runs integration tests

### Deploy

- `make deploy` # deploys app to google cloud engine

### Migrations

with goose installed `go get -u github.com/pressly/goose/cmd/goose`

#### create new migration

`goose -d migrations create add_some_column sql`

#### run migrations

`make migrate`

### TODO

* [x] create a separate database for integration tests
* [ ] separate configs for local, test and prod environments
