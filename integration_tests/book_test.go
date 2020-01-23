package integrationtests

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

func TestCreateBook(t *testing.T) {
	initialBook := &book.Book{
		ID: util.NewID(),
	}

	tt := []struct {
		name         string
		book         interface{}
		ID           string
		statusCode   int
		errorMessage string
	}{
		{
			name:         "success creating a book",
			book:         initialBook,
			statusCode:   201,
			ID:           initialBook.ID,
			errorMessage: "",
		},
		{
			name:         "request payload is not a book",
			book:         "definitely not a book, just a plain string",
			statusCode:   400,
			errorMessage: "",
		},
		{
			name: "made a book with same id (?), bad example though",
			book: &book.Book{
				ID: initialBook.ID,
			},
			statusCode:   500,
			errorMessage: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			jsonRequest, err := json.Marshal(tc.book)
			require.Nil(t, err)

			req := httptest.NewRequest("POST", "/books", bytes.NewBuffer(jsonRequest))

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

			createdBook := book.Book{}
			err = json.Unmarshal(body, &createdBook)
			require.Nil(t, err)

			require.Equal(t, tc.ID, createdBook.ID)
		})
	}

	repository.CleanUp()
}

func TestGetBook(t *testing.T) {
	initialBook := createBook()

	tt := []struct {
		name         string
		path         string
		statusCode   int
		errorMessage string
	}{
		{
			name:         "success retrieving a valid book",
			path:         "/" + initialBook.ID,
			statusCode:   200,
			errorMessage: "",
		},
		{
			name:         "invalid book id path",
			path:         "/" + util.NewID(),
			statusCode:   500,
			errorMessage: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("/books" + tc.path)

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

	repository.CleanUp()
}

func TestUpdateBook(t *testing.T) {
	initialBook := createBook()

	tt := []struct {
		name         string
		path         string
		book         interface{}
		title        string
		statusCode   int
		errorMessage string
	}{
		{
			name: "success updating a valid book",
			path: "/" + initialBook.ID,
			book: &book.Book{
				ID:    initialBook.ID,
				Title: "edited title",
			},
			title:        "edited title",
			statusCode:   200,
			errorMessage: "",
		},
		{
			name:         "request payload is not a book",
			path:         "/" + initialBook.ID,
			book:         "definitely not a book, just a plain string",
			statusCode:   400,
			errorMessage: "",
		},
		{
			name: "invalid book id path",
			path: "/" + util.NewID(),
			book: &book.Book{
				Title: "edited title",
			},
			statusCode:   500,
			errorMessage: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			jsonRequest, err := json.Marshal(tc.book)
			require.Nil(t, err)

			url := fmt.Sprintf("/books" + tc.path)

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

			updatedBook := book.Book{}
			err = json.Unmarshal(body, &updatedBook)
			require.Nil(t, err)

			require.Equal(t, tc.title, updatedBook.Title)
		})
	}

	repository.CleanUp()
}

func TestDeleteBook(t *testing.T) {
	initialBook := createBook()

	tt := []struct {
		name         string
		path         string
		statusCode   int
		errorMessage string
	}{
		{
			name:         "success deleting a valid book",
			path:         "/" + initialBook.ID,
			statusCode:   200,
			errorMessage: "",
		},
		{
			name:         "invalid book id path",
			path:         "/" + util.NewID(),
			statusCode:   500,
			errorMessage: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("/books" + tc.path)

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
