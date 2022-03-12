-- name: CreateUser :one
INSERT INTO users (
    user_name,
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
WHERE user_name = $1
LIMIT 1;

-- name: DeleteUser :one
DELETE FROM users
WHERE user_name = $1
RETURNING user_name;