.PHONY: deps server test build client

deps:
	go get -u ./...

build: deps
	docker-compose up --build

server:
	docker-compose up

test:
	go test -v ./...

client:
	cd client && npm start
