package auth

import (
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/stretchr/testify/require"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
)

var userRepository = &user.MockRepository{}
var authRepository = &MockRepository{}

var userService = user.NewUserService(userRepository)
var authService = NewAuthService(authRepository, userService)

func TestGetPassword(t *testing.T) {
	user := &user.User{
		ID:       util.NewID(),
		Username: "username",
		Password: "password",
	}

	authRepository.On("GetPassword", user.Username).Return(user.Password, nil)

	password, err := authService.GetStoredPasswordByUsername(user.Username)

	require.Nil(t, err)
	require.Equal(t, user.Password, password)
}

func TestComparePassword(t *testing.T) {
	incomingPassword := "incomingPassword"
	expectedPassword := hashAndSalt(incomingPassword)

	isSamePassword, err := authService.ComparePassword(incomingPassword, expectedPassword)

	require.Nil(t, err)
	require.True(t, isSamePassword)
}

func hashAndSalt(password string) string {
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	return string(hash)
}
