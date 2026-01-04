# google-authentication Specification

## ADDED Requirements

### Requirement: Google OAuth sets HttpOnly cookie
The system MUST set JWT tokens in HttpOnly cookies after successful Google OAuth.

#### Scenario: OAuth callback sets HttpOnly cookie
**Given** a user completes Google OAuth successfully
**When** the JWT token is generated
**Then** the token is set in a cookie named "token"
**And** the cookie has HttpOnly flag set to true
**And** the cookie has SameSite flag set to Strict
**And** the cookie has Secure flag set to true in production
**And** the cookie expiration matches JWT expiration (7 days)
**And** the cookie path is set to "/"

### Requirement: OAuth callback redirects to frontend with user data
The system MUST redirect to the frontend application with user data after successful OAuth.

#### Scenario: Successful OAuth redirects to frontend
**Given** a user completes Google OAuth successfully
**When** the callback is processed
**Then** the JWT is set in a cookie
**And** the backend redirects to the frontend root URL
**And** the redirect URL includes user data as query parameters
**And** the query parameters include: id, email, oauth_provider
**And** the redirect uses 302 Found status

#### Scenario: Failed OAuth redirects to login with error
**Given** the OAuth flow fails
**When** the callback is processed with an error
**Then** the backend redirects to `/login?error=oauth_failed`
**And** no cookie is set
**And** no user data is included

## MODIFIED Requirements

### Requirement: Handle Google OAuth callback
The system MUST handle the OAuth callback from Google, authenticate the user, and redirect to frontend.

#### Scenario: Successful OAuth callback for new user
**Given** a user completes Google OAuth consent
**And** the state parameter is valid and not expired
**And** no user exists with the Google ID
**When** Google redirects to `/api/auth/google/callback` with authorization code
**Then** the system exchanges the code for an access token
**And** the system fetches user info from Google (email, google_id)
**And** a new user is created with oauth_provider="google" and oauth_id set
**And** password_hash is null for the OAuth user
**And** a JWT token is set in an HttpOnly cookie
**And** the user is redirected to frontend with user data in URL parameters
**And** no token is in the response body

#### Scenario: Successful OAuth callback for existing user
**Given** a user with oauth_provider="google" and oauth_id="123456" exists
**And** the state parameter is valid
**When** that user completes OAuth and Google redirects with authorization code
**Then** the system exchanges the code for an access token
**And** the system fetches user info from Google
**And** the existing user is found by oauth_provider and oauth_id
**And** a JWT token is set in an HttpOnly cookie
**And** the user is redirected to frontend with user data in URL parameters
**And** no token is in the response body

#### Scenario: Invalid state parameter
**Given** a callback request with an invalid or expired state
**When** the callback is processed
**Then** the authentication fails
**And** the user is redirected to `/login?error=invalid_state`
**And** no cookie is set

#### Scenario: OAuth code exchange fails
**Given** Google returns an error during code exchange
**When** the callback is processed
**Then** the authentication fails
**And** the user is redirected to `/login?error=oauth_failed`
**And** no cookie is set

#### Scenario: Duplicate email from different provider
**Given** a user exists with email "user@example.com" via email/password
**And** a Google user attempts to sign in with the same email
**When** the OAuth callback is processed
**Then** a new separate user is created with oauth_provider="google"
**And** both users can coexist with the same email (no linking)
**And** a JWT cookie is set for the new Google user
