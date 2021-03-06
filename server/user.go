package server

import (
	"encoding/json"
	"net/http"

	"github.com/joshuabezaleel/library-server/pkg/auth"
	"github.com/joshuabezaleel/library-server/pkg/core/user"

	"github.com/gorilla/mux"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func (handler *userHandler) registerRouter(router *mux.Router) {
	// CRUD endpoints.
	router.HandleFunc("/users", handler.createUser).Methods("POST")
	router.HandleFunc("/users/{userID}", handler.getUser).Methods("GET")
	router.HandleFunc("/users/{userID}", handler.authService.CheckLoggedInMiddleware(handler.authService.CheckSameUser(handler.updateUser))).Methods("PUT")
	router.HandleFunc("/users/{userID}", handler.authService.CheckLoggedInMiddleware(handler.authService.CheckSameUser(handler.deleteUser))).Methods("DELETE")

	// Other endpoints.
}

func (handler *userHandler) createUser(w http.ResponseWriter, r *http.Request) {
	user := user.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, errInvalidRequestPayload.Error())
		return
	}
	defer r.Body.Close()

	newUser, err := handler.userService.Create(&user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, newUser)
}

func (handler *userHandler) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, ok := vars["userID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, errInvalidURLPath.Error())
		return
	}

	user, err := handler.userService.Get(userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (handler *userHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	user := user.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, errInvalidRequestPayload.Error())
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	userID, ok := vars["userID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, errInvalidURLPath.Error())
		return
	}
	user.ID = userID

	updatedUser, err := handler.userService.Update(&user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, updatedUser)
}

func (handler *userHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, ok := vars["userID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, errInvalidURLPath.Error())
		return
	}

	err := handler.userService.Delete(userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, "User "+userID+" deleted")
}
