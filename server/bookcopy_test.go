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
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
)

func TestBookCopyCreate(t *testing.T) {
	initialBook := &book.Book{
		ID: util.NewID(),
	}

	initialBookCopy := &bookcopy.BookCopy{
		ID:     util.NewID(),
		BookID: initialBook.ID,
	}

	failedBookCopy := &bookcopy.BookCopy{
		ID:     util.NewID(),
		BookID: initialBook.ID,
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
			name:              "success creating a valid Book Copy",
			requestPayload:    initialBookCopy,
			mockReturnPayload: initialBookCopy,
			ID:                initialBookCopy.ID,
			statusCode:        http.StatusCreated,
			err:               nil,
		},
		{
			name:              "invalid Book Copy ID path",
			requestPayload:    initialBookCopy,
			mockReturnPayload: nil,
			ID:                "invalidBookCopyID",
			statusCode:        http.StatusBadRequest,
			err:               nil,
		},
		{
			name:              "invalid request payload",
			requestPayload:    "a plain string, not a Book Copy",
			mockReturnPayload: nil,
			ID:                "",
			statusCode:        http.StatusBadRequest,
			err:               nil,
		},
		{
			name:              "failed creating a Book Copy",
			requestPayload:    failedBookCopy,
			mockReturnPayload: nil,
			ID:                failedBookCopy.ID,
			statusCode:        http.StatusInternalServerError,
			err:               errors.New("Book Copy not found"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bookCopyService.On("Create", tc.requestPayload).Return(tc.mockReturnPayload, tc.err)

			url := fmt.Sprintf("/books/" + initialBook.ID + "/bookcopies")

			reqByte, err := json.Marshal(tc.requestPayload)
			require.Nil(t, err)

			req := httptest.NewRequest("POST", url, bytes.NewReader(reqByte))

			if tc.statusCode != http.StatusBadRequest {
				req = mux.SetURLVars(req, map[string]string{"bookID": initialBook.ID, "bookCopyID": tc.ID})
			}

			w := httptest.NewRecorder()

			bookCopyTestingHandler.createBookCopy(w, req)

			require.Equal(t, tc.statusCode, w.Code)
		})
	}
}

func TestBookCopyGet(t *testing.T) {
	initialBook := &book.Book{
		ID: util.NewID(),
	}

	initialBookCopy := &bookcopy.BookCopy{
		ID:     util.NewID(),
		BookID: initialBook.ID,
	}

	tt := []struct {
		name              string
		mockReturnPayload interface{}
		ID                string
		statusCode        int
		err               error
	}{
		{
			name:              "success retrieving a valid Book Copy",
			mockReturnPayload: initialBookCopy,
			ID:                initialBookCopy.ID,
			statusCode:        http.StatusOK,
			err:               nil,
		},
		{
			name:              "invalid Book Copy ID path",
			mockReturnPayload: nil,
			ID:                "invalidBookCopyID",
			statusCode:        http.StatusBadRequest,
			err:               nil,
		},
		{
			name:              "Book Copy doesn't exist",
			mockReturnPayload: nil,
			ID:                util.NewID(),
			statusCode:        http.StatusInternalServerError,
			err:               errors.New("Book Copy not found"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bookCopyService.On("Get", tc.ID).Return(tc.mockReturnPayload, tc.err)

			url := fmt.Sprintf("/books/" + initialBook.ID + "/bookcopies/" + tc.ID)

			req := httptest.NewRequest("GET", url, nil)

			if tc.statusCode != http.StatusBadRequest {
				req = mux.SetURLVars(req, map[string]string{"bookID": initialBook.ID, "bookCopyID": tc.ID})
			}

			w := httptest.NewRecorder()

			bookCopyTestingHandler.getBookCopy(w, req)

			require.Equal(t, tc.statusCode, w.Code)
		})
	}
}

func TestBookCopyUpdate(t *testing.T) {
	initialBook := &book.Book{
		ID: util.NewID(),
	}

	initialBookCopy := &bookcopy.BookCopy{
		ID:        util.NewID(),
		BookID:    initialBook.ID,
		Condition: "initial condition",
	}

	editedBookCopy := initialBookCopy
	editedBookCopy.Condition = "edited condition"

	failedBookCopy := &bookcopy.BookCopy{
		ID:     util.NewID(),
		BookID: initialBook.ID,
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
			name:              "success updating a valid Book Copy",
			requestPayload:    editedBookCopy,
			mockReturnPayload: editedBookCopy,
			ID:                initialBookCopy.ID,
			statusCode:        http.StatusOK,
			err:               nil,
		},
		{
			name:              "invalid Book Copy ID path",
			requestPayload:    editedBookCopy,
			mockReturnPayload: nil,
			ID:                "invalidBookCopyID",
			statusCode:        http.StatusBadRequest,
			err:               nil,
		},
		{
			name:              "request payload is not a Book Copy",
			requestPayload:    "a plain string",
			mockReturnPayload: nil,
			ID:                initialBookCopy.ID,
			statusCode:        http.StatusBadRequest,
			err:               nil,
		},
		{
			name:              "failed updating a Book Copy",
			requestPayload:    failedBookCopy,
			mockReturnPayload: nil,
			ID:                failedBookCopy.ID,
			statusCode:        http.StatusInternalServerError,
			err:               errors.New("Book copy not found"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bookCopyService.On("Update", tc.requestPayload).Return(tc.mockReturnPayload, tc.err)

			url := fmt.Sprintf("/books/" + initialBook.ID + "/bookcopies/" + tc.ID)

			reqByte, err := json.Marshal(tc.requestPayload)
			require.Nil(t, err)

			req := httptest.NewRequest("PUT", url, bytes.NewReader(reqByte))

			if tc.statusCode != http.StatusBadRequest {
				req = mux.SetURLVars(req, map[string]string{"bookID": initialBook.ID, "bookCopyID": tc.ID})
			}

			w := httptest.NewRecorder()

			bookCopyTestingHandler.updateBookCopy(w, req)

			if tc.statusCode != http.StatusOK {
				require.Equal(t, tc.statusCode, w.Code)
				return
			}

			// If we want to test the title, we need to assert the type of the interface{}

		})
	}
}

func TestBookCopyDelete(t *testing.T) {
	initialBook := &book.Book{
		ID: util.NewID(),
	}

	initialBookCopy := &bookcopy.BookCopy{
		ID:     util.NewID(),
		BookID: initialBook.ID,
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
			mockReturnPayload: initialBookCopy,
			ID:                initialBookCopy.ID,
			statusCode:        http.StatusOK,
			err:               nil,
		},
		{
			name:              "invalid Book Copy ID path",
			mockReturnPayload: nil,
			ID:                "invalidBookCopyID",
			statusCode:        http.StatusBadRequest,
			err:               nil,
		},
		{
			name:              "book copy doesn't exist",
			mockReturnPayload: nil,
			ID:                util.NewID(),
			statusCode:        http.StatusInternalServerError,
			err:               errors.New("Book copy not found"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bookCopyService.On("Delete", tc.ID).Return(tc.err)

			url := fmt.Sprintf("/books/" + initialBook.ID + "/bookcopies/" + tc.ID)

			req := httptest.NewRequest("DELETE", url, nil)

			if tc.statusCode != http.StatusBadRequest {
				req = mux.SetURLVars(req, map[string]string{"bookID": initialBook.ID, "bookCopyID": tc.ID})
			}

			w := httptest.NewRecorder()

			bookCopyTestingHandler.deleteBookCopy(w, req)

			require.Equal(t, tc.statusCode, w.Code)
		})
	}
}
