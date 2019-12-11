package server

import (
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

}

func (handler *userHandler) getUser(w http.ResponseWriter, r *http.Request) {

}

func (handler *userHandler) updateUser(w http.ResponseWriter, r *http.Request) {

}

func (handler *userHandler) deleteUser(w http.ResponseWriter, r *http.Request) {

}
