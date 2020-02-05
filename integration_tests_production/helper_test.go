package integrationtestsproduction

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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

func createUserStudent() *user.User {
	initialUserStudent := &user.User{
		ID:       util.NewID(),
		Username: "student",
		Password: "student",
		Role:     "student",
	}

	jsonReq, _ := json.Marshal(initialUserStudent)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonReq))
	rr := httptest.NewRecorder()
	srv.Router.ServeHTTP(rr, req)

	return initialUserStudent
}

func createUserLibrarian() *user.User {
	initialUserLibrarian := &user.User{
		ID:       util.NewID(),
		Username: "librarian",
		Password: "librarian",
		Role:     "librarian",
	}

	jsonReq, _ := json.Marshal(initialUserLibrarian)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonReq))
	rr := httptest.NewRecorder()
	srv.Router.ServeHTTP(rr, req)

	return initialUserLibrarian
}

// login accepts a particular User with its role and return the token string for authorization purposes.
func login(user *user.User) string {
	jsonReq, err := json.Marshal(user)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonReq))
	if err != nil {
		panic(err)
	}

	rr := httptest.NewRecorder()
	srv.Router.ServeHTTP(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	token := string(body)
	token = token[1 : len(token)-1]

	return token
}
