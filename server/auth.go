package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/joshuabezaleel/library-server/pkg/auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type authHandler struct {
	authService auth.Service
}

func (handler *authHandler) registerRouter(router *mux.Router) {
	router.HandleFunc("/login", handler.login).Methods("POST")
}

func (handler *authHandler) login(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	storedPassword, err := handler.authService.GetStoredPasswordByUsername(request.Username)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = handler.authService.ComparePassword(request.Password, storedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Password is wrong")
		return
	}
	expirationTime := time.Now().Add(5 * time.Minute).Unix()
	claims := &auth.Claims{
		Username: request.Username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Library Server",
			ExpiresAt: expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(auth.SecretKey)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error issuing token")
		return
	}

	respondWithJSON(w, http.StatusOK, tokenString)
}
