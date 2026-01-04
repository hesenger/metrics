# Design: Frontend Foundation

## Context
The metrics platform needs a web interface for users to interact with analytics and feature flags. The backend currently only provides API endpoints. The project follows a monorepo structure with backend in `backend/` directory.

**Constraints:**
- Single binary deployment (frontend assets embedded in Go binary)
- Local development must support hot reload for frontend
- API endpoints already exist under `/api` prefix
- Tech stack predefined: Bun, Vite, React, TypeScript, Mantine UI, React Router

**Stakeholders:** Developers building UI features, end users accessing the platform

## Goals / Non-Goals

**Goals:**
- Set up modern React development environment with TypeScript
- Enable hot reload during local development
- Build static assets that embed into Go binary for production
- Maintain clear separation between frontend and backend code
- Support client-side routing (SPA) without breaking API routes

**Non-Goals:**
- Server-side rendering (SSR)
- Multiple frontends or micro-frontends
- Backend framework changes beyond static serving
- Authentication UI (will be separate change)

## Decisions

### Frontend Directory Structure
**Decision:** Create `frontend/` directory at repository root, parallel to `backend/`

**Rationale:**
- Maintains monorepo symmetry (backend/, frontend/)
- Clear ownership of dependencies and tooling
- Follows existing pattern from backend restructuring
- Easy to locate all frontend code

**Alternatives considered:**
- `web/` or `ui/` directory: Less explicit, could be confused with web assets
- Inside `backend/`: Violates separation of concerns

### Build Output Location
**Decision:** Vite builds to `backend/web/` directory

**Rationale:**
- Backend embeds assets using `//go:embed` directive
- Keeps embedded assets close to server code
- Clear that backend/web/ is generated, not source
- Simple relative path from frontend/

**Alternatives considered:**
- `dist/` at root: Would require copying to backend or changing embed path
- `backend/static/`: "web" better indicates SPA nature

### Go Embed Strategy
**Decision:** Use `//go:embed all:web` with io/fs package and Fiber's filesystem middleware

**Rationale:**
- Compiles assets directly into binary (single deployment artifact)
- No runtime file I/O needed for serving assets
- Standard Go 1.16+ feature, no external dependencies
- Fiber has built-in support for embedded filesystems

**Pattern:**
```go
//go:embed all:web
var webAssets embed.FS

// Serve from subdirectory to strip "web/" prefix
app.Use("/", filesystem.New(filesystem.Config{
    Root: http.FS(webAssets),
    PathPrefix: "web",
}))
```

### Local Development Workflow
**Decision:** Run Vite dev server separately, proxy API requests to backend, unified `make dev` command

**Rationale:**
- Preserves Vite's hot module replacement (HMR)
- No need to rebuild/restart backend for frontend changes
- Standard Vite development pattern
- Backend runs independently on configured port
- Single command (`make dev`) starts all services: database, backend, and frontend

**Configuration:**
```typescript
// vite.config.ts
export default {
  server: {
    proxy: {
      '/api': 'http://localhost:7701'
    }
  }
}
```

**Alternatives considered:**
- Separate commands for each service: Less developer-friendly, easy to forget starting a service
- Docker Compose for entire stack: Overkill for local development, slower iteration

### Routing Strategy
**Decision:** Client-side routing with catch-all fallback to index.html

**Rationale:**
- SPA pattern requires all routes to serve index.html
- API routes under `/api` are explicitly handled first
- Allows bookmarking and direct navigation to frontend routes
- Standard pattern for React Router

**Implementation order:**
1. API routes registered with `/api` prefix (already done)
2. Static files served from embedded FS at `/`
3. Catch-all `/*` → index.html for client-side routing

### Package Manager
**Decision:** Use Bun for frontend dependency management and runtime

**Rationale:**
- Project specification requires Bun
- Faster than npm/yarn for installs and execution
- Built-in TypeScript support
- Compatible with npm registry and package.json

### UI Component Library
**Decision:** Mantine UI v7

**Rationale:**
- Project specification requires Mantine
- Comprehensive component library with hooks
- Built-in dark mode support
- TypeScript-first design
- Good documentation

### Placeholder File for Local Development
**Decision:** Add `backend/web/embedded.txt` placeholder file, excluded from .gitignore

**Rationale:**
- Go embed directive requires directory to exist with at least one file
- Without placeholder, backend fails to compile during initial development
- Simple text file serves embedded assets during local dev before frontend is built
- Committed to version control so fresh clones work immediately
- Removed from .gitignore pattern (`backend/web/*` ignores built assets but not embedded.txt)

**Alternatives considered:**
- .gitkeep file: Not readable by embed, would still cause compile error
- Require frontend build before backend runs: Poor developer experience, circular dependency
- Conditional embed with build tags: Adds unnecessary complexity

## Risks / Trade-offs

### Risk: Embedded assets increase binary size
**Mitigation:**
- Vite tree-shaking and minification reduces bundle size
- Gzip/Brotli compression in production
- Monitor binary size over time
- Consider asset CDN if size becomes problematic

### Risk: Separate dev server increases complexity
**Mitigation:**
- Document workflow clearly in README
- Provide Makefile targets for common operations
- Single `make dev` starts all services (database, backend, frontend) concurrently
- Individual targets (`fe-dev`, `dev` for backend-only) available when needed

### Trade-off: Single binary vs separate frontend serving
**Chosen:** Single binary (embed)
**Given up:** Ability to update frontend without backend redeployment
**Rationale:** Simplifies deployment, hosting, and version consistency. Updates typically involve both frontend and backend.

## Migration Plan

**No migration needed** - this is a greenfield frontend setup. Existing API routes remain unchanged.

**Rollout:**
1. Merge frontend foundation (no breaking changes, API unaffected)
2. Backend continues serving API as before
3. Frontend served from `/` when accessed via browser
4. Subsequent changes will add actual UI pages

**Rollback:**
- Remove static serving middleware from main.go
- API-only mode restored immediately
- No database or data changes involved

## Open Questions

None - all technical decisions align with project specifications.
