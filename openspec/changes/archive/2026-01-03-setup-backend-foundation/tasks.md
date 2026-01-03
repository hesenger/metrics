# Tasks: Setup Backend Foundation

## Implementation Steps

### 1. Initialize Go project
- [x] Initialize Go module with `go mod init`
- [x] Add dependencies: fiber, pgx5, sqlc, golang-migrate, bcrypt, jwt-go
- [x] Create project directory structure (cmd, internal, migrations, sql)

### 2. Setup sqlc configuration
- [x] Create `sqlc.yaml` configuration file
- [x] Configure sqlc to generate Go code in `internal/database` package
- [x] Set pgx5 as the database driver

### 3. Create database configuration
- [x] Implement `internal/config/config.go` to load environment variables
- [x] Read `DATABASE_URL` from environment
- [x] Read `JWT_SECRET` from environment
- [x] Add validation for required environment variables

### 4. Create database connection package
- [x] Implement `internal/database/conn.go`
- [x] Create connection pool using pgx5
- [x] Add connection verification with ping
- [x] Handle connection errors gracefully

### 5. Setup migration system
- [x] Create `migrations` directory
- [x] Write `000001_create_users_table.up.sql` migration
- [x] Write `000001_create_users_table.down.sql` migration
- [x] Implement migration runner in `internal/database/migrate.go`
- [x] Run migrations on application startup

### 6. Define user SQL queries
- [x] Create `sql/queries/users.sql` file
- [x] Write `CreateUser` query (INSERT)
- [x] Write `GetUserByEmail` query (SELECT)
- [x] Write `GetUserByID` query (SELECT)

### 7. Generate sqlc code
- [x] Run `sqlc generate` to create Go code
- [x] Verify generated code in `internal/database`
- [x] Ensure type safety for all queries

### 8. Implement authentication utilities
- [x] Create `internal/auth/password.go` for bcrypt hashing
- [x] Implement `HashPassword` function (cost factor 12)
- [x] Implement `VerifyPassword` function
- [x] Create `internal/auth/jwt.go` for token generation
- [x] Implement `GenerateToken` function with 24h expiration
- [x] Implement `ValidateToken` function

### 9. Implement user registration handler
- [x] Create `internal/handlers/auth.go`
- [x] Implement `Register` handler function
- [x] Validate email format
- [x] Validate password length (minimum 8 characters)
- [x] Check for duplicate email
- [x] Hash password
- [x] Insert user into database
- [x] Generate JWT token
- [x] Return user details and token (201 Created)
- [x] Handle errors (400 Bad Request, 409 Conflict, 500 Internal Server Error)

### 10. Implement user authentication handler
- [x] Implement `Login` handler function in `internal/handlers/auth.go`
- [x] Validate request payload
- [x] Fetch user by email
- [x] Verify password hash
- [x] Generate JWT token
- [x] Return user details and token (200 OK)
- [x] Handle errors (400 Bad Request, 401 Unauthorized, 500 Internal Server Error)

### 11. Setup Fiber application
- [x] Create `cmd/server/main.go` entry point
- [x] Initialize configuration
- [x] Establish database connection
- [x] Run migrations
- [x] Create Fiber app instance
- [x] Register `/api/auth/register` route (POST)
- [x] Register `/api/auth/login` route (POST)
- [x] Start server on configured port

### 12. Add development tooling
- [x] Create `Makefile` with common commands
- [x] Add `make run` to start server
- [x] Add `make migrate-up` to run migrations
- [x] Add `make migrate-down` to rollback migrations
- [x] Add `make sqlc` to generate code

### 13. Testing and validation
- [x] Test database connection with valid `DATABASE_URL`
- [x] Test application fails with missing `DATABASE_URL`
- [x] Test application fails with missing `JWT_SECRET`
- [x] Test migrations create users table successfully
- [x] Test registration with valid data returns 201 and token
- [x] Test registration with duplicate email returns 409
- [x] Test registration with invalid email returns 400
- [x] Test registration with short password returns 400
- [x] Test login with valid credentials returns 200 and token
- [x] Test login with invalid password returns 401
- [x] Test login with non-existent user returns 401
- [x] Test JWT token contains correct claims and expiration
- [x] Run `openspec validate setup-backend-foundation --strict`

## Dependencies and Parallelization
- Tasks 1-2 can run in parallel
- Task 3 depends on task 1
- Task 4 depends on task 3
- Tasks 5-6 can run in parallel after task 4
- Task 7 depends on task 5
- Task 8 can run in parallel with tasks 5-7
- Task 9 depends on tasks 7 and 8
- Task 10 depends on tasks 7 and 8
- Task 11 depends on tasks 9 and 10
- Task 12 can run in parallel with task 11
- Task 13 depends on all previous tasks

## Validation Checklist
- [x] All environment variables properly validated
- [x] Password hashing uses bcrypt with cost 12
- [x] JWT tokens expire after 24 hours
- [x] All error responses include meaningful messages
- [x] Database constraints enforce email uniqueness
- [x] All scenarios in specs are testable
- [x] Code follows project conventions (no comments, idiomatic Go)
