-- name: CreateReminder :one
INSERT INTO reminders (full_name, personal_number, "user", phone_number) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetReminder :one
SELECT * FROM reminders
WHERE id = $1 LIMIT 1;

-- name: ListReminders :many
SELECT * FROM reminders
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateReminder :one
UPDATE reminders
SET personal_number = $2
WHERE id = $1
RETURNING *;

-- name: DeleteReminder :one
DELETE FROM reminders
WHERE id = $1
RETURNING *;