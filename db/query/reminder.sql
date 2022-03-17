-- name: CreateReminder :one
INSERT INTO reminders (full_name, personal_number, "user", phone_number) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetReminder :one
SELECT * FROM reminders
WHERE id = $1 LIMIT 1;

-- name: ListReminders :many
SELECT * FROM reminders
WHERE "user" = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateReminder :one
UPDATE reminders
SET personal_number = $2
WHERE id = $1
RETURNING *;

-- name: DeleteReminder :one
DELETE FROM reminders
WHERE id = $1
RETURNING *;