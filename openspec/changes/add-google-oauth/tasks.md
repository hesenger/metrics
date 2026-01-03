# Tasks: Add Google OAuth Authentication

## Implementation Tasks

- [x] 1. Add Google OAuth dependencies to go.mod
   - Add `golang.org/x/oauth2`
   - Add `google.golang.org/api/oauth2/v2`
   - Run `go mod tidy`

- [x] 2. Create database migration for OAuth fields
   - Create migration file `000002_add_oauth_fields.up.sql`
   - Add `oauth_provider` column (nullable VARCHAR)
   - Add `oauth_id` column (nullable VARCHAR)
   - Make `password_hash` nullable
   - Add unique constraint on (oauth_provider, oauth_id)
   - Create corresponding down migration

- [x] 3. Update sqlc queries for OAuth users
   - Add `CreateOAuthUser` query with oauth_provider and oauth_id
   - Add `GetUserByOAuthProvider` query
   - Update `User` struct in generated code to include new fields

- [x] 4. Extend Config struct for Google OAuth
   - Add `GoogleClientID` field
   - Add `GoogleClientSecret` field
   - Add `GoogleRedirectURL` field
   - Load from environment variables in `config.Load()`
   - Return error if Google OAuth env vars are missing

- [x] 5. Implement OAuth state management
   - Create `internal/auth/oauth-state.go`
   - Implement in-memory state store with expiration (5 minutes)
   - Add `GenerateState()` and `ValidateState()` functions
   - Include cleanup for expired states

- [x] 6. Implement Google OAuth handlers
   - Create `InitiateGoogleOAuth` handler in `internal/handlers/auth.go`
   - Generate and store state parameter
   - Redirect to Google OAuth consent screen with proper parameters
   - Create `GoogleOAuthCallback` handler
   - Validate state parameter
   - Exchange authorization code for access token
   - Fetch user info from Google
   - Create or find user in database
   - Generate JWT token and return response

- [x] 7. Register OAuth routes
   - Add `GET /api/auth/google` route in `cmd/server/main.go`
   - Add `GET /api/auth/google/callback` route
   - Pass Google OAuth config to AuthHandler

- [x] 8. Run and test database migration
   - Apply migration with `make migrate-up` or equivalent
   - Verify schema changes in database
   - Test rollback with down migration

- [x] 9. Manual end-to-end testing
   - Start application with Google OAuth credentials
   - Test OAuth flow with real Google account
   - Verify new user creation via OAuth
   - Verify existing user login via OAuth
   - Verify JWT token works for protected endpoints
   - Verify email/password auth still works

- [x] 10. Update environment configuration examples
    - Add Google OAuth env vars to `.env` file
    - Document where to obtain Google OAuth credentials (Google Cloud Console)

## Validation

- `openspec validate add-google-oauth --strict` passes
- Database migration applies and rolls back cleanly
- OAuth flow completes successfully for new users
- OAuth flow completes successfully for existing users
- JWT tokens from OAuth match structure of email/password tokens
- Existing email/password authentication unaffected

## Dependencies

- Tasks 1-2 can run in parallel
- Task 3 depends on task 2 (migration must exist first)
- Task 4 can run in parallel with tasks 1-3
- Task 5 can run in parallel with tasks 1-4
- Task 6 depends on tasks 1, 3, 4, 5
- Task 7 depends on task 6
- Task 8 depends on task 2
- Task 9 depends on all previous tasks
- Task 10 can run anytime
