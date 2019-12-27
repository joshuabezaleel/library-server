package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
)

func TestCreateBookCopy(t *testing.T) {
	initialBook := createBook()

	initialBookCopy := &bookcopy.BookCopy{
		ID:     util.NewID(),
		BookID: initialBook.ID,
	}

	tt := []struct {
		name         string
		bookCopy     interface{}
		ID           string
		statusCode   int
		errorMessage string
	}{
		{
			name:         "success creating a book copy",
			bookCopy:     initialBookCopy,
			ID:           initialBookCopy.ID,
			statusCode:   201,
			errorMessage: "",
		},
		{
			name:         "request payload is not a book copy",
			bookCopy:     "definitely not a book copy, just a plain string",
			statusCode:   400,
			errorMessage: "",
		},
		{
			name:         "book copy request is not complete",
			bookCopy:     &bookcopy.BookCopy{},
			statusCode:   500,
			errorMessage: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			jsonReq, _ := json.Marshal(tc.bookCopy)

			req, err := http.NewRequest("POST", "/books/"+initialBook.ID+"/bookcopies", bytes.NewBuffer(jsonReq))
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

			createdBookCopy := book.Book{}
			err = json.Unmarshal(body, &createdBookCopy)
			if err != nil {
				t.Fatalf("expected a BookCopy struct; got %s", body)
			}

			if createdBookCopy.ID != tc.ID {
				t.Errorf("expected book copy id to be %v; got %v", createdBookCopy.ID, tc.ID)
			}

		})
	}

	repository.CleanUp()
}

func TestGetBookCopy(t *testing.T) {
	initialBookCopy := createBookCopy()

	tt := []struct {
		name         string
		path         string
		statusCode   int
		errorMessage string
	}{
		{
			name:         "success retrieving a valid book copy",
			path:         initialBookCopy.ID,
			statusCode:   200,
			errorMessage: "",
		},
		{
			name:         "invalid book id path",
			path:         util.NewID(),
			statusCode:   500,
			errorMessage: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("/books/" + initialBookCopy.BookID + "/bookcopies/" + tc.path)

			req, err := http.NewRequest("GET", url, nil)
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

			newBookCopy := bookcopy.BookCopy{}
			err = json.Unmarshal(body, &newBookCopy)
			if err != nil {
				t.Fatalf("expected a Book Copy struct; got %s", body)
			}

			if initialBookCopy.ID != newBookCopy.ID {
				t.Fatalf("expected id %v; got %v", initialBookCopy.ID, newBookCopy.ID)
			}
		})
	}

	repository.CleanUp()
}

func TestUpdateBookCopy(t *testing.T) {
	initialBookCopy := createBookCopy()

	tt := []struct {
		name         string
		path         string
		bookCopy     interface{}
		condition    string
		statusCode   int
		errorMessage string
	}{
		{
			name: "success updating a valid book copy",
			path: initialBookCopy.ID,
			bookCopy: &bookcopy.BookCopy{
				Condition: "revised condition",
			},
			condition:    "revised condition",
			statusCode:   200,
			errorMessage: "",
		},
		{
			name:         "request payload is not a book copy",
			path:         initialBookCopy.ID,
			bookCopy:     "definitely not a book copy, just a plain string",
			statusCode:   400,
			errorMessage: "",
		},
		{
			name: "invalid book copy id path",
			path: util.NewID(),
			bookCopy: &bookcopy.BookCopy{
				Condition: "revised condition",
			},
			statusCode:   500,
			errorMessage: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			jsonReq, _ := json.Marshal(tc.bookCopy)
			url := fmt.Sprintf("/books/" + initialBookCopy.BookID + "/bookcopies/" + tc.path)

			req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonReq))
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

			updatedBookCopy := bookcopy.BookCopy{}
			err = json.Unmarshal(body, &updatedBookCopy)
			if err != nil {
				t.Fatalf("expected a Book Copy struct; got %s", body)
			}

			if updatedBookCopy.Condition != tc.condition {
				t.Errorf("expected condition to be %v; got %v", tc.condition, updatedBookCopy.Condition)
			}
		})
	}

	repository.CleanUp()
}

func TestDeleteBookCopy(t *testing.T) {
	initialBookCopy := createBookCopy()

	tt := []struct {
		name         string
		path         string
		statusCode   int
		errorMessage string
	}{
		{
			name:         "success deleting a valid book copy",
			path:         initialBookCopy.ID,
			statusCode:   200,
			errorMessage: "",
		},
		{
			name:         "invalid book copy id path",
			path:         util.NewID(),
			statusCode:   500,
			errorMessage: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("/books/" + initialBookCopy.BookID + "/bookcopies/" + tc.path)

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
