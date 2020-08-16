.PHONY: deps server test test-ci build client prod deploy down migrate

deps:
	go get ./...

build:
	cd client && npm run build

server:
	docker-compose up -d
	docker-compose logs -f

test:
	go test -v ./...

test-ci:
	docker-compose -f ./docker-compose.test.yaml build
	docker-compose -f ./docker-compose.test.yaml up -d
	go test -v --tags=integration ./...
	docker-compose -f ./docker-compose.test.yaml down

client:
	cd client && npm start

prod:
	docker build -t happydev .

deploy: build
	gcloud app deploy

down:
	docker-compose down

migrate:
	goose -dir migrations postgres "user=postgres password=postgres dbname=happydev sslmode=disable" up
