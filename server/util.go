package server

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	errInvalidRequestPayload = errors.New("Invalid request payload")
	errInvalidURLPath        = errors.New("Invalid URL path")
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"Error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
