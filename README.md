# `developerhappiness.app`

> A web application that syncs users from slack workspace and displays them as a list

Project layout based on [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

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
- (optional) [goose for migrations](https://github.com/pressly/goose)
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

With goose installed, run: `go get -u github.com/pressly/goose/cmd/goose`

#### create new migration

`goose -d migrations create add_some_column sql`

#### run migrations

`make migrate`

### TODO

With more time and bandwidth, this is a list of things I would like to add:

* [ ] Add separate configs for local, test and prod environments
* [ ] Abstract the API layer into an interaction and a repository so that we can easily mock database interactions without needing an actual DB connection (aka in-memory).
* [ ] Implement a graphql subscription for subscribing to user updates (slack API <> golang server <> event broker <> gql client)
* [ ] CI/CD

