package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/atercode/SimplySacco/db/sqlc"
	"github.com/atercode/SimplySacco/utils"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestCreateMemberAPI(t *testing.T) {
	member, password := randomMember(t)

	testCases := []struct {
		name string
		body gin.H
		// buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"full_name":   member.FullName,
				"email":       member.Email,
				"password":    password,
				"status_code": member.StatusCode,
			},
			// buildStubs: func(store *mockdb.MockStore) {
			// 	arg := db.CreateMemberParams{
			// 		FullName: member.FullName,
			// 		Email:    member.Email,
			// 		StatusCode: member.StatusCode,
			// 	}
			// 	store.EXPECT().
			// 		CreateMember(gomock.Any(), EqCreateMemberParams(arg, password)).
			// 		Times(1).
			// 		Return(user, nil)
			// },
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, member)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"full_name":   member.FullName,
				"email":       member.Email,
				"password":    password,
				"status_code": member.StatusCode,
			},
			// buildStubs: func(store *mockdb.MockStore) {
			// 	store.EXPECT().
			// 		CreateMember(gomock.Any(), gomock.Any()).
			// 		Times(1).
			// 		Return(db.Member{}, sql.ErrConnDone)
			// },
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		// {
		// 	name: "DuplicateUsername",
		// 	body: gin.H{
		// 		"full_name":   member.FullName,
		// 		"email":       member.Email,
		// 		"password":    password,
		// 		"status_code": member.StatusCode,
		// 	},
		// 	// buildStubs: func(store *mockdb.MockStore) {
		// 	// 	store.EXPECT().
		// 	// 		CreateMember(gomock.Any(), gomock.Any()).
		// 	// 		Times(1).
		// 	// 		Return(db.Member{}, &pq.Error{Code: "23505"})
		// 	// },
		// 	checkResponse: func(recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusForbidden, recorder.Code)
		// 	},
		// },
		{
			name: "InvalidEmail",
			body: gin.H{
				"full_name":   member.FullName,
				"email":       "ssaad",
				"password":    password,
				"status_code": member.StatusCode,
			},
			// buildStubs: func(store *mockdb.MockStore) {
			// 	store.EXPECT().
			// 		CreateMember(gomock.Any(), gomock.Any()).
			// 		Times(0)
			// },
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "TooShortPassword",
			body: gin.H{
				"full_name":   member.FullName,
				"email":       member.Email,
				"password":    "1345",
				"status_code": member.StatusCode,
			},
			// buildStubs: func(store *mockdb.MockStore) {
			// 	store.EXPECT().
			// 		CreateMember(gomock.Any(), gomock.Any()).
			// 		Times(0)
			// },
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			// ctrl := gomock.NewController(t)
			// defer ctrl.Finish()

			// store := mockdb.NewMockStore(ctrl)
			// tc.buildStubs(store)
			server := newTestServer(t)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/members"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomMember(t *testing.T) (member db.Member, password string) {
	password = gofakeit.DigitN(6)
	hashedPassword, err := utils.HashPassword(password)
	require.NoError(t, err)

	member = db.Member{
		FullName:       gofakeit.Name(),
		Email:          gofakeit.Email(),
		StatusCode:     "TEST",
		HashedPassword: hashedPassword,
	}
	return
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, member db.Member) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var givenMember db.Member
	err = json.Unmarshal(data, &givenMember)

	require.NoError(t, err)
	require.Equal(t, member.StatusCode, givenMember.StatusCode)
	require.Equal(t, member.FullName, givenMember.FullName)
	require.Equal(t, member.Email, givenMember.Email)
	// require.Empty(t, givenMember.HashedPassword)
}
