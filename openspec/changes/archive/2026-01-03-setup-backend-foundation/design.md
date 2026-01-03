# Design: Setup Backend Foundation

## Architecture Decisions

### Project Structure
```
/
├── cmd/
│   └── server/
│       └── main.go          # Entry point
├── internal/
│   ├── config/
│   │   └── config.go        # Environment configuration
│   ├── database/
│   │   └── db.go            # Database connection
│   ├── handlers/
│   │   └── auth.go          # Authentication handlers
│   └── models/
│       └── user.go          # User model
├── migrations/
│   └── 000001_create_users_table.up.sql
├── sql/
│   └── queries/
│       └── users.sql        # sqlc queries
└── sqlc.yaml                # sqlc configuration
```

### Database Configuration
- Use environment variables for database connection string
- Required env var: `DATABASE_URL` (PostgreSQL connection string)
- Format: `postgres://user:password@host:port/dbname?sslmode=disable`
- Use pgx5 driver for connection pooling and performance

### Migration System
- Golang Migrate for database migrations
- Migrations stored in `/migrations` directory
- Naming: `NNNNNN_description.up.sql` and `NNNNNN_description.down.sql`
- Run migrations on application startup

### User Schema
```sql
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
```

### Authentication Strategy
- Password-based authentication using bcrypt (cost factor 12)
- Session-based authentication using JWT tokens
- JWT secret from environment variable: `JWT_SECRET`
- Token expiration: 24 hours
- Return JWT token on successful registration and login

### API Endpoints
- `POST /api/auth/register` - Register new user
- `POST /api/auth/login` - Authenticate user

### Request/Response Format
**Registration Request:**
```json
{
  "email": "user@example.com",
  "password": "securepassword"
}
```

**Login Request:**
```json
{
  "email": "user@example.com",
  "password": "securepassword"
}
```

**Success Response:**
```json
{
  "token": "jwt.token.here",
  "user": {
    "id": 1,
    "email": "user@example.com",
    "created_at": "2026-01-03T10:00:00Z"
  }
}
```

**Error Response:**
```json
{
  "error": "error message"
}
```

### Validation Rules
- Email: Valid email format, required, max 255 characters
- Password: Minimum 8 characters, required
- Email uniqueness enforced at database level

### Technology Choices
- **Fiber**: Fast HTTP framework, Express-like API
- **sqlc**: Type-safe SQL code generation
- **pgx5**: High-performance PostgreSQL driver
- **Golang Migrate**: Simple and reliable migrations
- **bcrypt**: Industry standard for password hashing
- **JWT**: Stateless authentication tokens

### Trade-offs
1. **JWT vs Sessions**: Chose JWT for stateless authentication, simpler horizontal scaling
2. **bcrypt cost**: Cost factor 12 balances security with performance
3. **Migration timing**: Run on startup for simplicity, ensures schema is always current
