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
	"github.com/joshuabezaleel/library-server/pkg/core/book"
)

func TestBookCreate(t *testing.T) {
	initialBook := &book.Book{
		ID:    util.NewID(),
		Title: "title",
	}

	failedBook := &book.Book{
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
			name:              "success creating a valid Book",
			requestPayload:    initialBook,
			mockReturnPayload: initialBook,
			ID:                initialBook.ID,
			statusCode:        http.StatusCreated,
			err:               nil,
		},
		{
			name:              "invalid request payload",
			requestPayload:    "a plain string, not a Book",
			mockReturnPayload: nil,
			ID:                "",
			statusCode:        http.StatusBadRequest,
			err:               nil,
		},
		{
			name:              "failed creating a Book",
			requestPayload:    failedBook,
			mockReturnPayload: nil,
			ID:                failedBook.ID,
			statusCode:        http.StatusInternalServerError,
			err:               errors.New("Books not found"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bookService.On("Create", tc.requestPayload).Return(tc.mockReturnPayload, tc.err)

			url := fmt.Sprintf("/books")

			reqByte, err := json.Marshal(tc.requestPayload)
			require.Nil(t, err)

			req := httptest.NewRequest("POST", url, bytes.NewReader(reqByte))

			if tc.statusCode != http.StatusBadRequest {
				req = mux.SetURLVars(req, map[string]string{"bookID": tc.ID})
			}

			w := httptest.NewRecorder()

			bookTestingHandler.createBook(w, req)

			require.Equal(t, tc.statusCode, w.Code)
		})
	}
}

func TestBookGet(t *testing.T) {
	initialBook := &book.Book{
		ID:    util.NewID(),
		Title: "title",
	}

	tt := []struct {
		name              string
		mockReturnPayload interface{}
		ID                string
		statusCode        int
		err               error
	}{
		{
			name:              "success retrieving a valid book",
			mockReturnPayload: initialBook,
			ID:                initialBook.ID,
			statusCode:        http.StatusOK,
			err:               nil,
		},
		{
			name:              "invalid path",
			mockReturnPayload: nil,
			ID:                "invalidBookID",
			statusCode:        http.StatusBadRequest,
			err:               nil,
		},
		{
			name:              "book doesn't exist",
			mockReturnPayload: nil,
			ID:                util.NewID(),
			statusCode:        http.StatusInternalServerError,
			err:               errors.New("Books not found"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bookService.On("Get", tc.ID).Return(tc.mockReturnPayload, tc.err)

			url := fmt.Sprintf("/books/" + tc.ID)

			req := httptest.NewRequest("GET", url, nil)

			if tc.statusCode != http.StatusBadRequest {
				req = mux.SetURLVars(req, map[string]string{"bookID": tc.ID})
			}

			w := httptest.NewRecorder()

			bookTestingHandler.getBook(w, req)

			require.Equal(t, tc.statusCode, w.Code)
		})
	}
}

func TestBookUpdate(t *testing.T) {
	initialBook := &book.Book{
		ID:    util.NewID(),
		Title: "title",
	}

	editedBook := initialBook
	editedBook.Title = "edited title"

	failedBook := &book.Book{
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
			name:              "success updating a valid Book",
			requestPayload:    editedBook,
			mockReturnPayload: editedBook,
			ID:                initialBook.ID,
			statusCode:        http.StatusOK,
			err:               nil,
		},
		{
			name:              "invalid Book ID path",
			requestPayload:    editedBook,
			mockReturnPayload: nil,
			ID:                "invalidBookID",
			statusCode:        http.StatusBadRequest,
			err:               nil,
		},
		{
			name:              "request payload is not a Book",
			requestPayload:    "a plain string",
			mockReturnPayload: nil,
			ID:                initialBook.ID,
			statusCode:        http.StatusBadRequest,
			err:               nil,
		},
		{
			name:              "failed updating a Book",
			requestPayload:    failedBook,
			mockReturnPayload: nil,
			ID:                failedBook.ID,
			statusCode:        http.StatusInternalServerError,
			err:               errors.New("Books not found"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bookService.On("Update", tc.requestPayload).Return(tc.mockReturnPayload, tc.err)

			url := fmt.Sprintf("/books/" + tc.ID)

			reqByte, err := json.Marshal(tc.requestPayload)
			require.Nil(t, err)

			req := httptest.NewRequest("PUT", url, bytes.NewReader(reqByte))

			if tc.statusCode != http.StatusBadRequest {
				req = mux.SetURLVars(req, map[string]string{"bookID": tc.ID})
			}

			w := httptest.NewRecorder()

			bookTestingHandler.updateBook(w, req)

			if tc.statusCode != http.StatusOK {
				require.Equal(t, tc.statusCode, w.Code)
				return
			}

			// If we want to test the title, we need to assert the type of the interface{}

		})
	}
}

func TestBookDelete(t *testing.T) {
	initialBook := &book.Book{
		ID:    util.NewID(),
		Title: "title",
	}

	tt := []struct {
		name              string
		mockReturnPayload interface{}
		ID                string
		statusCode        int
		err               error
	}{
		{
			name:              "success deleting a valid book",
			mockReturnPayload: initialBook,
			ID:                initialBook.ID,
			statusCode:        http.StatusOK,
			err:               nil,
		},
		{
			name:              "invalid path",
			mockReturnPayload: nil,
			ID:                "invalidBookID",
			statusCode:        http.StatusBadRequest,
			err:               nil,
		},
		{
			name:              "book doesn't exist",
			mockReturnPayload: nil,
			ID:                util.NewID(),
			statusCode:        http.StatusInternalServerError,
			err:               errors.New("Books not found"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bookService.On("Delete", tc.ID).Return(tc.err)

			url := fmt.Sprintf("/books/" + tc.ID)

			req := httptest.NewRequest("DELETE", url, nil)

			if tc.statusCode != http.StatusBadRequest {
				req = mux.SetURLVars(req, map[string]string{"bookID": tc.ID})
			}

			w := httptest.NewRecorder()

			bookTestingHandler.deleteBook(w, req)

			require.Equal(t, tc.statusCode, w.Code)
		})
	}
}
