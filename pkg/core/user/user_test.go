package user

import (
	"errors"
	"testing"
	"time"

	"github.com/bouk/monkey"
	"github.com/stretchr/testify/require"

	util "github.com/joshuabezaleel/library-server/pkg"
)

var userRepository = &MockRepository{}
var userService = service{userRepository: userRepository}

func TestCreate(t *testing.T) {
	createdTime := time.Now()
	timePatch := monkey.Patch(time.Now, func() time.Time {
		return createdTime
	})
	defer timePatch.Unpatch()

	expectedPwd := "password"
	passwordPatch := monkey.Patch(hashAndSalt, func(password string) string {
		return expectedPwd
	})
	defer passwordPatch.Unpatch()

	user := &User{
		ID:           util.NewID(),
		RegisteredAt: createdTime,
		Password:     expectedPwd,
	}
	userRepository.On("Save", user).Return(user, nil)

	newUser, err := userService.Create(user)

	require.Nil(t, err)
	require.Equal(t, user.ID, newUser.ID)
}

func TestGet(t *testing.T) {
	user := &User{
		ID: util.NewID(),
	}
	userRepository.On("Get", user.ID).Return(user, nil)

	newUser, err := userService.Get(user.ID)

	require.Nil(t, err)
	require.Equal(t, user.ID, newUser.ID)
}
func TestUpdate(t *testing.T) {
	user := &User{
		ID:       util.NewID(),
		Username: "username",
	}

	expectedUser := &User{
		ID:       user.ID,
		Username: "editedusername",
	}

	userRepository.On("Update", user).Return(expectedUser, nil)

	updatedUser, err := userService.Update(user)

	require.Nil(t, err)
	require.Equal(t, user.ID, updatedUser.ID)
	require.Equal(t, expectedUser.Username, updatedUser.Username)
}

func TestDelete(t *testing.T) {
	user := &User{
		ID: util.NewID(),
	}

	userRepository.On("Delete", user.ID).Return(nil)

	err := userService.Delete(user.ID)

	require.Nil(t, err)
}

func TestGetUserIDByUsername(t *testing.T) {
	user := &User{
		ID:       util.NewID(),
		Username: "username",
	}

	userRepository.On("GetIDByUsername", user.Username).Return(user.ID, nil)

	userID, err := userService.GetUserIDByUsername(user.Username)

	require.Nil(t, err)
	require.Equal(t, user.ID, userID)
}

func TestCheckLibrarian(t *testing.T) {
	tt := []struct {
		name        string
		user        *User
		isLibrarian bool
		err         error
	}{
		{
			name: "user with role librarian",
			user: &User{
				ID:       util.NewID(),
				Username: "librarian1",
				Role:     "librarian",
			},
			isLibrarian: true,
			err:         nil,
		},
		{
			name: "user who is not a librarian",
			user: &User{
				ID:       util.NewID(),
				Username: "student1",
				Role:     "student",
			},
			isLibrarian: false,
			err:         errors.New("You are not authorized as a librarian to perform this action"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userRepository.On("GetIDByUsername", tc.user.Username).Return(tc.user.ID, nil)

			userRepository.On("GetRole", tc.user.ID).Return(tc.user.Role, nil)

			isLibrarian, err := userService.CheckLibrarian(tc.user.Username)

			require.Equal(t, tc.isLibrarian, isLibrarian)
			require.Equal(t, tc.err, err)
		})
	}

	// Testing error on CheckLibrarian
	userRepository.On("GetIDByUsername", "random username").Return("random user ID", nil)
	userRepository.On("GetRole", "random user ID").Return("", errors.New("another error"))

	_, err := userService.CheckLibrarian("random username")
	require.NotNil(t, err)

	// Testing error on GetIDByUsername
	userRepository.On("GetIDByUsername", "random username 2").Return("", errors.New("another error"))
	// userRepository.On("CheckLibrarian", "")

	_, err = userService.CheckLibrarian("random username 2")
	require.NotNil(t, err)
}

func TestGetTotalFine(t *testing.T) {
	user := &User{
		ID:        util.NewID(),
		TotalFine: 7000,
	}

	userRepository.On("GetTotalFine", user.ID).Return(user.TotalFine, nil)

	totalFine, err := userService.GetTotalFine(user.ID)

	require.Nil(t, err)
	require.Equal(t, user.TotalFine, totalFine)
}

func TestAddFine(t *testing.T) {
	// Happy path.
	user := &User{
		ID:        util.NewID(),
		TotalFine: 2000,
	}
	var fine uint32 = 7000

	userRepository.On("GetTotalFine", user.ID).Return(user.TotalFine, nil)

	userRepository.On("AddFine", user.ID, user.TotalFine+fine).Return(nil)

	addedFine, err := userService.AddFine(user.ID, fine)
	require.Nil(t, err)
	require.Equal(t, user.TotalFine+fine, addedFine)

	// Error on GetTotalFine
	userRepository.On("GetTotalFine", "another username").Return(uint32(0), errors.New("another error"))

	addedFine, err = userService.AddFine("another username", fine)
	require.NotNil(t, err)
	require.Equal(t, uint32(0), addedFine)
}
