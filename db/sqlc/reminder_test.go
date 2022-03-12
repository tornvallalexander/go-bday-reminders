package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"go-bday-reminders/utils"
	"testing"
)

// we don't want tests to depend on each other, therefore
// we create a reusable function for multiple tests
func createRandomReminder(t *testing.T) Reminder {
	user := createRandomUser(t)
	arg := CreateReminderParams{
		FullName:       utils.RandomFullName(),
		PersonalNumber: utils.RandomPnr(),
		User:           user.UserName,
		PhoneNumber:    utils.RandomPhoneNumber(),
	}

	reminder, err := testQueries.CreateReminder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, reminder)

	require.Equal(t, arg.FullName, reminder.FullName)

	require.NotZero(t, reminder.PersonalNumber)
	require.NotZero(t, reminder.ID)

	return reminder
}

func TestCreateReminder(t *testing.T) {
	createRandomReminder(t)
}

func TestGetReminder(t *testing.T) {
	reminder1 := createRandomReminder(t)
	reminder2, err := testQueries.GetReminder(context.Background(), reminder1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, reminder2)

	require.Equal(t, reminder1.ID, reminder2.ID)
	require.Equal(t, reminder1.PersonalNumber, reminder2.PersonalNumber)
	require.Equal(t, reminder1.FullName, reminder2.FullName)
}

func TestUpdateReminder(t *testing.T) {
	reminder1 := createRandomReminder(t)
	arg := UpdateReminderParams{
		ID:             reminder1.ID,
		PersonalNumber: utils.RandomPnr(),
	}

	reminder2, err := testQueries.UpdateReminder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, reminder2)
	require.NotEqual(t, reminder2.PersonalNumber, reminder1.PersonalNumber)

	require.Equal(t, reminder2.ID, reminder1.ID)
	require.Equal(t, reminder2.FullName, reminder1.FullName)
	require.Equal(t, reminder2.ID, arg.ID)

	require.NotZero(t, reminder2.PersonalNumber)
}

func TestDeleteReminder(t *testing.T) {
	reminder1 := createRandomReminder(t)
	_, err := testQueries.DeleteReminder(context.Background(), reminder1.ID)
	require.NoError(t, err)

	reminder2, err := testQueries.GetReminder(context.Background(), reminder1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, reminder2)
}

func TestListReminders(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomReminder(t)
	}

	arg := ListRemindersParams{
		Limit:  5,
		Offset: 5,
	}

	reminders, err := testQueries.ListReminders(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, reminders, 5)

	for _, reminder := range reminders {
		require.NotEmpty(t, reminder)

	}
}
