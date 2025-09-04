-- name: InsertUser :one
INSERT INTO users (name, password)
VALUES ($1, $2)
RETURNING id;

-- name: DeleteUserByIdAndName :exec
DELETE FROM users
WHERE id=$1 AND name=$2;

-- name: GetUserById :one
SELECT id, name, password
FROM users
WHERE id = $1;

-- name: GetUsersByName :many
SELECT id, name, password
FROM users
WHERE name = $1;