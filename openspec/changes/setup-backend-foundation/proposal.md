# Setup Backend Foundation

## Overview
Establish the Go backend foundation with Fiber web framework, sqlc for type-safe SQL queries, database configuration from environment variables, migration system, and user management with registration and authentication capabilities.

## Motivation
The project requires a robust backend foundation to support the metrics SaaS platform. This change provides the essential infrastructure including database connectivity, schema management, and user authentication that all future features will depend on.

## Scope
This change introduces:
- Database configuration from environment variables
- Migration system using Golang Migrate
- User table schema
- User registration endpoint
- User authentication endpoint

## Out of Scope
- Password reset functionality
- Email verification
- OAuth providers
- User profile management
- Role-based access control

## Dependencies
None - this is the foundational change for the backend.

## Risks and Mitigations
- **Risk**: Storing passwords securely
  - **Mitigation**: Use bcrypt for password hashing with appropriate cost factor

- **Risk**: Database connection failures
  - **Mitigation**: Implement connection retry logic and proper error handling

## Success Criteria
- Database connects successfully using environment variables
- Migrations run successfully creating user table
- Users can register with email and password
- Users can authenticate and receive session token
- All validations pass with `openspec validate setup-backend-foundation --strict`
