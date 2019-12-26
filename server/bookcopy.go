package server

import (
	"encoding/json"
	"net/http"

	"github.com/joshuabezaleel/library-server/pkg/auth"
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"

	"github.com/gorilla/mux"
)

type bookCopyHandler struct {
	bookCopyService bookcopy.Service
	authService     auth.Service
}

func (handler *bookCopyHandler) registerRouter(deployment string, router *mux.Router) {
	if deployment == "PRODUCTION" {
		// CRUD endpoints.
		router.HandleFunc("/books/{bookID}/bookcopies", handler.authService.CheckLoggedInMiddleware(handler.authService.CheckLibrarian(handler.createBookCopy))).Methods("POST")
		router.HandleFunc("/books/{bookID}/bookcopies/{bookCopyID}", handler.getBookCopy).Methods("GET")
		router.HandleFunc("/books/{bookID}/bookcopies/{bookCopyID}", handler.authService.CheckLoggedInMiddleware(handler.authService.CheckLibrarian(handler.updateBookCopy))).Methods("PUT")
		router.HandleFunc("/books/{bookID}/bookcopies/{bookCopyID}", handler.authService.CheckLoggedInMiddleware(handler.authService.CheckLibrarian(handler.deleteBookCopy))).Methods("DELETE")

		// Other endpoints.
	} else if deployment == "TESTING" {
		// CRUD endpoints.
		router.HandleFunc("/books/{bookID}/bookcopies", handler.createBookCopy).Methods("POST")
		router.HandleFunc("/books/{bookID}/bookcopies/{bookCopyID}", handler.getBookCopy).Methods("GET")
		router.HandleFunc("/books/{bookID}/bookcopies/{bookCopyID}", handler.updateBookCopy).Methods("PUT")
		router.HandleFunc("/books/{bookID}/bookcopies/{bookCopyID}", handler.deleteBookCopy).Methods("DELETE")

		// Other endpoints.
	}

}

func (handler *bookCopyHandler) createBookCopy(w http.ResponseWriter, r *http.Request) {
	bookCopy := bookcopy.BookCopy{}

	err := json.NewDecoder(r.Body).Decode(&bookCopy)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	bookID, ok := vars["bookID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}
	bookCopy.BookID = bookID

	newBookCopy, err := handler.bookCopyService.Create(&bookCopy)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, newBookCopy)
}

func (handler *bookCopyHandler) getBookCopy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookCopyID, ok := vars["bookCopyID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid book copy ID")
		return
	}

	bookCopy, err := handler.bookCopyService.Get(bookCopyID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, bookCopy)
}

func (handler *bookCopyHandler) updateBookCopy(w http.ResponseWriter, r *http.Request) {
	bookCopy := bookcopy.BookCopy{}

	err := json.NewDecoder(r.Body).Decode(&bookCopy)
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
	bookCopy.BookID = bookID

	bookCopyID, ok := vars["bookCopyID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid book copy ID")
		return
	}
	bookCopy.ID = bookCopyID

	updatedBookCopy, err := handler.bookCopyService.Update(&bookCopy)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, updatedBookCopy)
}

func (handler *bookCopyHandler) deleteBookCopy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookCopyID, ok := vars["bookCopyID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid book copy ID")
		return
	}

	err := handler.bookCopyService.Delete(bookCopyID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, "Book copy "+bookCopyID+" deleted")
}
