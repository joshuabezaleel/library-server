package server

import (
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

func TestBookGet(t *testing.T) {
	initialBook := &book.Book{
		ID:    util.NewID(),
		Title: "title",
	}

	tt := []struct {
		name              string
		returnMockPayload interface{}
		ID                string
		statusCode        int
		err               error
	}{
		{
			name:              "success retrieving a valid book",
			returnMockPayload: initialBook,
			ID:                initialBook.ID,
			statusCode:        http.StatusOK,
			err:               nil,
		},
		{
			name:              "invalid path",
			returnMockPayload: nil,
			ID:                "invalidBookID",
			statusCode:        http.StatusBadRequest,
			err:               nil,
		},
		{
			name:              "book doesn't exist",
			returnMockPayload: nil,
			ID:                util.NewID(),
			statusCode:        http.StatusInternalServerError,
			err:               errors.New("Books not found"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bookService.On("Get", tc.ID).Return(tc.returnMockPayload, tc.err)

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
