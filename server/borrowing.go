package server

import (
	"net/http"

	"github.com/joshuabezaleel/library-server/pkg/auth"
	"github.com/joshuabezaleel/library-server/pkg/borrowing"

	"github.com/gorilla/mux"
)

type borrowingHandler struct {
	borrowingService borrowing.Service
	authService      auth.Service
}

func (handler *borrowingHandler) registerRouter(deployment string, router *mux.Router) {
	if deployment == "PRODUCTION" {
		router.HandleFunc("/books/{bookID}/bookcopies/{bookCopyID}/borrow", handler.authService.CheckLoggedInMiddleware(handler.borrowBookCopy)).Methods("POST")
		router.HandleFunc("/books/{bookID}/bookcopies/{bookCopyID}/return", handler.authService.CheckLoggedInMiddleware(handler.returnBookCopy)).Methods("POST")
	} else if deployment == "TESTING" {
		router.HandleFunc("/books/{bookID}/bookcopies/{bookCopyID}/borrow", handler.borrowBookCopy).Methods("POST")
		router.HandleFunc("/books/{bookID}/bookcopies/{bookCopyID}/return", handler.returnBookCopy).Methods("POST")
	}

}

func (handler *borrowingHandler) borrowBookCopy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookCopyID, ok := vars["bookCopyID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid book copy ID")
		return
	}

	username := r.Context().Value("username").(string)

	borrow, err := handler.borrowingService.Borrow(username, bookCopyID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, borrow)
}

func (handler *borrowingHandler) returnBookCopy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookCopyID, ok := vars["bookCopyID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid book copy ID")
		return
	}

	username := r.Context().Value("username").(string)

	borrow, err := handler.borrowingService.Return(username, bookCopyID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, borrow)
}
