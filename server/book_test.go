package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
)

func TestCreateBook(t *testing.T) {
	book := &book.Book{
		ID:    util.NewID(),
		Title: "title",
	}
	jsonReq, _ := json.Marshal(book)

	req, err := http.NewRequest("POST", "localhost:8082/books", bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	srv.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("rec.Code = %d; want = %d", rr.Code, http.StatusCreated)
	}
	// srv.Router.c
	// testingSrv.Router.ServeHTTP(rr, req)
	// t.Log("test")
}
func TestGetBook(t *testing.T) {
}
