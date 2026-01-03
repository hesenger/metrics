-- name: CreateUser :one
INSERT INTO users (email, password_hash, created_at, updated_at)
VALUES ($1, $2, NOW(), NOW())
RETURNING id, email, created_at, updated_at;

-- name: GetUserByEmail :one
SELECT id, email, password_hash, created_at, updated_at
FROM users
WHERE email = $1;

-- name: GetUserByID :one
SELECT id, email, password_hash, created_at, updated_at
FROM users
WHERE id = $1;
