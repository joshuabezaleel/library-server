package integrationtestsproduction

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
)

func TestCreateBook(t *testing.T) {
	initialBook := &book.Book{
		ID: util.NewID(),
	}

	tt := []struct {
		name               string
		requestPayload     interface{}
		ID                 string
		authorizationToken string
		statusCode         int
		errorMessage       string
	}{
		{
			name:               "success creating a Book by Librarian",
			requestPayload:     initialBook,
			ID:                 initialBook.ID,
			authorizationToken: userLibrarianToken,
			statusCode:         http.StatusCreated,
			errorMessage:       "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			jsonRequest, err := json.Marshal(tc.requestPayload)
			require.Nil(t, err)

			req := httptest.NewRequest("POST", "/books", bytes.NewBuffer(jsonRequest))
			req.Header.Add("Authorization", tc.authorizationToken)

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

			createdBook := &book.Book{}
			err = json.Unmarshal(body, createdBook)
			require.Nil(t, err)

			require.Equal(t, tc, createdBook.ID)
		})
	}

	repository.CleanUp()
}
