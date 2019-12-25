package persistence

import (
	"testing"

	"github.com/stretchr/testify/require"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
)

func TestGetPassword(t *testing.T) {
	// Create a User with a password and save it.
	user := &user.User{
		ID:       util.NewID(),
		Username: "usernamefortesting2",
		Password: "passwordfortesting2",
	}
	user1, err := repository.UserRepository.Save(user)
	require.Nil(t, err)

	// Get the User's password
	password, err := repository.AuthRepository.GetPassword(user.Username)
	require.Nil(t, err)
	require.Equal(t, password, user1.Password)
}
