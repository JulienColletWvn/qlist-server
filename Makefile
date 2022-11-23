include .env
export

start:
	docker-compose up -d --build

stop:
	docker-compose down

swagger:
	swag init --parseDependency --parseInternal

migrateup:
	migrate -path db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_URL}:5432/${POSTGRES_DB}?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_URL}:5432/${POSTGRES_DB}?sslmode=disable" -verbose down

sqlc:
	sqlc generate


.PHONY: start stop swagger migrateup migratedown sqlc
