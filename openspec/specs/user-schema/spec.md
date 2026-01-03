# user-schema Specification

## Purpose
TBD - created by archiving change setup-backend-foundation. Update Purpose after archive.
## Requirements
### Requirement: User table structure
The system MUST provide a users table with essential fields for authentication including OAuth support.

#### Scenario: User table supports OAuth fields
**Given** migrations have been executed
**Then** the `users` table has an optional `oauth_provider` column (nullable string)
**And** the table has an optional `oauth_id` column (nullable string)
**And** there is a unique constraint on (oauth_provider, oauth_id) for non-null values
**And** the `password_hash` column is nullable to support OAuth users

#### Scenario: OAuth user without password
**Given** a user created via OAuth
**Then** the user has oauth_provider set (e.g., "google")
**And** the user has oauth_id set to the provider's user ID
**And** the user's password_hash is null
**And** the user's email is populated from OAuth provider

#### Scenario: Email/password user remains unchanged
**Given** a user created via email/password registration
**Then** the user has password_hash set
**And** the user's oauth_provider is null
**And** the user's oauth_id is null

#### Scenario: OAuth provider uniqueness
**Given** a user exists with oauth_provider="google" and oauth_id="123456"
**When** attempting to insert another user with the same oauth_provider and oauth_id
**Then** the database rejects the insert
**And** a unique constraint violation error is raised

### Requirement: User model type safety
The system MUST provide type-safe access to user data using sqlc-generated code.

#### Scenario: Type-safe user queries
**Given** sqlc has generated Go code from SQL queries
**Then** user queries have strongly-typed parameters
**And** query results map to Go structs
**And** the compiler catches type mismatches

