package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
)

func createBook() *book.Book {
	initialBook := &book.Book{
		ID:    util.NewID(),
		Title: "initial title",
	}

	jsonReq, _ := json.Marshal(initialBook)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonReq))
	rr := httptest.NewRecorder()
	srv.Router.ServeHTTP(rr, req)

	return initialBook
}

func createBookCopy() *bookcopy.BookCopy {
	initialBook := &book.Book{
		ID:    util.NewID(),
		Title: "initial title",
	}

	jsonReq, _ := json.Marshal(initialBook)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonReq))
	rr := httptest.NewRecorder()
	srv.Router.ServeHTTP(rr, req)

	initialBookCopy := &bookcopy.BookCopy{
		ID:        util.NewID(),
		BookID:    initialBook.ID,
		Condition: "initial condition",
	}

	jsonReq, _ = json.Marshal(initialBookCopy)
	req, _ = http.NewRequest("POST", "/books/"+initialBook.ID+"/bookcopies", bytes.NewBuffer(jsonReq))
	rr = httptest.NewRecorder()
	srv.Router.ServeHTTP(rr, req)

	return initialBookCopy
}
