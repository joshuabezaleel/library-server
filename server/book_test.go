package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateBook(t *testing.T) {
	req, err := http.NewRequest("POST", "/books", nil)
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
