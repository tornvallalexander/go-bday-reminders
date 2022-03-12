-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email,
    phone_number
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1
LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;