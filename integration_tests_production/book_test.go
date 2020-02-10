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
	"github.com/joshuabezaleel/library-server/pkg/core/book"
)

func TestBookCreate(t *testing.T) {
	initialBook := &book.Book{
		ID: util.NewID(),
	}

	tt := []struct {
		name               string
		ID                 string
		requestPayload     interface{}
		authorizationToken string
		statusCode         int
		errorMessage       string
	}{
		{
			name:               "success creating a Book by User Librarian",
			ID:                 initialBook.ID,
			requestPayload:     initialBook,
			authorizationToken: userLibrarianToken,
			statusCode:         http.StatusCreated,
			errorMessage:       "",
		},
		{
			name:               "failed authorization when creating a Book by User Student",
			ID:                 initialBook.ID,
			requestPayload:     initialBook,
			authorizationToken: userStudentToken,
			statusCode:         http.StatusUnauthorized,
			errorMessage:       "",
		},
		{
			name:               "request payload is not a Book",
			ID:                 "",
			requestPayload:     "plain string, not a Book",
			authorizationToken: userLibrarianToken,
			statusCode:         http.StatusBadRequest,
			errorMessage:       "",
		},
		{
			name:               "made a Book with the same ID",
			ID:                 initialBook.ID,
			requestPayload:     initialBook,
			authorizationToken: userLibrarianToken,
			statusCode:         http.StatusInternalServerError,
			errorMessage:       "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			jsonRequest, err := json.Marshal(tc.requestPayload)
			require.Nil(t, err)

			req := httptest.NewRequest("POST", "/books", bytes.NewBuffer(jsonRequest))
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

			createdBook := &book.Book{}
			err = json.Unmarshal(body, createdBook)
			require.Nil(t, err)

			require.Equal(t, tc, createdBook.ID)
		})
	}
}

func TestBookGet(t *testing.T) {
	initialBook := createBook()

	tt := []struct {
		name         string
		ID           string
		statusCode   int
		errorMessage string
	}{
		{
			name:         "success retrieving a valid Book",
			ID:           initialBook.ID,
			statusCode:   http.StatusOK,
			errorMessage: "",
		},
		{
			name:         "invalid Book ID path",
			ID:           util.NewID(),
			statusCode:   http.StatusInternalServerError,
			errorMessage: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("/books/" + tc.ID)

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

			newBook := book.Book{}
			err = json.Unmarshal(body, &newBook)
			require.Nil(t, err)

			require.Equal(t, initialBook.ID, newBook.ID)
		})
	}
}

func TestBookUpdate(t *testing.T) {
	initialBook := createBook()

	tt := []struct {
		name               string
		ID                 string
		requestPayload     interface{}
		authorizationToken string
		expectedTitle      string
		statusCode         int
		errorMessage       string
	}{
		{
			name: "success updating a valid Book by User Librarian",
			ID:   initialBook.ID,
			requestPayload: &book.Book{
				ID:    initialBook.ID,
				Title: "edited title",
			},
			authorizationToken: userLibrarianToken,
			expectedTitle:      "edited title",
			statusCode:         http.StatusOK,
			errorMessage:       "",
		},
		{
			name:               "request payload is not a Book",
			ID:                 initialBook.ID,
			requestPayload:     "definitely not a Book, just a plain string",
			authorizationToken: userLibrarianToken,
			statusCode:         http.StatusBadRequest,
			errorMessage:       "",
		},
		{
			name:               "failed authorization when updating a Book by User Student",
			ID:                 initialBook.ID,
			requestPayload:     "definitely not a Book, just a plain string",
			authorizationToken: userStudentToken,
			statusCode:         http.StatusUnauthorized,
			errorMessage:       "",
		},
		{
			name: "invalid Book ID path",
			ID:   util.NewID(),
			requestPayload: &book.Book{
				Title: "edited title",
			},
			authorizationToken: userLibrarianToken,
			statusCode:         http.StatusInternalServerError,
			errorMessage:       "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			jsonRequest, err := json.Marshal(tc.requestPayload)
			require.Nil(t, err)

			url := fmt.Sprintf("/books/" + tc.ID)

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

			updatedBook := book.Book{}
			err = json.Unmarshal(body, &updatedBook)
			require.Nil(t, err)

			require.Equal(t, tc.expectedTitle, updatedBook.Title)
		})
	}
}

func TestBookDelete(t *testing.T) {
	initialBook := createBook()

	tt := []struct {
		name               string
		ID                 string
		authorizationToken string
		statusCode         int
		errorMessage       string
	}{
		{
			name:               "success deleting a valid Book by User Librarian",
			ID:                 initialBook.ID,
			authorizationToken: userLibrarianToken,
			statusCode:         http.StatusOK,
			errorMessage:       "",
		},
		{
			name:               "failed authorization on deleting a Book by User Student",
			ID:                 initialBook.ID,
			authorizationToken: userStudentToken,
			statusCode:         http.StatusUnauthorized,
			errorMessage:       "",
		},
		{
			name:               "invalid Book ID path",
			ID:                 util.NewID(),
			authorizationToken: userLibrarianToken,
			statusCode:         http.StatusInternalServerError,
			errorMessage:       "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("/books/" + tc.ID)

			req := httptest.NewRequest("DELETE", url, nil)
			req.Header.Add("Authorization", tc.authorizationToken)

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
}
