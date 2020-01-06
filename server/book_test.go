package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/stretchr/testify/require"

	"github.com/joshuabezaleel/library-server/pkg/auth"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
)

// var bookKRepository = &book.MockRepository{}
// var userService = &book.MockService{}

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
			jsonReq, _ := json.Marshal(tc.book)

			req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonReq))
			if err != nil {
				t.Errorf("Cannot perform HTTP request: %v", err)
			}

			rr := httptest.NewRecorder()
			srv.Router.ServeHTTP(rr, req)

			resp := rr.Result()
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				if resp.StatusCode != tc.statusCode {
					t.Errorf("expected %v; got %v", tc.statusCode, resp.Status)
				}
				return
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			createdBook := book.Book{}
			err = json.Unmarshal(body, &createdBook)
			if err != nil {
				t.Fatalf("expected a Book struct; got %s", body)
			}

			if createdBook.ID != tc.ID {
				t.Errorf("expected book id to be %v; got %v", createdBook.ID, tc.ID)
			}

		})
	}

	repository.CleanUp()
}

// func TestGetBook(t *testing.T) {
// 	initialBook := createBook()

// 	tt := []struct {
// 		name         string
// 		path         string
// 		statusCode   int
// 		errorMessage string
// 	}{
// 		{
// 			name:         "success retrieving a valid book",
// 			path:         "/" + initialBook.ID,
// 			statusCode:   200,
// 			errorMessage: "",
// 		},
// 		{
// 			name:         "invalid book id path",
// 			path:         "/" + util.NewID(),
// 			statusCode:   500,
// 			errorMessage: "",
// 		},
// 	}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			url := fmt.Sprintf("/books" + tc.path)

// 			req, err := http.NewRequest("GET", url, nil)
// 			if err != nil {
// 				t.Errorf("Cannot perform HTTP request: %v", err)
// 			}

// 			rr := httptest.NewRecorder()
// 			srv.Router.ServeHTTP(rr, req)

// 			resp := rr.Result()
// 			defer resp.Body.Close()

// 			if resp.StatusCode != http.StatusOK {
// 				if resp.StatusCode != tc.statusCode {
// 					t.Errorf("expected %v; got %v", tc.statusCode, resp.Status)
// 				}
// 				return
// 			}

// 			body, err := ioutil.ReadAll(resp.Body)
// 			if err != nil {
// 				t.Fatalf("could not read response: %v", err)
// 			}

// 			newBook := book.Book{}
// 			err = json.Unmarshal(body, &newBook)
// 			if err != nil {
// 				t.Fatalf("expected a Book struct; got %s", body)
// 			}

// 			if initialBook.ID != newBook.ID {
// 				t.Fatalf("expected id %v; got %v", initialBook.ID, newBook.ID)
// 			}
// 		})
// 	}

// 	repository.CleanUp()
// }

func TestUpdateBook(t *testing.T) {
	initialBook := createBook()
	// log.Println(initialBook.ID)

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
			jsonReq, _ := json.Marshal(tc.book)
			url := fmt.Sprintf("/books" + tc.path)
			// log.Println(tc.book)
			// log.Println(url)

			req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonReq))
			if err != nil {
				t.Errorf("Cannot perform HTTP request: %v", err)
			}

			rr := httptest.NewRecorder()
			srv.Router.ServeHTTP(rr, req)
			// log.Println(rr.Body.String())

			resp := rr.Result()
			// log.Println(resp.Body)
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				if resp.StatusCode != tc.statusCode {
					t.Errorf("expected %v; got %v", tc.statusCode, resp.Status)
				}
				return
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			updatedBook := book.Book{}
			err = json.Unmarshal(body, &updatedBook)
			if err != nil {
				t.Fatalf("expected a Book struct; got %s", body)
			}

			if updatedBook.Title != tc.title {
				t.Errorf("expected title to be %v; got %v", tc.title, updatedBook.Title)
			}
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

			req, err := http.NewRequest("DELETE", url, nil)
			if err != nil {
				t.Errorf("Cannot perform HTTP request: %v", err)
			}

			rr := httptest.NewRecorder()
			srv.Router.ServeHTTP(rr, req)

			resp := rr.Result()
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				if resp.StatusCode != tc.statusCode {
					t.Errorf("expected %v; got %v", tc.statusCode, resp.Status)
				}
				return
			}

			// body, err := ioutil.ReadAll(resp.Body)
			// if err != nil {
			// 	t.Fatalf("could not read response: %v", err)
			// }

			// msg := string(body)
			// if msg :=
			// if initialBook.ID != newBook.ID {
			// 	t.Fatalf("expected id %v; got %v", initialBook.ID, newBook.ID)
			// }
		})
	}

	repository.CleanUp()
}

func TestGetMock(t *testing.T) {
	bookService := &book.MockService{}
	authService := &auth.MockService{}

	bookHandler := bookHandler{bookService, authService}

	book := &book.Book{
		ID:    util.NewID(),
		Title: "title",
	}

	url := fmt.Sprintf("/books/" + "invalidbookID")
	req := httptest.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	bookHandler.getBook(w, req)
	require.Equal(t, http.StatusBadRequest, w.Code)

	bookService.On("Get", book.ID).Return(book, nil)
	req = httptest.NewRequest("GET", url, nil)
	req = mux.SetURLVars(req, map[string]string{"bookID": book.ID})
	w = httptest.NewRecorder()
	bookHandler.getBook(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	wrongID := util.NewID()
	bookService.On("Get", wrongID).Return(nil, errors.New("another error"))
	url = fmt.Sprintf("/books/" + wrongID)
	req = httptest.NewRequest("GET", url, nil)
	req = mux.SetURLVars(req, map[string]string{"bookID": wrongID})
	w = httptest.NewRecorder()
	bookHandler.getBook(w, req)
	require.Equal(t, http.StatusInternalServerError, w.Code)
}
