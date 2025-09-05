-- name: InsertUser :one
INSERT INTO users (username, email, password, name)
VALUES ($1, $2, $3, $4)
RETURNING id, username, email, name, created_at, updated_at;

-- name: DeleteUser :exec
DELETE FROM users WHERE id=$1;

-- name: GetUserByUsername :one
SELECT * FROM users where username = $1;

-- name: GetUserByEmail :one
SELECT * FROM users where email = $1;

-- name: ListUsers :many
SELECT id, username, email, name, created_at, updated_at
FROM users
ORDER BY id;

-- name: UpdateUser :one
UPDATE users
SET name = $2, updated_at = now()
WHERE id = $1
RETURNING id, username, email, name, created_at, updated_at;