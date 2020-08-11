# `developerhappiness.app`

> A web application that syncs users from slack workspace and displays them as a list

### Project Layout

```
{root}

├── client   # React FE Client
├── cmd      # Main driver for Golang app
├── pkg      # Library code to be used by external applications
    └── event
        └──  # Slack Event API Interactor
...
```

### Requirements

- [Node + NPM](https://nodejs.org/en/) 
- [golang](https://golang.org/)
- [docker](https://docs.docker.com/get-docker/)
- [docker-compose](https://docs.docker.com/compose/install/)

### Setup

- `cd client && npm ci`
- `cp sample.config.yml config.yml`

### Running

- `make server`
- `make client`

### Tests

- `make test` # runs server tests

### Terms Used
- *Clean Architecture*: TODO: define
- *driver*: TODO: define
- *interactor*: In an MVC style app, this would be your controller. This Go is not an OOP language this is loosely one of the many styles
in which a Go app can be organized. 

