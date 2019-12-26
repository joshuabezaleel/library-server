package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	util "github.com/joshuabezaleel/library-server/pkg"
	book "github.com/joshuabezaleel/library-server/pkg/core/book"
)

func TestCreateBook(t *testing.T) {
	tt := []struct {
		book         interface{}
		statusCode   int
		errorMessage string
	}{
		{
			book: &book.Book{
				ID: util.NewID(),
			},
			statusCode:   201,
			errorMessage: "",
		},
		{
			book:         "Definitely not a book, just a plain string",
			statusCode:   400,
			errorMessage: "",
		},
		{
			book:         &book.Book{},
			statusCode:   500,
			errorMessage: "",
		},
	}

	for _, tc := range tt {
		jsonReq, _ := json.Marshal(tc.book)

		req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonReq))
		if err != nil {
			t.Errorf("Cannot perform HTTP request: %v", err)
		}

		rr := httptest.NewRecorder()
		srv.Router.ServeHTTP(rr, req)

		if rr.Code != tc.statusCode {
			t.Errorf("rec.Code = %d; want = %d", rr.Code, tc.statusCode)
		}
	}

	repository.CleanUp()
}
func TestGetBook(t *testing.T) {
}
