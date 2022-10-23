include .env
export

start:
	docker-compose up -d --build

stop:
	docker-compose down

.PHONY: start stop migrateup migratedown sqlc
