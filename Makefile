# Makefile

dev:
	docker-compose stop
	docker-compose build
	docker-compose up -d
	docker-compose ps

genmodel:
	go run github.com/volatiletech/sqlboiler mysql