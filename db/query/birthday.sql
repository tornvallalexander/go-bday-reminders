-- name: CreateBirthday :one
INSERT INTO birthdays (
    full_name,
    future_birthday
) VALUES (
   $1, $2
) RETURNING *;

-- name: GetBirthday :one
SELECT * FROM birthdays
WHERE id = $1 LIMIT 1;

-- name: ListBirthdays :many
SELECT * FROM birthdays
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateBirthday :one
UPDATE birthdays
SET future_birthday = $2
WHERE id = $1
RETURNING *;

-- name: DeleteBirthday :one
DELETE FROM birthdays
WHERE id = $1
RETURNING *;