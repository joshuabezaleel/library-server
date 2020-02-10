package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
)

func TestUserCreate(t *testing.T) {
	initialUser := &user.User{
		ID: util.NewID(),
	}

	failedUser := &user.User{
		ID: util.NewID(),
	}

	tt := []struct {
		name              string
		requestPayload    interface{}
		mockReturnPayload interface{}
		ID                string
		statusCode        int
		err               error
	}{
		{
			name:              "success creating a valid User",
			requestPayload:    initialUser,
			mockReturnPayload: initialUser,
			ID:                initialUser.ID,
			statusCode:        http.StatusCreated,
			err:               nil,
		},
		{
			name:              "invalid request payload",
			requestPayload:    "a plain string, not a User",
			mockReturnPayload: nil,
			ID:                "",
			statusCode:        http.StatusBadRequest,
			err:               nil,
		},
		{
			name:              "failed creating a User",
			requestPayload:    failedUser,
			mockReturnPayload: nil,
			ID:                failedUser.ID,
			statusCode:        http.StatusInternalServerError,
			err:               errors.New("Users not found"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userService.On("Create", tc.requestPayload).Return(tc.mockReturnPayload, tc.err)

			url := fmt.Sprintf("/users")

			reqByte, err := json.Marshal(tc.requestPayload)
			require.Nil(t, err)

			req := httptest.NewRequest("POST", url, bytes.NewReader(reqByte))

			if tc.statusCode != http.StatusBadRequest {
				req = mux.SetURLVars(req, map[string]string{"userID": tc.ID})
			}

			w := httptest.NewRecorder()

			userTestingHandler.createUser(w, req)

			require.Equal(t, tc.statusCode, w.Code)
		})
	}
}

func TestUserGet(t *testing.T) {
	initialUser := &user.User{
		ID: util.NewID(),
	}

	tt := []struct {
		name              string
		mockReturnPayload interface{}
		ID                string
		statusCode        int
		err               error
	}{
		{
			name:              "success retrieving a valid user",
			mockReturnPayload: initialUser,
			ID:                initialUser.ID,
			statusCode:        http.StatusOK,
			err:               nil,
		},
		{
			name:              "invalid path",
			mockReturnPayload: nil,
			ID:                "invalidUserID",
			statusCode:        http.StatusBadRequest,
			err:               nil,
		},
		{
			name:              "user doesn't exist",
			mockReturnPayload: nil,
			ID:                util.NewID(),
			statusCode:        http.StatusInternalServerError,
			err:               errors.New("Users not found"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userService.On("Get", tc.ID).Return(tc.mockReturnPayload, tc.err)

			url := fmt.Sprintf("/users/" + tc.ID)

			req := httptest.NewRequest("GET", url, nil)

			if tc.statusCode != http.StatusBadRequest {
				req = mux.SetURLVars(req, map[string]string{"userID": tc.ID})
			}

			w := httptest.NewRecorder()

			userTestingHandler.getUser(w, req)

			require.Equal(t, tc.statusCode, w.Code)
		})
	}
}

func TestUserUpdate(t *testing.T) {
	initialUser := &user.User{
		ID: util.NewID(),
	}

	editedUser := initialUser
	editedUser.Username = "edited username"

	failedUser := &user.User{
		ID: util.NewID(),
	}

	tt := []struct {
		name              string
		requestPayload    interface{}
		mockReturnPayload interface{}
		ID                string
		statusCode        int
		err               error
	}{
		{
			name:              "success updating a valid User",
			requestPayload:    editedUser,
			mockReturnPayload: editedUser,
			ID:                initialUser.ID,
			statusCode:        http.StatusOK,
			err:               nil,
		},
		{
			name:              "invalid User ID path",
			requestPayload:    editedUser,
			mockReturnPayload: nil,
			ID:                "invalidUserID",
			statusCode:        http.StatusBadRequest,
			err:               nil,
		},
		{
			name:              "request payload is not a User",
			requestPayload:    "a plain string",
			mockReturnPayload: nil,
			ID:                initialUser.ID,
			statusCode:        http.StatusBadRequest,
			err:               nil,
		},
		{
			name:              "failed updating a User",
			requestPayload:    failedUser,
			mockReturnPayload: nil,
			ID:                failedUser.ID,
			statusCode:        http.StatusInternalServerError,
			err:               errors.New("Users not found"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userService.On("Update", tc.requestPayload).Return(tc.mockReturnPayload, tc.err)

			url := fmt.Sprintf("/users/" + tc.ID)

			reqByte, err := json.Marshal(tc.requestPayload)
			require.Nil(t, err)

			req := httptest.NewRequest("PUT", url, bytes.NewReader(reqByte))

			if tc.statusCode != http.StatusBadRequest {
				req = mux.SetURLVars(req, map[string]string{"userID": tc.ID})
			}

			w := httptest.NewRecorder()

			userTestingHandler.updateUser(w, req)

			if tc.statusCode != http.StatusOK {
				require.Equal(t, tc.statusCode, w.Code)
				return
			}

			// If we want to test the title, we need to assert the type of the interface{}

		})
	}
}

func TestUserDelete(t *testing.T) {
	initialUser := &user.User{
		ID: util.NewID(),
	}

	tt := []struct {
		name              string
		mockReturnPayload interface{}
		ID                string
		statusCode        int
		err               error
	}{
		{
			name:              "success deleting a valid user",
			mockReturnPayload: initialUser,
			ID:                initialUser.ID,
			statusCode:        http.StatusOK,
			err:               nil,
		},
		{
			name:              "invalid path",
			mockReturnPayload: nil,
			ID:                "invalidUserID",
			statusCode:        http.StatusBadRequest,
			err:               nil,
		},
		{
			name:              "user doesn't exist",
			mockReturnPayload: nil,
			ID:                util.NewID(),
			statusCode:        http.StatusInternalServerError,
			err:               errors.New("Users not found"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userService.On("Delete", tc.ID).Return(tc.err)

			url := fmt.Sprintf("/users/" + tc.ID)

			req := httptest.NewRequest("DELETE", url, nil)

			if tc.statusCode != http.StatusBadRequest {
				req = mux.SetURLVars(req, map[string]string{"userID": tc.ID})
			}

			w := httptest.NewRecorder()

			userTestingHandler.deleteUser(w, req)

			require.Equal(t, tc.statusCode, w.Code)
		})
	}
}
