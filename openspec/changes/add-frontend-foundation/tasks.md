# Implementation Tasks

## 1. Frontend Project Setup
- [x] 1.1 Initialize Bun project in `frontend/` directory
- [x] 1.2 Install Vite, React, TypeScript dependencies
- [x] 1.3 Configure Vite build output to `../backend/web`
- [x] 1.4 Create basic Vite configuration with proxy for `/api` routes
- [x] 1.5 Set up TypeScript configuration with strict mode
- [x] 1.6 Create basic project structure (src/, public/, index.html)

## 2. UI Framework Setup
- [x] 2.1 Install Mantine UI and dependencies
- [x] 2.2 Install React Router
- [x] 2.3 Configure Mantine provider in root component
- [x] 2.4 Set up basic routing structure
- [x] 2.5 Create placeholder home page component

## 3. Backend Static Serving
- [x] 3.1 Create `backend/web/` directory
- [x] 3.2 Add placeholder file `backend/web/embedded.txt` for local development
- [x] 3.3 Add Go embed directive in main.go for web assets
- [x] 3.4 Configure Fiber static middleware to serve from embedded filesystem
- [x] 3.5 Add catch-all route for SPA client-side routing
- [x] 3.6 Ensure API routes remain under `/api` prefix

## 4. Development Tooling
- [x] 4.1 Add `fe-dev` target to root Makefile for Vite dev server
- [x] 4.2 Add `fe-build` target to root Makefile for production build
- [x] 4.3 Add `fe-install` target for installing frontend dependencies
- [x] 4.4 Update existing `dev` target to start database, backend, and frontend concurrently
- [x] 4.5 Update `.gitignore` with frontend build artifacts
- [x] 4.6 Add `backend/web/*` to .gitignore (except embedded.txt)

## 5. Validation
- [x] 5.1 Test local development: frontend dev server with API proxy
- [x] 5.2 Test production build: embedded assets served by backend
- [x] 5.3 Verify SPA routing works (refresh on non-root routes)
- [x] 5.4 Verify API routes still work under `/api`
- [x] 5.5 Run `openspec validate add-frontend-foundation --strict`
