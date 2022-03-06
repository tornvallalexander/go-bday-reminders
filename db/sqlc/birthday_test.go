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
func createRandomBirthday(t *testing.T) Birthday {
	arg := CreateBirthdayParams{
		FullName:       utils.RandomFullName(),
		FutureBirthday: utils.RandomDate(),
	}

	birthday, err := testQueries.CreateBirthday(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, birthday)

	require.Equal(t, arg.FullName, birthday.FullName)

	require.NotZero(t, birthday.FutureBirthday)
	require.NotZero(t, birthday.ID)

	return birthday
}

func TestCreateBirthday(t *testing.T) {
	createRandomBirthday(t)
}

func TestGetBirthday(t *testing.T) {
	birthday1 := createRandomBirthday(t)
	birthday2, err := testQueries.GetBirthday(context.Background(), birthday1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, birthday2)

	require.Equal(t, birthday1.ID, birthday2.ID)
	require.Equal(t, birthday1.FutureBirthday, birthday2.FutureBirthday)
	require.Equal(t, birthday1.FullName, birthday2.FullName)
}

func TestUpdateBirthday(t *testing.T) {
	birthday1 := createRandomBirthday(t)
	arg := UpdateBirthdayParams{
		ID:             birthday1.ID,
		FutureBirthday: utils.RandomDate(),
	}

	birthday2, err := testQueries.UpdateBirthday(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, birthday2)
	require.NotEqual(t, birthday2.FutureBirthday, birthday1.FutureBirthday)

	require.Equal(t, birthday2.ID, birthday1.ID)
	require.Equal(t, birthday2.FullName, birthday1.FullName)
	require.Equal(t, birthday2.ID, arg.ID)
	require.Equal(t, birthday2.FutureBirthday, arg.FutureBirthday)
}

func TestDeleteBirthday(t *testing.T) {
	birthday1 := createRandomBirthday(t)
	_, err := testQueries.DeleteBirthday(context.Background(), birthday1.ID)
	require.NoError(t, err)

	birthday2, err := testQueries.GetBirthday(context.Background(), birthday1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, birthday2)
}

func TestListBirthdays(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomBirthday(t)
	}

	arg := ListBirthdaysParams{
		Limit:  5,
		Offset: 5,
	}

	birthdays, err := testQueries.ListBirthdays(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, birthdays, 5)

	for _, birthday := range birthdays {
		require.NotEmpty(t, birthday)

	}
}
