package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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
	book := &book.Book{
		ID:    util.NewID(),
		Title: "TITLE WOOOOOI 2",
	}
	log.Printf("BOOK ID = %v", book.ID)
	jsonReq, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonReq))
	rr := httptest.NewRecorder()
	srv.Router.ServeHTTP(rr, req)

	log.Println(book.ID)
	log.Println(rr.Code)

	tt := []struct {
		path         string
		statusCode   int
		errorMessage string
	}{
		{
			path:         "/" + book.ID,
			statusCode:   200,
			errorMessage: "",
		},
		{
			path:         "/" + util.NewID(),
			statusCode:   500,
			errorMessage: "",
		},
	}

	for _, tc := range tt {
		// jsonReq, _ := json.Marshal(tc.book)
		url := fmt.Sprintf("/books%s", tc.path)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Errorf("Cannot perform HTTP request: %v", err)
		}
		log.Println(req)

		rr := httptest.NewRecorder()
		srv.Router.ServeHTTP(rr, req)

		if rr.Code != tc.statusCode {
			t.Errorf("rec.Code = %d; want = %d", rr.Code, tc.statusCode)
		}
		log.Println(rr.Body)
	}

	// repository.CleanUp()
}
