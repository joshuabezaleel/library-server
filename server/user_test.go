package server

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
			jsonReq, _ := json.Marshal(tc.user)

			req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonReq))
			if err != nil {
				t.Errorf("Cannot perform HTTP request: %v", err)
			}

			rr := httptest.NewRecorder()
			srv.Router.ServeHTTP(rr, req)

			resp := rr.Result()
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				if resp.StatusCode != tc.statusCode {
					t.Errorf("expected %v; got %v", tc.statusCode, resp.Status)
				}
				return
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			createdUser := user.User{}
			err = json.Unmarshal(body, &createdUser)
			if err != nil {
				t.Fatalf("expected a User struct; got %s", body)
			}

			if createdUser.ID != tc.ID {
				t.Errorf("expected user id to be %v; got %v", createdUser.ID, tc.ID)
			}

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

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				t.Errorf("Cannot perform HTTP request: %v", err)
			}

			rr := httptest.NewRecorder()
			srv.Router.ServeHTTP(rr, req)

			resp := rr.Result()
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				if resp.StatusCode != tc.statusCode {
					t.Errorf("expected %v; got %v", tc.statusCode, resp.Status)
				}
				return
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			newUser := user.User{}
			err = json.Unmarshal(body, &newUser)
			if err != nil {
				t.Fatalf("expected a User struct; got %s", body)
			}

			if initialUser.ID != newUser.ID {
				t.Fatalf("expected id %v; got %v", initialUser.ID, newUser.ID)
			}
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
			jsonReq, _ := json.Marshal(tc.user)
			url := fmt.Sprintf("/users" + tc.path)

			req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonReq))
			if err != nil {
				t.Errorf("Cannot perform HTTP request: %v", err)
			}

			rr := httptest.NewRecorder()
			srv.Router.ServeHTTP(rr, req)

			resp := rr.Result()
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				if resp.StatusCode != tc.statusCode {
					t.Errorf("expected %v; got %v", tc.statusCode, resp.Status)
				}
				return
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			updatedUser := user.User{}
			err = json.Unmarshal(body, &updatedUser)
			if err != nil {
				t.Fatalf("expected a User struct; got %s", body)
			}

			if updatedUser.Username != tc.username {
				t.Errorf("expected username to be %v; got %v", tc.username, updatedUser.Username)
			}
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

			req, err := http.NewRequest("DELETE", url, nil)
			if err != nil {
				t.Errorf("Cannot perform HTTP request: %v", err)
			}

			rr := httptest.NewRecorder()
			srv.Router.ServeHTTP(rr, req)

			resp := rr.Result()
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				if resp.StatusCode != tc.statusCode {
					t.Errorf("expected %v; got %v", tc.statusCode, resp.Status)
				}
				return
			}

			// body, err := ioutil.ReadAll(resp.Body)
			// if err != nil {
			// 	t.Fatalf("could not read response: %v", err)
			// }

			// msg := string(body)
			// if msg :=
			// if initialBook.ID != newBook.ID {
			// 	t.Fatalf("expected id %v; got %v", initialBook.ID, newBook.ID)
			// }
		})
	}

	repository.CleanUp()
}
