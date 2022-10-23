include .env
export

start:
	docker-compose up -d --build

stop:
	docker-compose down

swagger:
	swag init --parseDependency --parseInternal

.PHONY: start stop swagger
