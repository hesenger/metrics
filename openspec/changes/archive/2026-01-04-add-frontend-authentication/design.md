# Design: Frontend Authentication

## Architecture Overview

This change implements the frontend authentication UI that connects to existing backend endpoints, with a security improvement to use HttpOnly cookies for JWT storage.

### Authentication Flow

**Email/Password Registration:**
1. User fills form on `/register` page
2. Frontend validates input (email format, password length)
3. React Query mutation POSTs to `/api/auth/register`
4. Backend creates user, sets JWT in HttpOnly cookie, returns user data
5. Frontend stores user data in localStorage
6. Auth context updates with authenticated state
7. User redirected to home/dashboard

**Email/Password Login:**
1. User fills form on `/login` page
2. React Query mutation POSTs to `/api/auth/login`
3. Backend verifies credentials, sets JWT in HttpOnly cookie, returns user data
4. Frontend stores user data in localStorage
5. Auth context updates with authenticated state
6. User redirected to home/dashboard

**Google OAuth:**
1. User clicks "Sign in with Google" button on `/login` or `/register`
2. Frontend redirects to `/api/auth/google`
3. Backend redirects to Google OAuth consent screen
4. User completes OAuth, Google redirects to `/api/auth/google/callback`
5. Backend sets JWT in HttpOnly cookie, redirects to frontend with user data in URL params
6. Frontend extracts user data from URL, stores in localStorage
7. Auth context updates with authenticated state

**Logout:**
1. User clicks logout button
2. React Query mutation POSTs to new `/api/auth/logout` endpoint
3. Backend clears JWT cookie
4. Frontend clears localStorage
5. Auth context updates to unauthenticated
6. User redirected to `/login`

**Returning User Authentication Check:**
1. User opens app or refreshes page
2. Frontend checks localStorage for user data
3. If user data exists, auth context initializes as authenticated
4. Auth context makes GET request to `/api/auth/me` on mount
5. Backend validates JWT from cookie, refreshes cookie expiration to 7 days
6. Backend returns user data (id, email, oauth_provider)
7. Frontend updates localStorage and auth context with fresh user data
8. If `/me` returns 401, frontend clears localStorage and redirects to `/login`
9. This flow keeps users logged in and extends session with activity

### State Management

**Auth Context:**
- Provides global auth state: `{ user: User | null, isAuthenticated: boolean }`
- Loads initial state from localStorage on app mount
- Calls `/api/auth/me` on mount to verify and refresh authentication
- Updates when login/register/logout mutations succeed or /me response received
- Wraps entire app in `main.tsx`

**React Query:**
- Manages server state and mutations
- Mutations for: login, register, logout
- Query for: /me (verify and refresh authentication)
- Automatic error handling
- Loading states for UI feedback

**Protected Routes:**
- Wrapper component checks `isAuthenticated` from context
- Redirects to `/login` if not authenticated
- Preserves intended destination for post-login redirect

### Security Considerations

**HttpOnly Cookie for JWT:**
- Prevents XSS attacks from accessing token
- Cookie flags: HttpOnly, Secure (in production), SameSite=Strict
- Backend sets cookie on all successful auth responses
- Frontend doesn't need to manually handle JWT

**User Data in localStorage:**
- Not sensitive (id, email, provider)
- Enables client-side auth state persistence
- Cleared on logout
- Synced with cookie state (no token = no user data)

**CSRF Protection:**
- SameSite=Strict cookie prevents CSRF
- No additional CSRF token needed for this initial implementation

### Component Structure

```
frontend/src/
  contexts/
    auth-context.tsx         # Auth provider and hooks
  pages/
    login.tsx                # Login page with form + Google button
    register.tsx             # Register page with form
    auth-callback.tsx        # OAuth callback handler
    home.tsx                 # Existing home (will be protected)
  components/
    protected-route.tsx      # Route guard wrapper
  api/
    auth.ts                  # React Query mutations and types
  types/
    user.ts                  # User type definition
```

### Backend Changes Required

**Cookie Management:**
- Fiber cookie utilities for setting HttpOnly cookies
- JWT secret from environment (already exists)
- Cookie expiration matches JWT expiration (7 days)
- Cookie is refreshed on every `/api/auth/me` request, extending session

**Response Format Changes:**
- Register: Returns `{ user: { id, email, oauth_provider, created_at } }` instead of `{ token, user }`
- Login: Returns `{ user: { id, email, oauth_provider, created_at } }` instead of `{ token, user }`
- Google callback: Redirects to frontend with user data in query params

**New Endpoints:**
- POST `/api/auth/logout` - clears cookie, returns 204
- GET `/api/auth/me` - validates JWT from cookie, refreshes cookie expiration, returns user data

### Frontend Libraries

**New Dependencies:**
- `@tanstack/react-query` - server state management
- No additional dependencies needed (Mantine has form components)

### Route Structure

```
Public routes:
  /login          - Login page
  /register       - Register page
  /auth/callback  - OAuth callback handler

Protected routes:
  /               - Home/dashboard (existing)
  /* future routes
```

### Error Handling

**Frontend:**
- Display validation errors inline on forms
- Show API errors via Mantine notifications
- Network errors handled by React Query with retry logic

**Backend:**
- Existing error responses remain unchanged
- Cookie setting failures logged but don't block response

### Testing Approach

**Frontend:**
- No component tests (per project conventions)
- Complex auth logic in hooks can be tested if needed

**Backend:**
- Update existing auth handler tests to verify cookie setting
- Test cookie flags and expiration (7 days)
- Test logout endpoint clears cookie
- Test /me endpoint validates JWT and refreshes cookie
- Test /me endpoint returns 401 for invalid/expired JWT

## Trade-offs

**HttpOnly Cookie vs localStorage JWT:**
- ✅ More secure (XSS-proof)
- ✅ Browser handles cookie lifecycle
- ❌ Requires backend changes
- ❌ CORS complexity for cross-domain (not relevant for this monorepo)

**React Query vs Custom Fetch:**
- ✅ Built-in loading/error states
- ✅ Mutation handling
- ✅ Industry standard
- ❌ Additional dependency

**Separate Login/Register Pages vs Single Page:**
- ✅ Simpler component logic
- ✅ Clearer URL structure
- ✅ Easier to implement different forms
- ❌ Extra route/page to maintain

**User Data in localStorage:**
- ✅ Enables instant auth state on initial render
- ✅ Kept fresh via /me endpoint on app mount
- ✅ Reduces layout shift during auth check
- ❌ Can become stale if updated on server between sessions
- ❌ Not synced across tabs (acceptable for v1)

**/me Endpoint with Cookie Refresh:**
- ✅ Extends session automatically with activity
- ✅ 7-day expiration provides good UX (stay logged in)
- ✅ No frontend JWT expiration logic needed
- ❌ Extra API call on every app mount
- ❌ Slightly more complex backend logic

## Open Questions

None - all requirements clarified with user.
