# Change: Add Frontend Authentication

## Why
Enable users to register, log in, and authenticate via Google OAuth through a web interface. Currently, authentication endpoints exist but there is no UI for users to interact with them. Additionally, JWT tokens are returned in response bodies which prevents HttpOnly cookie usage for improved security.

## What Changes
1. Update backend authentication endpoints to set JWT in HttpOnly strict cookie
2. Update backend to return only user data (id, email, oauth_provider) in response body
3. Add React Query to frontend for server state management
4. Create authentication context provider for global auth state
5. Implement login page at `/login` with email/password form and Google SSO button
6. Implement registration page at `/register` with email/password form
7. Add protected route wrapper component for authentication guards
8. Implement logout functionality that clears cookie and local state
9. Add auto-redirect logic for authenticated users visiting auth pages
10. Store user data in localStorage for client-side auth state persistence
11. JWT cookie expires after 7 days
12. Returning users will check authentication through a new endpoint api/auth/me
13. api/auth/me refreshes JWT cookie on request, restarting expiration timer
14. FE is unaware of JWT expiration, but checking the response from ME endpoint knows its state
15. ME endpoint returns all data needed to be stored in the FE (id, email, oauth_provider)

## Impact
- Affected specs: user-authentication, user-registration, google-authentication, frontend-foundation
- Affected code: backend/cmd/server/main.go, backend auth handlers, frontend routing
- New capability: frontend-authentication
- Breaking change: JWT token now in cookie instead of response body (clients must handle cookies)
