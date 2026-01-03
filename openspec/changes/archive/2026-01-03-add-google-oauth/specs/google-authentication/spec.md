# google-authentication Specification Delta

## Purpose
Enable users to authenticate using their Google accounts via OAuth 2.0.

## ADDED Requirements

### Requirement: Initiate Google OAuth flow
The system MUST provide an endpoint to initiate the Google OAuth 2.0 authentication flow.

#### Scenario: Start OAuth flow
**Given** the application is configured with Google OAuth credentials
**When** a GET request is made to `/api/auth/google`
**Then** the system generates a random state parameter
**And** the state is stored with expiration time
**And** the system redirects to Google's OAuth consent screen
**And** the redirect URL includes the client ID, redirect URI, scope, and state

#### Scenario: Missing Google OAuth configuration
**Given** `GOOGLE_CLIENT_ID` or `GOOGLE_CLIENT_SECRET` is not set
**When** the application starts
**Then** the application fails to start
**And** an error message indicates missing Google OAuth configuration

### Requirement: Handle Google OAuth callback
The system MUST handle the OAuth callback from Google and authenticate the user.

#### Scenario: Successful OAuth callback for new user
**Given** a user completes Google OAuth consent
**And** the state parameter is valid and not expired
**And** no user exists with the Google ID
**When** Google redirects to `/api/auth/google/callback` with authorization code
**Then** the system exchanges the code for an access token
**And** the system fetches user info from Google (email, google_id)
**And** a new user is created with oauth_provider="google" and oauth_id set
**And** password_hash is null for the OAuth user
**And** a JWT token is generated
**And** the response includes the JWT token
**And** the response includes user details (id, email, created_at)

#### Scenario: Successful OAuth callback for existing user
**Given** a user with oauth_provider="google" and oauth_id="123456" exists
**And** the state parameter is valid
**When** that user completes OAuth and Google redirects with authorization code
**Then** the system exchanges the code for an access token
**And** the system fetches user info from Google
**And** the existing user is found by oauth_provider and oauth_id
**And** a JWT token is generated for the existing user
**And** the response includes the JWT token and user details

#### Scenario: Invalid state parameter
**Given** a callback request with an invalid or expired state
**When** the callback is processed
**Then** the authentication fails
**And** the response status is 400 Bad Request
**And** an error message indicates invalid state

#### Scenario: OAuth code exchange fails
**Given** Google returns an error during code exchange
**When** the callback is processed
**Then** the authentication fails
**And** the response status is 500 Internal Server Error
**And** an error message indicates OAuth failure

#### Scenario: Duplicate email from different provider
**Given** a user exists with email "user@example.com" via email/password
**And** a Google user attempts to sign in with the same email
**When** the OAuth callback is processed
**Then** a new separate user is created with oauth_provider="google"
**And** both users can coexist with the same email (no linking)

### Requirement: OAuth user info retrieval
The system MUST fetch user information from Google's OAuth API.

#### Scenario: Fetch user info from Google
**Given** a valid Google access token
**When** the system requests user info from Google
**Then** the response includes the user's email
**And** the response includes the user's Google ID
**And** the system uses this data to create or identify the user
