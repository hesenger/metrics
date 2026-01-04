# frontend-foundation Specification

## Purpose
TBD - created by archiving change add-frontend-foundation. Update Purpose after archive.
## Requirements
### Requirement: Frontend project structure
The system MUST organize frontend code in a dedicated directory with modern React tooling.

#### Scenario: Frontend directory exists at root
**Given** the monorepo structure
**Then** the `frontend/` directory exists at repository root
**And** the `frontend/src/` directory contains application source code
**And** the `frontend/public/` directory contains static assets
**And** the `frontend/package.json` defines dependencies and scripts
**And** the `frontend/tsconfig.json` configures TypeScript compilation
**And** the `frontend/vite.config.ts` configures the build tool

#### Scenario: Node modules excluded from version control
**Given** the frontend directory structure
**Then** the `frontend/node_modules/` directory is in .gitignore
**And** dependencies are installed via `bun install`

### Requirement: Vite build configuration
The system MUST build frontend assets to the backend web directory for embedded serving.

#### Scenario: Build output configured for backend embedding
**Given** the Vite configuration
**Then** the build output directory is set to `../backend/web`
**And** the build creates an index.html entry point
**And** the build bundles and minifies JavaScript and CSS
**And** the build generates production-optimized assets

#### Scenario: Development server with API proxy
**Given** the Vite development server configuration
**Then** the dev server proxies requests matching `/api` to the backend server
**And** the backend server URL is configurable
**And** hot module replacement (HMR) is enabled
**And** the dev server runs on a separate port from backend

### Requirement: React application foundation
The system MUST provide a React application with TypeScript, routing, and UI components.

#### Scenario: React app with TypeScript
**Given** the frontend source directory
**Then** the application is written in TypeScript
**And** TypeScript strict mode is enabled
**And** React components use TypeScript type definitions
**And** the main entry point renders the root React component

#### Scenario: React Router configured
**Given** the React application
**Then** React Router is installed as a dependency
**And** client-side routing is configured in the root component
**And** route definitions exist in the source code
**And** a default home route is defined

#### Scenario: Mantine UI integrated
**Given** the React application
**Then** Mantine UI is installed as a dependency
**And** the MantineProvider wraps the root component
**And** Mantine components are available for use
**And** Mantine's default theme is applied

### Requirement: Backend static file serving
The system MUST serve frontend assets from an embedded filesystem in the Go binary.

#### Scenario: Assets embedded in binary
**Given** the Go server application
**Then** the `backend/web/` directory is embedded using `//go:embed` directive
**And** the embedded filesystem includes all built frontend assets
**And** the embed includes the index.html file
**And** the embed includes bundled JavaScript and CSS files

#### Scenario: Static files served from root
**Given** a request to a non-API route
**When** the route does not start with `/api`
**Then** the server attempts to serve a matching file from embedded assets
**And** the server serves with appropriate content-type headers
**And** the server returns 404 if no matching file exists

#### Scenario: SPA fallback to index.html
**Given** a request to an unknown route
**When** the route does not match any static file
**And** the route does not start with `/api`
**Then** the server responds with index.html
**And** the response status is 200 OK
**And** client-side routing handles the route

#### Scenario: API routes take precedence
**Given** a request to an API endpoint
**When** the route starts with `/api`
**Then** the request is handled by API route handlers
**And** static file middleware is not invoked
**And** the response is JSON (not HTML)

### Requirement: Frontend build automation
The system MUST provide Makefile targets for frontend development and build operations.

#### Scenario: Frontend development target
**Given** the root Makefile
**Then** a `fe-dev` target exists
**And** the target starts the Vite development server
**And** the command runs from the frontend directory

#### Scenario: Frontend build target
**Given** the root Makefile
**Then** a `fe-build` target exists
**And** the target builds production assets
**And** the output is placed in `backend/web/`
**And** the command runs from the frontend directory

#### Scenario: Frontend install target
**Given** the root Makefile
**Then** a `fe-install` target exists
**And** the target runs `bun install` in frontend directory
**And** dependencies are installed from package.json

