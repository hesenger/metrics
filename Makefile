.PHONY: run migrate-up migrate-down sqlc db-up db-down

run:
	go run cmd/server/main.go

db-up:
	docker compose up -d

db-down:
	docker compose down

migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down

sqlc:
	sqlc generate
