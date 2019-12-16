package server

import (
	"net/http"

	"github.com/joshuabezaleel/library-server/pkg/auth"
	"github.com/joshuabezaleel/library-server/pkg/borrowing"
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
	"github.com/joshuabezaleel/library-server/pkg/core/user"

	"github.com/gorilla/mux"
)

type borrowingHandler struct {
	borrowingService borrowing.Service
	bookCopyService  bookcopy.Service
	userService      user.Service
	authService      auth.Service
}

func (handler *borrowingHandler) registerRouter(router *mux.Router) {
	router.HandleFunc("/books/{bookID}/bookcopies/{bookCopyID}/borrow", handler.authService.CheckLoggedInMiddleware(handler.borrowBookCopy)).Methods("POST")
	router.HandleFunc("/books/{bookID}/bookcopies/{bookCopyID}/return", handler.authService.CheckLoggedInMiddleware(handler.returnBookCopy)).Methods("POST")
}

func (handler *borrowingHandler) borrowBookCopy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookCopyID, ok := vars["bookCopyID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid book copy ID")
		return
	}

	username := r.Context().Value("username").(string)
	userID, err := handler.userService.GetUserIDByUsername(username)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	borrow, err := handler.borrowingService.Borrow(userID, bookCopyID)
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
	userID, err := handler.userService.GetUserIDByUsername(username)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	borrow, err := handler.borrowingService.Return(userID, bookCopyID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, borrow)
}
