# Proposal: Restructure Backend Directory

## Summary
Move all backend code into a dedicated `backend/` folder to prepare the monorepo for future frontend code and improve project organization.

## Problem
Currently, backend code (Go modules, migrations, SQL queries) lives at the repository root, making it unclear where frontend code will live when added. This flat structure makes it harder to manage build tooling and dependencies separately for backend and frontend.

## Solution
Restructure the monorepo by:
- Moving all backend-related code and configuration into `backend/` folder
- Keeping shared/root-level files at repository root (docker-compose.yml, openspec/)
- Updating build tooling paths (Makefile, sqlc.yaml)
- Maintaining all import paths and functionality unchanged

## Scope
This change introduces a new project-structure capability that defines the monorepo organization.

## Out of Scope
- Frontend implementation or scaffolding
- Changes to actual backend code logic
- Changes to API contracts or database schema

## Related Changes
- Prepares foundation for future frontend integration
- Affects all build and development tooling

## Success Criteria
- All backend code resides in `backend/` folder
- `make dev`, `make migrate-up`, `make sqlc` commands work correctly
- All Go imports resolve correctly
- Application runs in development mode without errors
- Tests pass (when implemented)
