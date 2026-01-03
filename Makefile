.PHONY: dev migrate-up migrate-down sqlc db-up db-down

dev:
	cd backend && go run cmd/server/main.go

db-up:
	docker compose up -d

db-down:
	docker compose down

migrate-up:
	migrate -path backend/migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path backend/migrations -database "$(DATABASE_URL)" down

sqlc:
	cd backend && sqlc generate
