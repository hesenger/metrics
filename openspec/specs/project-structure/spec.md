# project-structure Specification

## Purpose
TBD - created by archiving change restructure-backend-directory. Update Purpose after archive.
## Requirements
### Requirement: Backend code isolation
The system MUST organize all backend code within a dedicated backend directory.

#### Scenario: Backend code in backend folder
**Given** the project repository structure
**Then** the `backend/` directory exists at repository root
**And** the `backend/cmd/` directory contains the server entry point
**And** the `backend/internal/` directory contains internal Go packages
**And** the `backend/migrations/` directory contains database migrations
**And** the `backend/sql/` directory contains SQL query files
**And** the `backend/go.mod` file defines the Go module

#### Scenario: Backend build tooling configured
**Given** the backend directory structure
**Then** the `backend/sqlc.yaml` exists with sqlc configuration
**And** the `backend/.env` exists for local environment variables
**And** all paths in sqlc.yaml are relative to backend directory

#### Scenario: Backend binaries in backend folder
**Given** a successful backend build
**Then** compiled binaries are placed in `backend/bin/`
**And** the bin directory is not committed to version control

### Requirement: Shared project files at root
The system MUST keep project-wide shared files at the repository root.

#### Scenario: Shared files remain at root
**Given** the project repository structure
**Then** the `Makefile` file exists at repository root
**And** the `docker-compose.yml` file exists at repository root
**And** the `openspec/` directory exists at repository root
**And** the `.gitignore` file exists at repository root
**And** project documentation files (CLAUDE.md, AGENTS.md) exist at repository root

#### Scenario: Makefile orchestrates monorepo
**Given** the Makefile at repository root
**Then** the Makefile contains targets for backend operations
**And** the Makefile contains targets for frontend operations
**And** backend targets reference `backend/` directory in their paths
**And** frontend targets reference `frontend/` directory in their paths
**And** all make commands run from repository root

#### Scenario: Database configuration is shared
**Given** the docker-compose.yml at repository root
**Then** database services are accessible from both backend and frontend
**And** the DATABASE_URL environment variable can be used by backend services

### Requirement: Go module path consistency
The system MUST maintain consistent Go import paths after restructuring.

#### Scenario: Import paths remain unchanged
**Given** the backend code moved to backend/ directory
**And** the Go module remains `github.com/hesen/metrics`
**Then** existing import paths like `github.com/hesen/metrics/internal/handlers` continue to work
**And** all Go packages resolve correctly
**And** no code changes are required for imports

#### Scenario: Go module in backend directory
**Given** the backend directory structure
**Then** the `backend/go.mod` file defines module `github.com/hesen/metrics`
**And** Go commands run from backend/ directory find the correct module
**And** dependencies are managed within backend/go.sum

### Requirement: Frontend code isolation
The system MUST organize all frontend code within a dedicated frontend directory.

#### Scenario: Frontend code in frontend folder
**Given** the project repository structure
**Then** the `frontend/` directory exists at repository root
**And** the `frontend/src/` directory contains React application code
**And** the `frontend/public/` directory contains static assets
**And** the `frontend/package.json` file defines Node.js dependencies
**And** the `frontend/tsconfig.json` file configures TypeScript
**And** the `frontend/vite.config.ts` file configures Vite

#### Scenario: Frontend build tooling configured
**Given** the frontend directory structure
**Then** Vite is configured to build to `../backend/web`
**And** the package.json includes build and dev scripts
**And** Bun is the package manager (bun.lockb exists)

#### Scenario: Frontend build output in backend folder
**Given** a successful frontend build
**Then** compiled assets are placed in `backend/web/`
**And** the web directory contains index.html as entry point
**And** the web directory is not committed to version control
**And** the web directory is embedded in the Go binary

