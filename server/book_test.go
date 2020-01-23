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

func TestGetMock(t *testing.T) {
	book := &book.Book{
		ID:    util.NewID(),
		Title: "title",
	}

	url := fmt.Sprintf("/books/" + "invalidbookID")
	req := httptest.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	bookTestingHandler.getBook(w, req)
	require.Equal(t, http.StatusBadRequest, w.Code)

	bookService.On("Get", book.ID).Return(book, nil)
	req = httptest.NewRequest("GET", url, nil)
	req = mux.SetURLVars(req, map[string]string{"bookID": book.ID})
	w = httptest.NewRecorder()
	bookTestingHandler.getBook(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	wrongID := util.NewID()
	bookService.On("Get", wrongID).Return(nil, errors.New("another error"))
	url = fmt.Sprintf("/books/" + wrongID)
	req = httptest.NewRequest("GET", url, nil)
	req = mux.SetURLVars(req, map[string]string{"bookID": wrongID})
	w = httptest.NewRecorder()
	bookTestingHandler.getBook(w, req)
	require.Equal(t, http.StatusInternalServerError, w.Code)
}
