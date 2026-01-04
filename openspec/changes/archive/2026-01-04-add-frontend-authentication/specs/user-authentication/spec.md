# user-authentication Specification

## ADDED Requirements

### Requirement: JWT in HttpOnly cookie
The system MUST set JWT tokens in HttpOnly cookies for improved security.

#### Scenario: Login sets HttpOnly cookie
**Given** a user successfully authenticates via `/api/auth/login`
**When** the JWT token is generated
**Then** the token is set in a cookie named "token"
**And** the cookie has HttpOnly flag set to true
**And** the cookie has SameSite flag set to Strict
**And** the cookie has Secure flag set to true in production
**And** the cookie expiration matches JWT expiration (7 days)
**And** the cookie path is set to "/"

#### Scenario: Login response excludes token from body
**Given** a user successfully authenticates
**When** the response is sent
**Then** the response body contains only user data: `{ user: { id, email, oauth_provider, created_at } }`
**And** the response body does NOT contain a "token" field
**And** the JWT is only in the cookie header

### Requirement: Current user endpoint
The system MUST provide an endpoint to verify authentication and return current user data.

#### Scenario: Me endpoint with valid JWT
**Given** an authenticated user with a valid JWT in cookie
**When** a GET request is made to `/api/auth/me`
**Then** the JWT is validated from the cookie
**And** the cookie expiration is refreshed to 7 days from now
**And** the response status is 200 OK
**And** the response body contains user data: `{ user: { id, email, oauth_provider, created_at } }`

#### Scenario: Me endpoint without JWT
**Given** a request without a JWT cookie
**When** a GET request is made to `/api/auth/me`
**Then** the response status is 401 Unauthorized
**And** an error message indicates missing authentication

#### Scenario: Me endpoint with expired JWT
**Given** a request with an expired JWT cookie
**When** a GET request is made to `/api/auth/me`
**Then** the response status is 401 Unauthorized
**And** an error message indicates expired token
**And** the cookie is not refreshed

#### Scenario: Me endpoint with invalid JWT
**Given** a request with an invalid or malformed JWT cookie
**When** a GET request is made to `/api/auth/me`
**Then** the response status is 401 Unauthorized
**And** an error message indicates invalid token

#### Scenario: Me endpoint refreshes cookie on each call
**Given** an authenticated user makes a request to `/api/auth/me`
**When** the request is processed successfully
**Then** the JWT cookie expiration is extended to 7 days from the request time
**And** this allows active users to maintain their session indefinitely

### Requirement: Logout endpoint
The system MUST provide an endpoint to clear authentication cookies.

#### Scenario: Logout clears cookie
**Given** an authenticated user
**When** a POST request is made to `/api/auth/logout`
**Then** the "token" cookie is cleared
**And** the cookie expiration is set to a past date
**And** the response status is 204 No Content
**And** no response body is sent

#### Scenario: Logout without authentication
**Given** an unauthenticated user
**When** a POST request is made to `/api/auth/logout`
**Then** the response status is 204 No Content
**And** no error is returned

## MODIFIED Requirements

### Requirement: JWT token generation
The system MUST generate secure JWT tokens for authenticated users and set them in HttpOnly cookies.

#### Scenario: Token includes user claims
**Given** a user successfully authenticates
**When** a JWT token is generated
**Then** the token includes the user ID in claims
**And** the token includes the user email in claims
**And** the token has an expiration time of 7 days
**And** the token is signed with the `JWT_SECRET` from environment
**And** the token is set in an HttpOnly cookie (not response body)

#### Scenario: Token secret from environment
**Given** the `JWT_SECRET` environment variable is not set
**When** the application starts
**Then** the application fails to start
**And** an error message indicates the missing JWT secret configuration
