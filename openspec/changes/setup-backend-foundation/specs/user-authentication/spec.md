# Spec: User Authentication

## ADDED Requirements

### Requirement: Authenticate existing user
The system MUST allow registered users to authenticate with email and password.

#### Scenario: Successful authentication
**Given** a user exists with email "user@example.com" and password "correctpassword"
**When** a POST request is made to `/api/auth/login` with correct credentials
**Then** the password is verified against the stored hash
**And** a JWT token is generated
**And** the response includes the JWT token
**And** the response includes user details (id, email, created_at)
**And** the response status is 200 OK

#### Scenario: Invalid password
**Given** a user exists with email "user@example.com"
**When** a POST request is made to `/api/auth/login` with incorrect password
**Then** the authentication fails
**And** the response status is 401 Unauthorized
**And** an error message indicates invalid credentials

#### Scenario: Non-existent user
**Given** no user exists with email "nonexistent@example.com"
**When** a POST request is made to `/api/auth/login` with that email
**Then** the authentication fails
**And** the response status is 401 Unauthorized
**And** an error message indicates invalid credentials

#### Scenario: Missing credentials
**Given** a login request missing email or password
**When** the request is processed
**Then** the authentication fails
**And** the response status is 400 Bad Request
**And** an error message indicates which fields are missing

### Requirement: JWT token generation
The system MUST generate secure JWT tokens for authenticated users.

#### Scenario: Token includes user claims
**Given** a user successfully authenticates
**When** a JWT token is generated
**Then** the token includes the user ID in claims
**And** the token includes the user email in claims
**And** the token has an expiration time of 24 hours
**And** the token is signed with the `JWT_SECRET` from environment

#### Scenario: Token secret from environment
**Given** the `JWT_SECRET` environment variable is not set
**When** the application starts
**Then** the application fails to start
**And** an error message indicates the missing JWT secret configuration
