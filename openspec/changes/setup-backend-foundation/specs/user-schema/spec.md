# Spec: User Schema

## ADDED Requirements

### Requirement: User table structure
The system MUST provide a users table with essential fields for authentication.

#### Scenario: User table exists
**Given** migrations have been executed
**Then** a `users` table exists in the database
**And** the table has an `id` column as primary key
**And** the table has an `email` column with unique constraint
**And** the table has a `password_hash` column
**And** the table has `created_at` and `updated_at` timestamp columns

#### Scenario: Email uniqueness is enforced
**Given** a user exists with email "user@example.com"
**When** attempting to insert another user with the same email
**Then** the database rejects the insert
**And** a unique constraint violation error is raised

### Requirement: User model type safety
The system MUST provide type-safe access to user data using sqlc-generated code.

#### Scenario: Type-safe user queries
**Given** sqlc has generated Go code from SQL queries
**Then** user queries have strongly-typed parameters
**And** query results map to Go structs
**And** the compiler catches type mismatches
