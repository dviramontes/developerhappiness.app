# `developerhappiness.app`

> A web application that syncs users from slack workspace and displays them as a list

Project layout is based on [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

### Project Layout

```
{root}

├── client   # React FE Client
├── cmd      # Main driver* for Golang app
├── pkg      # Library code to be used by external applications
    └── event
        └──  # Slack Event API Interactor*
...
```

### Requirements

- [Node + NPM](https://nodejs.org/en/) 
- [golang](https://golang.org/)
- [docker](https://docs.docker.com/get-docker/)
- [docker-compose](https://docs.docker.com/compose/install/)
- (optional) [google cloud sdk for deployment](https://cloud.google.com/sdk/docs/downloads-versioned-archives)

### Setup

- `cd client && npm ci`
- `cp sample.config.yaml config.yaml`

### Development

- `make server`
- `make client`

### Tests

- `make test` # runs server tests

### Deploy

- `make deploy` # deploys app to google cloud engine

### Terms
- *driver*: TODO: define
- *interactor*: TODO: define

