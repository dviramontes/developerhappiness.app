.PHONY: deps server test build client prod

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

prod:
	docker build -t happydev .
