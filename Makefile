# Makefile

dev:
	docker-compose stop
	docker-compose build
	docker-compose up -d
	docker-compose ps

genmodel:
	go run github.com/volatiletech/sqlboiler mysql

generate:
	go generate ./
	go generate ./...

fmt:
	go fmt ./...

test:
	go test ./...
