# user-registration Specification

## ADDED Requirements

### Requirement: Registration sets HttpOnly cookie
The system MUST set JWT tokens in HttpOnly cookies after successful registration.

#### Scenario: Registration sets HttpOnly cookie
**Given** a new user successfully registers via `/api/auth/register`
**When** the JWT token is generated
**Then** the token is set in a cookie named "token"
**And** the cookie has HttpOnly flag set to true
**And** the cookie has SameSite flag set to Strict
**And** the cookie has Secure flag set to true in production
**And** the cookie expiration matches JWT expiration (7 days)
**And** the cookie path is set to "/"

#### Scenario: Registration response excludes token from body
**Given** a new user successfully registers
**When** the response is sent
**Then** the response body contains only user data: `{ user: { id, email, oauth_provider, created_at } }`
**And** the response body does NOT contain a "token" field
**And** the JWT is only in the cookie header

## MODIFIED Requirements

### Requirement: Register new user
The system MUST allow new users to register with email and password and set authentication cookie.

#### Scenario: Successful user registration
**Given** no user exists with email "newuser@example.com"
**When** a POST request is made to `/api/auth/register` with valid email and password
**Then** a new user is created in the database
**And** the password is hashed using bcrypt
**And** a JWT token is set in an HttpOnly cookie
**And** the response includes user details (id, email, oauth_provider, created_at)
**And** the response does NOT include a token field
**And** the response status is 201 Created

#### Scenario: Duplicate email registration
**Given** a user exists with email "existing@example.com"
**When** a POST request is made to `/api/auth/register` with the same email
**Then** the registration fails
**And** the response status is 409 Conflict
**And** an error message indicates the email is already registered
**And** no cookie is set

#### Scenario: Invalid email format
**Given** a registration request with invalid email format
**When** the request is processed
**Then** the registration fails
**And** the response status is 400 Bad Request
**And** an error message indicates invalid email format
**And** no cookie is set

#### Scenario: Password too short
**Given** a registration request with password less than 8 characters
**When** the request is processed
**Then** the registration fails
**And** the response status is 400 Bad Request
**And** an error message indicates password minimum length requirement
**And** no cookie is set

#### Scenario: Missing required fields
**Given** a registration request missing email or password
**When** the request is processed
**Then** the registration fails
**And** the response status is 400 Bad Request
**And** an error message indicates which fields are missing
**And** no cookie is set
