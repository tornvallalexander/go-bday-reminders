package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "go-bday-reminders/db/mock"
	db "go-bday-reminders/db/sqlc"
	"go-bday-reminders/utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetReminderAPI(t *testing.T) {
	reminder := randomReminder(t)

	testCases := []struct {
		name          string
		reminderID    int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:       "OK",
			reminderID: reminder.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetReminder(gomock.Any(), gomock.Eq(reminder.ID)).
					Times(1).
					Return(reminder, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchReminder(t, recorder.Body, reminder)
			},
		},
		{
			name:       "NotFound",
			reminderID: reminder.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetReminder(gomock.Any(), gomock.Eq(reminder.ID)).
					Times(1).
					Return(db.Reminder{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:       "InternalError",
			reminderID: reminder.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetReminder(gomock.Any(), gomock.Eq(reminder.ID)).
					Times(1).
					Return(db.Reminder{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:       "InvalidID",
			reminderID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetReminder(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// start test server and send request
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/reminders/%d", tc.reminderID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}

func randomReminder(t *testing.T) db.Reminder {
	user, _ := randomUser(t)
	return db.Reminder{
		ID:             utils.RandomInt(1, 1000),
		FullName:       utils.RandomFullName(),
		PersonalNumber: utils.RandomPnr(),
		User:           user.Username,
		PhoneNumber:    utils.RandomPhoneNumber(),
		CreatedAt:      utils.RandomDate(),
	}
}

func requireBodyMatchReminder(t *testing.T, body *bytes.Buffer, reminder db.Reminder) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotReminder db.Reminder
	err = json.Unmarshal(data, &gotReminder)
	require.NoError(t, err)
	require.Equal(t, reminder, gotReminder)
}
