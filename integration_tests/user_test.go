package integrationtests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	initialUser := &user.User{
		ID: util.NewID(),
	}

	tt := []struct {
		name         string
		user         interface{}
		ID           string
		statusCode   int
		errorMessage string
	}{
		{
			name:         "success creating a user",
			user:         initialUser,
			statusCode:   201,
			ID:           initialUser.ID,
			errorMessage: "",
		},
		{
			name:         "request payload is not a user",
			user:         "definitely not a user, just a plain string",
			statusCode:   400,
			errorMessage: "",
		},
		{
			name: "made a user with same id (?), bad example though",
			user: &user.User{
				ID: initialUser.ID,
			},
			statusCode:   500,
			errorMessage: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			jsonReq, err := json.Marshal(tc.user)
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

	repository.CleanUp()
}

func TestGetUser(t *testing.T) {
	initialUser := createUser()

	tt := []struct {
		name         string
		path         string
		statusCode   int
		errorMessage string
	}{
		{
			name:         "success retrieving a valid user",
			path:         "/" + initialUser.ID,
			statusCode:   200,
			errorMessage: "",
		},
		{
			name:         "invalid user id path",
			path:         "/" + util.NewID(),
			statusCode:   500,
			errorMessage: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("/users" + tc.path)

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

	repository.CleanUp()
}

func TestUpdateUser(t *testing.T) {
	initialUser := createUser()

	tt := []struct {
		name         string
		path         string
		user         interface{}
		username     string
		statusCode   int
		errorMessage string
	}{
		{
			name: "success updating a valid user",
			path: "/" + initialUser.ID,
			user: &user.User{
				Username: "edited username",
			},
			username:     "edited username",
			statusCode:   200,
			errorMessage: "",
		},
		{
			name:         "request payload is not a user",
			path:         "/" + initialUser.ID,
			user:         "definitely not a user, just a plain string",
			statusCode:   400,
			errorMessage: "",
		},
		{
			name: "invalid user id path",
			path: "/" + util.NewID(),
			user: &user.User{
				Username: "edited username",
			},
			statusCode:   500,
			errorMessage: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			jsonRequest, err := json.Marshal(tc.user)
			require.Nil(t, err)

			url := fmt.Sprintf("/users" + tc.path)

			req := httptest.NewRequest("PUT", url, bytes.NewBuffer(jsonRequest))

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

			require.Equal(t, tc.username, updatedUser.Username)
		})
	}

	repository.CleanUp()
}

func TestDeleteUser(t *testing.T) {
	initialUser := createUser()

	tt := []struct {
		name         string
		path         string
		statusCode   int
		errorMessage string
	}{
		{
			name:         "success deleting a valid user",
			path:         "/" + initialUser.ID,
			statusCode:   200,
			errorMessage: "",
		},
		{
			name:         "invalid user id path",
			path:         "/" + util.NewID(),
			statusCode:   500,
			errorMessage: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("/users" + tc.path)

			req := httptest.NewRequest("DELETE", url, nil)

			rr := httptest.NewRecorder()
			srv.Router.ServeHTTP(rr, req)

			resp := rr.Result()
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				require.Equal(t, tc.statusCode, resp.StatusCode)
				return
			}
		})
	}

	repository.CleanUp()
}
