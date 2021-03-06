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

func (handler *bookCopyHandler) registerRouter(router *mux.Router) {
	// CRUD endpoints.
	router.HandleFunc("/books/{bookID}/bookcopies", handler.authService.CheckLoggedInMiddleware(handler.authService.CheckLibrarian(handler.createBookCopy))).Methods("POST")
	router.HandleFunc("/books/{bookID}/bookcopies/{bookCopyID}", handler.getBookCopy).Methods("GET")
	router.HandleFunc("/books/{bookID}/bookcopies/{bookCopyID}", handler.authService.CheckLoggedInMiddleware(handler.authService.CheckLibrarian(handler.updateBookCopy))).Methods("PUT")
	router.HandleFunc("/books/{bookID}/bookcopies/{bookCopyID}", handler.authService.CheckLoggedInMiddleware(handler.authService.CheckLibrarian(handler.deleteBookCopy))).Methods("DELETE")

	// Other endpoints.
}

func (handler *bookCopyHandler) createBookCopy(w http.ResponseWriter, r *http.Request) {
	bookCopy := bookcopy.BookCopy{}

	err := json.NewDecoder(r.Body).Decode(&bookCopy)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, errInvalidRequestPayload.Error())
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	bookID, ok := vars["bookID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, errInvalidURLPath.Error())
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
		respondWithError(w, http.StatusBadRequest, errInvalidURLPath.Error())
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
		respondWithError(w, http.StatusBadRequest, errInvalidRequestPayload.Error())
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	bookCopyID, ok := vars["bookCopyID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, errInvalidURLPath.Error())
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
		respondWithError(w, http.StatusBadRequest, errInvalidURLPath.Error())
		return
	}

	err := handler.bookCopyService.Delete(bookCopyID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, "Book copy "+bookCopyID+" deleted")
}
