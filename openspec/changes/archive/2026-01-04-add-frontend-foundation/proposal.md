# Change: Add Frontend Foundation

## Why
Enable user interface development by establishing the frontend foundation with modern React tooling. Currently, the backend provides API-only endpoints with no way for users to interact with the system through a web interface.

## What Changes
1. Initialize Vite + React + TypeScript project in `frontend/` directory
2. Configure Bun as package manager and runtime
3. Set up Mantine UI component library
4. Configure React Router for client-side routing
5. Configure Vite build to output to `backend/web/` for embedded serving
6. Add backend static file serving with Go embed
7. Configure local development proxy for API calls to backend on port 7701
8. Configure frontend dev server to run on port 7702
9. Extend Makefile with frontend build and development targets
10. `make dev` starts backend and frontend apps, making sure DB docker is running
11. Add simple file to backend/web/embedded.txt so Fiber serves it on local development

## Impact
- Affected specs: project-structure
- Affected code: backend/cmd/server/main.go, Makefile, .gitignore
- New directories: frontend/
- New capability: frontend-foundation
