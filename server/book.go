package server

import (
	"net/http"

	"github.com/joshuabezaleel/library-server/pkg/core/book"

	"github.com/gorilla/mux"
)

type bookHandler struct {
	bookService book.Service
}

func (handler *bookHandler) registerRouter(router *mux.Router) {
	// CRUD endpoints.
	router.HandleFunc("/books", handler.createBook).Methods("POST")
	router.HandleFunc("/books/{bookID}", handler.getBook).Methods("GET")
	router.HandleFunc("/books/{bookID}", handler.updateBook).Methods("PUT")
	router.HandleFunc("/books/{bookID}", handler.deleteBook).Methods("DELETE")

	// Other endpoints.
}

func (handler *bookHandler) createBook(w http.ResponseWriter, r *http.Request) {

}

func (handler *bookHandler) getBook(w http.ResponseWriter, r *http.Request) {

}

func (handler *bookHandler) updateBook(w http.ResponseWriter, r *http.Request) {

}

func (handler *bookHandler) deleteBook(w http.ResponseWriter, r *http.Request) {

}
