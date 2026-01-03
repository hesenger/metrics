# Tasks: Restructure Backend Directory

## Implementation Tasks

- [x] 1. Create backend directory structure
   - Create `backend/` directory at repository root
   - Prepare for file migration

- [x] 2. Move Go code and modules
   - Move `cmd/` to `backend/cmd/`
   - Move `internal/` to `backend/internal/`
   - Move `go.mod` to `backend/go.mod`
   - Move `go.sum` to `backend/go.sum`

- [x] 3. Move database and SQL files
   - Move `migrations/` to `backend/migrations/`
   - Move `sql/` to `backend/sql/`

- [x] 4. Move build tooling and configuration
   - Move `sqlc.yaml` to `backend/sqlc.yaml`
   - Move `.env` to `backend/.env`
   - Move `bin/` to `backend/bin/` (if exists)
   - Keep `Makefile` at root (will be updated in next task)

- [x] 5. Update Makefile targets
   - Rename `run` target to `dev`
   - Update `dev` target to use `cd backend && go run cmd/server/main.go` for development
   - Update `migrate-up` to use `-path backend/migrations`
   - Update `migrate-down` to use `-path backend/migrations`
   - Update `sqlc` target to run in backend/ directory context (use `cd backend && sqlc generate`)
   - Keep `db-up` and `db-down` referencing root-level docker-compose.yml

- [x] 6. Update sqlc.yaml paths
   - Paths already relative to backend/ (sql/queries, migrations, internal/database)
   - No changes needed

- [x] 7. Update .gitignore paths
   - Update `bin/` to `backend/bin/`
   - Update `.env` to `backend/.env`
   - Ensure all backend-specific ignores are updated

- [x] 8. Verify build commands
   - Run `cd backend && go build ./...` to verify compilation
   - Run `make sqlc` from root to verify sqlc generation
   - Run `make dev` from root to verify server starts with go run
   - Verify no import errors or missing files

- [x] 9. Verify database operations
   - Database starts with `make db-up`
   - Migration paths correctly set to `backend/migrations`
   - Database connection works

- [x] 10. Update documentation
   - No README to update
   - Updated openspec/project.md with new directory structure

- [x] 11. Verify end-to-end functionality
   - Started database with `make db-up` from root
   - Started server with `make dev` from root
   - Tested API endpoints (register, login, OAuth)
   - All existing functionality works

## Validation

- `openspec validate restructure-backend-directory --strict` passes
- Backend builds successfully from `backend/` directory
- All Makefile targets work from root directory
- Server starts and responds to requests
- Database migrations apply successfully
- No broken imports or compilation errors

## Dependencies

- Tasks 1-4 can be done sequentially (file moves)
- Tasks 5-6 depend on tasks 1-4 (configuration updates depend on files being moved)
- Task 7 can be done in parallel with tasks 5-6
- Tasks 8-9 depend on tasks 5-6 (verification depends on config being updated)
- Task 10 can be done anytime after task 1
- Task 11 depends on all previous tasks
