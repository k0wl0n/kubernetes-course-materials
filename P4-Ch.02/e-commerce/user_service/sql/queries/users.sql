-- name: CreateUser :one
INSERT INTO users (id, created_at, modified_at, email, password) 
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE id=$1 FOR UPDATE NOWAIT;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email=$1 FOR UPDATE NOWAIT;

-- name: UpdateUserDetails :one
UPDATE users SET email=$1, name=$2, modified_at=$3
WHERE id=$4
RETURNING *;

-- name: UpdateUserPassword :exec
UPDATE users SET password=$1, modified_at=$2
WHERE id=$3;

-- name: DeleteUser :exec
DELETE FROM users WHERE id=$1;