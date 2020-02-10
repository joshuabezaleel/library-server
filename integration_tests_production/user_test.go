package integrationtestsproduction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
)

func TestCreateUser(t *testing.T) {
	initialUser := &user.User{
		ID: util.NewID(),
	}

	tt := []struct {
		name           string
		ID             string
		requestPayload interface{}
		statusCode     int
		errorMessage   string
	}{
		{
			name:           "success creating a User",
			ID:             initialUser.ID,
			requestPayload: initialUser,
			statusCode:     http.StatusCreated,
			errorMessage:   "",
		},
		{
			name:           "request payload is not a User",
			ID:             initialUser.ID,
			requestPayload: "definitely not a User, just a plain string",
			statusCode:     http.StatusBadRequest,
			errorMessage:   "",
		},
		{
			name: "made a User ID",
			ID:   initialUser.ID,
			requestPayload: &user.User{
				ID: initialUser.ID,
			},
			statusCode:   http.StatusInternalServerError,
			errorMessage: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			jsonReq, err := json.Marshal(tc.requestPayload)
			require.Nil(t, err)

			req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(jsonReq))

			rr := httptest.NewRecorder()
			srv.Router.ServeHTTP(rr, req)

			resp := rr.Result()
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				require.Equal(t, tc.statusCode, resp.StatusCode)
				return
			}

			body, err := ioutil.ReadAll(resp.Body)
			require.Nil(t, err)

			createdUser := user.User{}
			err = json.Unmarshal(body, &createdUser)
			require.Nil(t, err)

			require.Equal(t, tc.ID, createdUser.ID)
		})
	}
}

func TestGetUser(t *testing.T) {
	initialUser := userLibrarian

	tt := []struct {
		name         string
		ID           string
		statusCode   int
		errorMessage string
	}{
		{
			name:         "success retrieving a valid User",
			ID:           initialUser.ID,
			statusCode:   http.StatusOK,
			errorMessage: "",
		},
		{
			name:         "invalid User ID path",
			ID:           util.NewID(),
			statusCode:   http.StatusInternalServerError,
			errorMessage: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("/users/" + tc.ID)

			req := httptest.NewRequest("GET", url, nil)

			rr := httptest.NewRecorder()
			srv.Router.ServeHTTP(rr, req)

			resp := rr.Result()
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				require.Equal(t, tc.statusCode, resp.StatusCode)
				return
			}

			body, err := ioutil.ReadAll(resp.Body)
			require.Nil(t, err)

			newUser := user.User{}
			err = json.Unmarshal(body, &newUser)
			require.Nil(t, err)

			require.Equal(t, initialUser.ID, newUser.ID)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	librarian := userLibrarian
	librarianAuthorizationToken := "Bearer " + login(librarian)

	student := userStudent
	studentAuthorizationToken := "Bearer " + login(student)

	tt := []struct {
		name               string
		ID                 string
		requestPayload     interface{}
		authorizationToken string
		expectedUsername   string
		statusCode         int
		errorMessage       string
	}{
		{
			name: "success updating a valid User",
			ID:   librarian.ID,
			requestPayload: &user.User{
				Username: "edited username",
			},
			authorizationToken: librarianAuthorizationToken,
			expectedUsername:   "edited username",
			statusCode:         http.StatusOK,
			errorMessage:       "",
		},
		{
			name:               "request payload is not a User",
			ID:                 student.ID,
			requestPayload:     "definitely not a User, just a plain string",
			authorizationToken: studentAuthorizationToken,
			statusCode:         http.StatusBadRequest,
			errorMessage:       "",
		},
		{
			name: "invalid User ID path",
			ID:   util.NewID(),
			requestPayload: &user.User{
				Username: "edited username",
			},
			authorizationToken: librarianAuthorizationToken,
			statusCode:         http.StatusInternalServerError,
			errorMessage:       "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			jsonRequest, err := json.Marshal(tc.requestPayload)
			require.Nil(t, err)

			url := fmt.Sprintf("/users/" + tc.ID)

			req := httptest.NewRequest("PUT", url, bytes.NewBuffer(jsonRequest))
			req.Header.Add("Authorization", tc.authorizationToken)

			rr := httptest.NewRecorder()
			srv.Router.ServeHTTP(rr, req)

			resp := rr.Result()
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				require.Equal(t, tc.statusCode, resp.StatusCode)
				return
			}

			body, err := ioutil.ReadAll(resp.Body)
			require.Nil(t, err)

			updatedUser := user.User{}
			err = json.Unmarshal(body, &updatedUser)
			require.Nil(t, err)

			require.Equal(t, tc.expectedUsername, updatedUser.Username)
		})
	}
}
