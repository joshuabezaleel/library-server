package auth

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

// SecretKey defines secret string for JWT Token.
var SecretKey = []byte(os.Getenv("SECRET_KEY"))

// Claims defines claims for JWT Token along with User's Username.
type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}
