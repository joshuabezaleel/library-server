package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joshuabezaleel/library-server/pkg/auth"
	"github.com/joshuabezaleel/library-server/pkg/core/book"

	"github.com/gorilla/mux"
)

type bookHandler struct {
	bookService book.Service
	authService auth.Service
}

func (handler *bookHandler) registerRouter(deployment string, router *mux.Router) {
	log.Printf("registerRouter dipanggil kok ... deployment = %s", deployment)
	if deployment == "PRODUCTION" {
		// CRUD endpoints.
		router.HandleFunc("/books", handler.authService.CheckLoggedInMiddleware(handler.authService.CheckLibrarian(handler.CreateBook))).Methods("POST")
		router.HandleFunc("/books/{bookID}", handler.getBook).Methods("GET")
		router.HandleFunc("/books/{bookID}", handler.authService.CheckLoggedInMiddleware(handler.authService.CheckLibrarian(handler.updateBook))).Methods("PUT")
		router.HandleFunc("/books/{bookID}", handler.authService.CheckLoggedInMiddleware(handler.authService.CheckLibrarian(handler.deleteBook))).Methods("DELETE")

		// Other endpoints.
	} else if deployment == "TESTING" {
		// CRUD endpoints.
		// router.HandleFunc("/books", handler.authService.CheckLoggedInMiddleware(handler.authService.CheckLibrarian(handler.CreateBook))).Methods("POST")
		router.HandleFunc("/books", handler.CreateBook).Methods("POST")
		router.HandleFunc("/books/{bookID}", handler.getBook).Methods("GET")
		router.HandleFunc("/books/{bookID}", handler.updateBook).Methods("PUT")
		router.HandleFunc("/books/{bookID}", handler.deleteBook).Methods("DELETE")

		// Other endpoints.
	}

}

// CreateBook handles the creation of a Book.
func (handler *bookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	book := book.Book{}

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	newBook, err := handler.bookService.Create(&book)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, newBook)
}

func (handler *bookHandler) getBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, ok := vars["bookID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	book, err := handler.bookService.Get(bookID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, book)
}

func (handler *bookHandler) updateBook(w http.ResponseWriter, r *http.Request) {
	book := book.Book{}

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	bookID, ok := vars["bookID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}
	book.ID = bookID

	updatedBook, err := handler.bookService.Update(&book)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, updatedBook)
}

func (handler *bookHandler) deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, ok := vars["bookID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	err := handler.bookService.Delete(bookID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, "Book "+bookID+" deleted")
}
