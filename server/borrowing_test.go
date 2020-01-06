package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/borrowing"
)

func TestBorrow(t *testing.T) {
	initialBookCopy := createBookCopy()
	initialUser := createUser()

	tt := []struct {
		name         string
		bookID       string
		bookCopyID   string
		username     string
		statusCode   int
		errorMessage string
	}{
		{
			name:         "success borrowing a book copy",
			bookID:       initialBookCopy.BookID,
			bookCopyID:   initialBookCopy.ID,
			username:     initialUser.Username,
			statusCode:   200,
			errorMessage: "",
		},
		{
			name:         "invalid book",
			bookID:       util.NewID(),
			bookCopyID:   initialBookCopy.ID,
			username:     initialUser.Username,
			statusCode:   500,
			errorMessage: "",
		},
		{
			name:         "invalid bookcopy",
			bookID:       initialBookCopy.BookID,
			bookCopyID:   util.NewID(),
			username:     initialUser.Username,
			statusCode:   500,
			errorMessage: "",
		},
		{
			name:         "invalid user",
			bookID:       initialBookCopy.BookID,
			bookCopyID:   initialBookCopy.ID,
			username:     "another random username",
			statusCode:   500,
			errorMessage: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("/books/" + tc.bookID + "/bookcopies/" + tc.bookCopyID + "/borrow")

			req, err := http.NewRequest("POST", url, nil)
			if err != nil {
				t.Errorf("Cannot perform HTTP request: %v", err)
			}

			ctx := req.Context()
			ctx = context.WithValue(ctx, "username", tc.username)

			req = req.WithContext(ctx)
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

			newBorrow := borrowing.Borrow{}
			err = json.Unmarshal(body, &newBorrow)
			if err != nil {
				t.Fatalf("expected a Borrow struct; got %s", body)
			}

			if newBorrow.BookCopyID != initialBookCopy.ID {
				t.Fatalf("expected book copy id %v; got %v", initialBookCopy.ID, newBorrow.BookCopyID)
			}

			if initialUser.ID != newBorrow.UserID {
				t.Fatalf("expected user id %v; got %v", initialUser.ID, newBorrow.UserID)
			}
		})
	}

	repository.CleanUp()
}

func TestReturn(t *testing.T) {
	initialBookCopy := createBookCopy()
	otherBookCopy := createBookCopy()
	initialUser := createUser()
	// log.Printf("book copy id: %v\n", initialBookCopy.ID)
	// log.Printf("other book copy id: %v\n", otherBookCopy.ID)
	// log.Printf("user id: %v\n", initialUser.ID)

	// Borrow the initialBookCopy by initialUser.
	url := fmt.Sprintf("/books/" + initialBookCopy.BookID + "/bookcopies/" + initialBookCopy.ID + "/borrow")

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		t.Errorf("Cannot perform HTTP request: %v", err)
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx, "username", initialUser.Username)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	srv.Router.ServeHTTP(rr, req)

	// log.Println(rr.Body.String())

	tt := []struct {
		name             string
		bookID           string
		bookCopyID       string
		borrowerUsername string
		returnerUsername string
		statusCode       int
		errorMessage     string
	}{
		{
			name:             "success returning a book copy",
			bookID:           initialBookCopy.BookID,
			bookCopyID:       initialBookCopy.ID,
			borrowerUsername: initialUser.Username,
			returnerUsername: initialUser.Username,
			statusCode:       200,
			errorMessage:     "",
		},
		{
			name:             "returning a not borrowed bookcopy",
			bookID:           initialBookCopy.BookID,
			bookCopyID:       otherBookCopy.ID,
			borrowerUsername: initialUser.Username,
			returnerUsername: initialUser.Username,
			statusCode:       500,
			errorMessage:     "",
		},
		{
			name:             "returning other user's bookcopy",
			bookID:           initialBookCopy.BookID,
			bookCopyID:       initialBookCopy.ID,
			borrowerUsername: initialUser.Username,
			returnerUsername: "otherUsername",
			statusCode:       500,
			errorMessage:     "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("/books/" + tc.bookID + "/bookcopies/" + tc.bookCopyID + "/return")
			// log.Println(url)

			req, err := http.NewRequest("POST", url, nil)
			if err != nil {
				t.Errorf("Cannot perform HTTP request: %v", err)
			}

			ctx := req.Context()
			ctx = context.WithValue(ctx, "username", tc.returnerUsername)

			req = req.WithContext(ctx)
			rr := httptest.NewRecorder()
			srv.Router.ServeHTTP(rr, req)

			// log.Println(rr.Body.String())

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

			newBorrow := borrowing.Borrow{}
			err = json.Unmarshal(body, &newBorrow)
			if err != nil {
				t.Fatalf("expected a Borrow struct; got %s", body)
			}

			if newBorrow.BookCopyID != initialBookCopy.ID {
				t.Fatalf("expected book copy id %v; got %v", initialBookCopy.ID, newBorrow.BookCopyID)
			}

			if initialUser.ID != newBorrow.UserID {
				t.Fatalf("expected user id %v; got %v", initialUser.ID, newBorrow.UserID)
			}
		})
	}

	repository.CleanUp()
}
