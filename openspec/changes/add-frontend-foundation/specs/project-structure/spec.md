## ADDED Requirements

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

## MODIFIED Requirements

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
