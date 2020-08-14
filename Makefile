.PHONY: deps server test build client prod deploy down migrate

deps:
	go get ./...

build:
	cd client && npm run build

server:
	docker-compose up -d
	docker-compose logs -f

test:
	go test -v ./...

client:
	cd client && npm start

prod:
	docker build -t happydev .

deploy: build
	gcloud app deploy

down:
	docker-compose down

migrate:
	goose -dir migrations postgres "user=postgres password=postgres dbname=postgres sslmode=disable" up
