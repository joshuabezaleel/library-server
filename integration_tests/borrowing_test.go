package integrationtests

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

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

			req := httptest.NewRequest("POST", url, nil)

			ctx := req.Context()
			ctx = context.WithValue(ctx, "username", tc.username)

			req = req.WithContext(ctx)
			rr := httptest.NewRecorder()
			srv.Router.ServeHTTP(rr, req)

			resp := rr.Result()
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				require.Equal(t, tc.statusCode, resp.StatusCode)
				return
			}

			body, err := ioutil.ReadAll(resp.Body)
			require.Nil(t, err)

			newBorrow := borrowing.Borrow{}
			err = json.Unmarshal(body, &newBorrow)
			require.Nil(t, err)

			require.Equal(t, initialBookCopy.ID, newBorrow.BookCopyID)

			require.Equal(t, initialUser.ID, newBorrow.UserID)
		})
	}

	repository.CleanUp()
}

func TestReturn(t *testing.T) {
	initialBookCopy := createBookCopy()
	otherBookCopy := createBookCopy()
	initialUser := createUser()

	// Borrow the initialBookCopy by initialUser.
	url := fmt.Sprintf("/books/" + initialBookCopy.BookID + "/bookcopies/" + initialBookCopy.ID + "/borrow")

	req := httptest.NewRequest("POST", url, nil)

	ctx := req.Context()
	ctx = context.WithValue(ctx, "username", initialUser.Username)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	srv.Router.ServeHTTP(rr, req)

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

			req := httptest.NewRequest("POST", url, nil)

			ctx := req.Context()
			ctx = context.WithValue(ctx, "username", tc.returnerUsername)

			req = req.WithContext(ctx)
			rr := httptest.NewRecorder()
			srv.Router.ServeHTTP(rr, req)

			resp := rr.Result()
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				require.Equal(t, tc.statusCode, resp.StatusCode)
				return
			}

			body, err := ioutil.ReadAll(resp.Body)
			require.Nil(t, err)

			newBorrow := borrowing.Borrow{}
			err = json.Unmarshal(body, &newBorrow)
			require.Nil(t, err)

			require.Equal(t, initialBookCopy.ID, newBorrow.BookCopyID)

			require.Equal(t, initialUser.ID, newBorrow.UserID)
		})
	}

	repository.CleanUp()
}
