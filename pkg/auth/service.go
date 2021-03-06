package auth

import (
	"context"
	"encoding/json"

	"github.com/joshuabezaleel/library-server/pkg/core/user"

	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"github.com/gorilla/mux"

	"golang.org/x/crypto/bcrypt"
)

type contextKeyType string

// Service provides basic operations for Auth service.
type Service interface {
	GetStoredPasswordByUsername(username string) (string, error)
	ComparePassword(incomingPassword, storedPassword string) (bool, error)
	CheckLoggedInMiddleware(next http.HandlerFunc) http.HandlerFunc
	CheckLibrarian(next http.HandlerFunc) http.HandlerFunc
	CheckSameUser(next http.HandlerFunc) http.HandlerFunc
}

type service struct {
	authRepository Repository
	userService    user.Service
}

// NewAuthService creates an instance of the service for user domain model with necessary dependencies.
func NewAuthService(authRepository Repository, userService user.Service) Service {
	return &service{
		authRepository: authRepository,
		userService:    userService,
	}
}

func (s *service) GetStoredPasswordByUsername(username string) (string, error) {
	password, err := s.authRepository.GetPassword(username)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (s *service) ComparePassword(incomingPassword, storedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(incomingPassword))
	if err != nil {
		return false, err
	}

	return true, nil
}

// nolint
// NOLINT
func (s *service) CheckLoggedInMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := jwt.MapClaims{}
		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")

			if len(bearerToken) == 2 {
				token, err := jwt.ParseWithClaims(bearerToken[1], claims, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return SecretKey, nil
				})
				if err != nil {
					w.Write([]byte(err.Error()))
					return
				}

				if token.Valid {
					username := claims["username"].(string)

					// var contextKey contextKeyType
					// contextKey = "username"
					// nolint
					ctx := context.WithValue(r.Context(), "username", username) // NOLINT
					next(w, r.WithContext(ctx))
				} else {
					respondWithError(w, http.StatusUnauthorized, "Invalid authorization token")
				}
			}
		} else {
			respondWithError(w, http.StatusNotAcceptable, "An authorization header is required")
		}
	})
}

func (s *service) CheckLibrarian(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.Context().Value("username").(string)

		role, err := s.userService.GetRole(username)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if role != "librarian" {
			respondWithError(w, http.StatusUnauthorized, "You are not authorized as a librarian to perform this action")
		} else {
			next(w, r)
		}
	})
}

func (s *service) CheckSameUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usernameLoggedIn := r.Context().Value("username").(string)

		IDLoggedIn, err := s.userService.GetUserIDByUsername(usernameLoggedIn)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		vars := mux.Vars(r)
		userID, ok := vars["userID"]
		if !ok {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		if IDLoggedIn != userID {
			respondWithError(w, http.StatusUnauthorized, "You are not authorized to perform this action")
		} else {
			next(w, r)
		}
	})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"Error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
