# user-registration Specification

## Purpose
TBD - created by archiving change setup-backend-foundation. Update Purpose after archive.
## Requirements
### Requirement: Register new user
The system MUST allow new users to register with email and password.

#### Scenario: Successful user registration
**Given** no user exists with email "newuser@example.com"
**When** a POST request is made to `/api/auth/register` with valid email and password
**Then** a new user is created in the database
**And** the password is hashed using bcrypt
**And** a JWT token is returned
**And** the response includes user details (id, email, created_at)
**And** the response status is 201 Created

#### Scenario: Duplicate email registration
**Given** a user exists with email "existing@example.com"
**When** a POST request is made to `/api/auth/register` with the same email
**Then** the registration fails
**And** the response status is 409 Conflict
**And** an error message indicates the email is already registered

#### Scenario: Invalid email format
**Given** a registration request with invalid email format
**When** the request is processed
**Then** the registration fails
**And** the response status is 400 Bad Request
**And** an error message indicates invalid email format

#### Scenario: Password too short
**Given** a registration request with password less than 8 characters
**When** the request is processed
**Then** the registration fails
**And** the response status is 400 Bad Request
**And** an error message indicates password minimum length requirement

#### Scenario: Missing required fields
**Given** a registration request missing email or password
**When** the request is processed
**Then** the registration fails
**And** the response status is 400 Bad Request
**And** an error message indicates which fields are missing

### Requirement: Password security
User passwords MUST be securely hashed before storage.

#### Scenario: Password is hashed
**Given** a user registers with password "mysecretpassword"
**When** the user is created
**Then** the password is hashed using bcrypt cost factor 12
**And** the plain password is never stored
**And** the hash is stored in the `password_hash` column

