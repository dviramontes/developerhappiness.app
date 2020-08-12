.PHONY: deps server test build client prod deploy

deps:
	go get -u ./...

build: deps
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
