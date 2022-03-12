// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
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
) RETURNING user_name, hashed_password, full_name, email, phone_number, password_changed_at, created_at
`

type CreateUserParams struct {
	UserName       string `json:"user_name"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	PhoneNumber    int64  `json:"phone_number"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.UserName,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
		arg.PhoneNumber,
	)
	var i User
	err := row.Scan(
		&i.UserName,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PhoneNumber,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :one
DELETE FROM users
WHERE user_name = $1
RETURNING user_name
`

func (q *Queries) DeleteUser(ctx context.Context, userName string) (string, error) {
	row := q.db.QueryRowContext(ctx, deleteUser, userName)
	var user_name string
	err := row.Scan(&user_name)
	return user_name, err
}

const getUser = `-- name: GetUser :one
SELECT user_name, hashed_password, full_name, email, phone_number, password_changed_at, created_at FROM users
WHERE user_name = $1
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, userName string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, userName)
	var i User
	err := row.Scan(
		&i.UserName,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PhoneNumber,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}