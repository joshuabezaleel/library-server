package persistence

import (
	"testing"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
	"github.com/stretchr/testify/require"
)

func TestUserSave(t *testing.T) {
	// Create a new User and save it.
	user := &user.User{
		ID:       util.NewID(),
		Username: "username",
	}
	user1, err := repository.UserRepository.Save(user)

	// Happy path.
	require.Nil(t, err)
	require.Equal(t, user1.ID, user.ID)

	repository.CleanUp()
}

func TestUserGet(t *testing.T) {
	// Create a new User and save it.
	user := &user.User{
		ID: util.NewID(),
	}
	user1, err := repository.UserRepository.Save(user)
	require.Nil(t, err)

	// Get the User.
	user2, err := repository.UserRepository.Get(user.ID)
	require.Nil(t, err)
	require.Equal(t, user2.ID, user1.ID)

	// Get invalid User.
	_, err = repository.UserRepository.Get(util.NewID())
	require.NotNil(t, err)

	repository.CleanUp()
}

func TestUserUpdate(t *testing.T) {
	// Create a new User and save it.
	user := &user.User{
		ID:       util.NewID(),
		Username: "username",
	}
	user1, err := repository.UserRepository.Save(user)
	require.Nil(t, err)

	// Update the User's username.
	user1.Username = "edited username"
	user2, err := repository.UserRepository.Update(user1)
	require.Nil(t, err)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user.Username, "edited username")

	repository.CleanUp()
}

func TestUserDelete(t *testing.T) {
	// Create a new User and save it.
	user := &user.User{
		ID: util.NewID(),
	}
	_, err := repository.UserRepository.Save(user)
	require.Nil(t, err)

	// Delete the User that was just created.
	err = repository.UserRepository.Delete(user.ID)
	require.Nil(t, err)

	// Unable to retrieve the User that was just deleted.
	_, err = repository.UserRepository.Get(user.ID)
	require.NotNil(t, err)

	repository.CleanUp()
}

func TestUserGetIDByUsername(t *testing.T) {
	// Create a new User and save it.
	user := &user.User{
		ID:       util.NewID(),
		Username: "username",
	}
	_, err := repository.UserRepository.Save(user)
	require.Nil(t, err)

	// Get User ID by the username.
	user2ID, err := repository.UserRepository.GetIDByUsername(user.Username)
	require.Nil(t, err)
	require.Equal(t, user2ID, user.ID)

	repository.CleanUp()
}

func TestUserCheckLibrarian(t *testing.T) {
	// Create a new User with the role "librarian" and save it.
	user := &user.User{
		ID:   util.NewID(),
		Role: "librarian",
	}
	user1, err := repository.UserRepository.Save(user)
	require.Nil(t, err)

	// Check the User's role.
	role, err := repository.UserRepository.CheckLibrarian(user1.ID)
	require.Nil(t, err)
	require.Equal(t, role, "librarian")

	repository.CleanUp()
}

func TestUserAddFine(t *testing.T) {
	// Create a new User and save it.
	user := &user.User{
		ID: util.NewID(),
	}
	user1, err := repository.UserRepository.Save(user)
	require.Nil(t, err)

	// Add fine to the user.
	var fine uint32 = 7000
	err = repository.UserRepository.AddFine(user.ID, fine)
	require.Nil(t, err)

	// Get the User and check the fine amount.
	user2, err := repository.UserRepository.Get(user.ID)
	require.Nil(t, err)
	require.Equal(t, user2.ID, user1.ID)
	require.Equal(t, user2.TotalFine, fine)

	repository.CleanUp()
}

func TestUserGetTotalFine(t *testing.T) {

}
