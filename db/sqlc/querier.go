// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateReminder(ctx context.Context, arg CreateReminderParams) (Reminder, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteReminder(ctx context.Context, id int64) (Reminder, error)
	DeleteUser(ctx context.Context, username string) error
	GetReminder(ctx context.Context, id int64) (Reminder, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetSmsReminders(ctx context.Context, user string) ([]Reminder, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListReminders(ctx context.Context, arg ListRemindersParams) ([]Reminder, error)
	UpdateReminder(ctx context.Context, arg UpdateReminderParams) (Reminder, error)
}

var _ Querier = (*Queries)(nil)
