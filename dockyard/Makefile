include .env

.PHONY: migrate-up migrate-down sqlc-generate docker-up docker-down build-api

build-api:
	cd api && go mod tidy

migrate-up:
	migrate -path api/internal/db/migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path api/internal/db/migrations -database "$(DATABASE_URL)" down 1

sqlc-generate:
	cd api && sqlc generate

docker-up:
	docker compose up --build

docker-down:
	docker compose down
