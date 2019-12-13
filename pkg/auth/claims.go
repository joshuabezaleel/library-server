package auth

import (
	"github.com/dgrijalva/jwt-go"
)

// SecretKey defines secret string for JWT Token.
var SecretKey = []byte("CdfiJ4E73IOGWt8MC")

// Claims defines claims for JWT Token along with User's Username.
type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}
