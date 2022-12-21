include .env
export

start:
	docker-compose up -d --build

stop:
	docker-compose down

run:
	go run main.go

swagger:
	swag init --parseDependency --parseInternal

migrateup:
	migrate -path db/migration -database "${POSTGRES_URL}" -verbose up

migratedown:
	migrate -path db/migration -database "${POSTGRES_URL}" -verbose down

sqlc:
	sqlc generate


.PHONY: start stop swagger migrateup migratedown sqlc run
