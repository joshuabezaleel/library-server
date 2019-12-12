package server

import (
	"encoding/json"
	"net/http"

	"github.com/joshuabezaleel/library-server/pkg/core/user"

	"github.com/gorilla/mux"
)

type userHandler struct {
	userService user.Service
}

func (handler *userHandler) registerRouter(router *mux.Router) {
	// CRUD endpoints.
	router.HandleFunc("/users", handler.createUser).Methods("POST")
	router.HandleFunc("/users/{userID}", handler.getUser).Methods("GET")
	router.HandleFunc("/users/{userID}", handler.updateUser).Methods("PUT")
	router.HandleFunc("/users/{userID}", handler.deleteUser).Methods("DELETE")

	// Other endpoints.
}

func (handler *userHandler) createUser(w http.ResponseWriter, r *http.Request) {
	user := user.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
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
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	newUser, err := handler.userService.Get(userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, newUser)
}

func (handler *userHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	user := user.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	userID, ok := vars["userID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	user.ID = userID

	newUser, err := handler.userService.Update(&user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, newUser)
}

func (handler *userHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, ok := vars["userID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	err := handler.userService.Delete(userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, "User "+userID+" deleted")
}
