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

func TestGetBirthdayAPI(t *testing.T) {
	birthday := randomBirthday()

	testCases := []struct {
		name          string
		birthdayID    int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:       "OK",
			birthdayID: birthday.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetBirthday(gomock.Any(), gomock.Eq(birthday.ID)).
					Times(1).
					Return(birthday, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchBirthday(t, recorder.Body, birthday)
			},
		},
		{
			name:       "NotFound",
			birthdayID: birthday.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetBirthday(gomock.Any(), gomock.Eq(birthday.ID)).
					Times(1).
					Return(db.Birthday{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:       "InternalError",
			birthdayID: birthday.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetBirthday(gomock.Any(), gomock.Eq(birthday.ID)).
					Times(1).
					Return(db.Birthday{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:       "InvalidID",
			birthdayID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetBirthday(gomock.Any(), gomock.Any()).
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
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/birthdays/%d", tc.birthdayID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}

func randomBirthday() db.Birthday {
	return db.Birthday{
		ID:             utils.RandomInt(1, 1000),
		FullName:       utils.RandomFullName(),
		FutureBirthday: utils.RandomDate(),
	}
}

func requireBodyMatchBirthday(t *testing.T, body *bytes.Buffer, birthday db.Birthday) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotBirthday db.Birthday
	err = json.Unmarshal(data, &gotBirthday)
	require.NoError(t, err)
	require.Equal(t, birthday, gotBirthday)
}
