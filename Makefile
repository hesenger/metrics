.PHONY: dev be-dev fe-dev fe-build fe-install migrate-up migrate-down sqlc db-up db-down

dev:
	@echo "Starting database, backend, and frontend..."
	@make db-up
	@(set -a && source .env && cd backend && go run cmd/server/main.go) & \
	(cd frontend && ~/.bun/bin/bun run dev) & \
	wait

be-dev:
	cd backend && go run cmd/server/main.go

fe-dev:
	cd frontend && ~/.bun/bin/bun run dev

fe-build:
	cd frontend && ~/.bun/bin/bun run build

fe-install:
	cd frontend && ~/.bun/bin/bun install

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
