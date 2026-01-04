-- name: CreateUser :one
INSERT INTO users (email, password_hash, created_at, updated_at)
VALUES ($1, $2, NOW(), NOW())
RETURNING id, email, oauth_provider, oauth_id, created_at, updated_at;

-- name: GetUserByEmail :one
SELECT id, email, password_hash, oauth_provider, oauth_id, created_at, updated_at
FROM users
WHERE email = $1;

-- name: GetUserByID :one
SELECT id, email, password_hash, oauth_provider, oauth_id, created_at, updated_at
FROM users
WHERE id = $1;

-- name: CreateOAuthUser :one
INSERT INTO users (email, oauth_provider, oauth_id, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
RETURNING id, email, oauth_provider, oauth_id, created_at, updated_at;

-- name: GetUserByOAuthProvider :one
SELECT id, email, password_hash, oauth_provider, oauth_id, created_at, updated_at
FROM users
WHERE oauth_provider = $1 AND oauth_id = $2;
