package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
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

	// log.Println("ASDFGHJKL")
	// log.Println(rr.Body.String())
	// log.Println("ASDFGHJKL")

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

func createUser() *user.User {
	initialUser := &user.User{
		ID:       util.NewID(),
		Username: "initial username",
	}

	jsonReq, _ := json.Marshal(initialUser)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonReq))
	rr := httptest.NewRecorder()
	srv.Router.ServeHTTP(rr, req)

	return initialUser
}
