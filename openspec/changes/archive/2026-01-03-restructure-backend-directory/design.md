# Design: Restructure Backend Directory

## Overview
Reorganize the monorepo to clearly separate backend code from future frontend code and shared project files.

## Target Directory Structure

```
/
├── backend/
│   ├── cmd/
│   ├── internal/
│   ├── migrations/
│   ├── sql/
│   ├── bin/
│   ├── go.mod
│   ├── go.sum
│   ├── sqlc.yaml
│   └── .env
├── Makefile
├── docker-compose.yml
├── .gitignore
├── openspec/
├── CLAUDE.md
└── AGENTS.md
```

## Design Decisions

### What Moves to backend/
**Backend-specific code and tooling:**
- `cmd/` - Go server entry point
- `internal/` - Go packages
- `migrations/` - Database migrations
- `sql/` - SQL query files
- `bin/` - Built binaries
- `go.mod`, `go.sum` - Go module files
- `sqlc.yaml` - sqlc configuration
- `.env` - backend environment variables

**Rationale:** These are all backend-specific and will be independent from frontend tooling.

### What Stays at Root
**Shared project files:**
- `Makefile` - monorepo build orchestration (coordinates backend and future frontend)
- `docker-compose.yml` - database services (shared by both BE and FE during dev)
- `openspec/` - project specifications (project-wide)
- `.gitignore` - repository-wide ignore rules
- `CLAUDE.md`, `AGENTS.md` - project-wide documentation

**Rationale:** These files are either shared between frontend/backend or are project-level concerns. The Makefile serves as the main entry point for all build commands across the monorepo.

### Module Path Strategy
**Decision:** Keep module path as `github.com/hesen/metrics` but move it to `backend/go.mod`

**Alternative considered:** Change module path to `github.com/hesen/metrics/backend`
- Rejected because it would require updating all import paths throughout the codebase
- Current approach is simpler and less error-prone

**Rationale:** Import paths like `github.com/hesen/metrics/internal/handlers` will continue to work because Go looks for go.mod starting from the current directory and walks up the tree.

### Makefile Strategy
**Decision:** Keep Makefile at root and update paths to reference backend/ directory

**Alternative considered:** Move Makefile to backend/ directory
- Rejected because it requires developers to cd into backend/ for all commands
- Makes it harder to add frontend commands later

**Rationale:** Root Makefile serves as the central orchestration point for the entire monorepo. Backend commands will be prefixed or use -C flag to run in backend context. When frontend is added, we can add frontend-specific targets to the same Makefile (e.g., `make fe-dev`, `make be-dev`).

### Development Command Strategy
**Decision:** Rename `make run` to `make dev` and use `go run` instead of building binaries

**Rationale:**
- `dev` is more intuitive for development workflow
- `go run` compiles and runs in one step, better for rapid iteration
- No need to track/clean built binaries during development
- Aligns with common monorepo patterns where `dev` starts development servers

## Migration Strategy

1. Create `backend/` directory
2. Move all backend files/folders to `backend/`
3. Update paths in configuration files (sqlc.yaml, Makefile)
4. Update .gitignore paths if needed
5. Verify build and run commands work
6. Update documentation if needed

## Trade-offs

### Pros
- Clear separation of concerns
- Easier to add frontend later
- Standard monorepo structure
- Independent tooling for BE/FE

### Cons
- Requires path updates in several config files
- Need to run commands from backend/ directory
- Slightly longer paths in development

## Testing Approach
- Verify `make dev` works from repository root
- Verify `make migrate-up` works from repository root
- Verify `make sqlc` works from repository root
- Verify server starts and responds to requests using `go run`
- Verify all paths resolve correctly
