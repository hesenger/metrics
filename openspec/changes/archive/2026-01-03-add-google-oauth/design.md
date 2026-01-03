# Design: Google OAuth Authentication

## Overview
Implement Google OAuth 2.0 authentication using the Authorization Code Flow. This extends the existing JWT-based authentication system to support social login via Google.

## Architecture

### OAuth Flow
1. Client requests `/api/auth/google` → Backend redirects to Google OAuth consent screen
2. User authorizes on Google → Google redirects to `/api/auth/google/callback` with authorization code
3. Backend exchanges code for access token → Fetches user info from Google
4. Backend creates/finds user → Generates JWT token
5. Backend redirects to frontend with token (or returns JSON for API clients)

### Database Schema Changes
The `users` table requires modifications to support OAuth:
- `password_hash` becomes nullable (OAuth users don't have passwords)
- Add `oauth_provider` (e.g., "google", null for email/password)
- Add `oauth_id` (provider's user ID)
- Add uniqueness constraint on (oauth_provider, oauth_id)

### Configuration
New environment variables:
- `GOOGLE_CLIENT_ID`: OAuth 2.0 client ID from Google Cloud Console
- `GOOGLE_CLIENT_SECRET`: OAuth 2.0 client secret
- `GOOGLE_REDIRECT_URL`: Callback URL (e.g., `http://localhost:7701/api/auth/google/callback`)

### Dependencies
- Add `golang.org/x/oauth2` for OAuth 2.0 flow
- Add `google.golang.org/api/oauth2/v2` for Google user info API

## Trade-offs

### Password Hash Nullable vs Sentinel Value
**Decision**: Make `password_hash` nullable
**Rationale**: Clearer intent, avoids magic values, simplifies queries

### Account Linking vs Separate Accounts
**Decision**: Separate accounts (out of scope for this change)
**Rationale**: Simpler implementation, can be added later if needed

### State Parameter Storage
**Decision**: Use in-memory map with expiration (simple, stateless alternatives complex)
**Rationale**: Good enough for MVP, can migrate to Redis if needed

## Security Considerations
- State parameter validation prevents CSRF attacks
- HTTPS required in production for OAuth redirects
- OAuth tokens are not stored (only used to fetch user info)
- Same JWT token security as existing authentication

## Implementation Notes
- Reuse existing `AuthResponse` and `UserDetails` structures
- Reuse existing `GenerateToken` function for JWT creation
- Add new handler methods `InitiateGoogleOAuth` and `GoogleOAuthCallback`
- Migration required for schema changes
