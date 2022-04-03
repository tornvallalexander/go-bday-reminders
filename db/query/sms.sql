-- name: GetSmsReminders :many
SELECT * FROM reminders
WHERE "user" = $1
ORDER BY personal_number;