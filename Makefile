.PHONY: deps server test build client prod deploy

deps:
	go get -u ./...

build: deps
	docker-compose up --build

run:
	docker-compose up -d
	docker-compose logs -f

test:
	go test -v ./...

client:
	cd client && npm run build

prod:
	docker build -t happydev .

deploy: client
	gcloud app deploy
