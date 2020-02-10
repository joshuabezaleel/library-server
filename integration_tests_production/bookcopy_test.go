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
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
)

func TestBookCopyCreate(t *testing.T) {
	initialBook := createBook()

	initialBookCopy := &bookcopy.BookCopy{
		ID:     util.NewID(),
		BookID: initialBook.ID,
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
			name:               "success creating a BookCopy by User Librarian",
			ID:                 initialBookCopy.ID,
			requestPayload:     initialBookCopy,
			authorizationToken: userLibrarianToken,
			statusCode:         http.StatusCreated,
			errorMessage:       "",
		},
		{
			name:               "failed authorization on creating a BookCopy by User Student",
			ID:                 initialBookCopy.ID,
			requestPayload:     initialBookCopy,
			authorizationToken: userStudentToken,
			statusCode:         http.StatusUnauthorized,
			errorMessage:       "",
		},
		{
			name:               "request payload is not a BookCopy",
			ID:                 initialBookCopy.ID,
			requestPayload:     "definitely not a Book Copy, just a plain string",
			authorizationToken: userLibrarianToken,
			statusCode:         http.StatusBadRequest,
			errorMessage:       "",
		},
		{
			name:               "made a Book Copy with the same ID",
			ID:                 initialBookCopy.ID,
			requestPayload:     initialBookCopy,
			authorizationToken: userLibrarianToken,
			statusCode:         http.StatusInternalServerError,
			errorMessage:       "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			jsonReq, err := json.Marshal(tc.requestPayload)
			require.Nil(t, err)

			req := httptest.NewRequest("POST", "/books/"+initialBook.ID+"/bookcopies", bytes.NewBuffer(jsonReq))
			req.Header.Add("Authorization", tc.authorizationToken)
			require.Nil(t, err)

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

			createdBookCopy := book.Book{}
			err = json.Unmarshal(body, &createdBookCopy)
			require.Nil(t, err)

			require.Equal(t, tc.ID, createdBookCopy.ID)
		})
	}
}

func TestBookCopyGet(t *testing.T) {
	initialBookCopy := createBookCopy()

	tt := []struct {
		name         string
		ID           string
		statusCode   int
		errorMessage string
	}{
		{
			name:         "success retrieving a valid Book Copy",
			ID:           initialBookCopy.ID,
			statusCode:   http.StatusOK,
			errorMessage: "",
		},
		{
			name:         "invalid Book Copy ID path",
			ID:           util.NewID(),
			statusCode:   http.StatusInternalServerError,
			errorMessage: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("/books/" + initialBookCopy.BookID + "/bookcopies/" + tc.ID)

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

			newBookCopy := bookcopy.BookCopy{}
			err = json.Unmarshal(body, &newBookCopy)
			require.Nil(t, err)

			require.Equal(t, initialBookCopy.ID, newBookCopy.ID)
		})
	}
}

func TestBookCopyUpdate(t *testing.T) {
	initialBookCopy := createBookCopy()

	tt := []struct {
		name               string
		ID                 string
		requestPayload     interface{}
		authorizationToken string
		expectedCondition  string
		statusCode         int
		errorMessage       string
	}{
		{
			name: "success updating a valid Book Copy by User Librarian",
			ID:   initialBookCopy.ID,
			requestPayload: &bookcopy.BookCopy{
				Condition: "revised condition",
			},
			authorizationToken: userLibrarianToken,
			expectedCondition:  "revised condition",
			statusCode:         http.StatusOK,
			errorMessage:       "",
		},
		{
			name: "failed authorization on updating a valid Book Copy by User Student",
			ID:   initialBookCopy.ID,
			requestPayload: &bookcopy.BookCopy{
				Condition: "revised condition",
			},
			authorizationToken: userStudentToken,
			expectedCondition:  "revised condition",
			statusCode:         http.StatusUnauthorized,
			errorMessage:       "",
		},
		{
			name:               "request payload is not a Book Copy",
			ID:                 initialBookCopy.ID,
			requestPayload:     "definitely not a Book Copy, just a plain string",
			authorizationToken: userLibrarianToken,
			statusCode:         http.StatusBadRequest,
			errorMessage:       "",
		},
		{
			name: "invalid Book Copy ID path",
			ID:   util.NewID(),
			requestPayload: &bookcopy.BookCopy{
				Condition: "revised condition",
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

			url := fmt.Sprintf("/books/" + initialBookCopy.BookID + "/bookcopies/" + tc.ID)

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

			updatedBookCopy := bookcopy.BookCopy{}
			err = json.Unmarshal(body, &updatedBookCopy)
			require.Nil(t, err)

			require.Equal(t, tc.expectedCondition, updatedBookCopy.Condition)
		})
	}
}

func TestBookCopyDelete(t *testing.T) {
	initialBookCopy := createBookCopy()

	tt := []struct {
		name               string
		ID                 string
		authorizationToken string
		statusCode         int
		errorMessage       string
	}{
		{
			name:               "success deleting a valid Book Copy by User Librarian",
			ID:                 initialBookCopy.ID,
			authorizationToken: userLibrarianToken,
			statusCode:         http.StatusOK,
			errorMessage:       "",
		},
		{
			name:               "failed authorization on deleting a valid Book Copy by User Student",
			ID:                 initialBookCopy.ID,
			authorizationToken: userStudentToken,
			statusCode:         http.StatusUnauthorized,
			errorMessage:       "",
		},
		{
			name:               "invalid Book Copy ID path",
			ID:                 util.NewID(),
			authorizationToken: userLibrarianToken,
			statusCode:         http.StatusInternalServerError,
			errorMessage:       "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("/books/" + initialBookCopy.BookID + "/bookcopies/" + tc.ID)

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
