package user

import (
	"testing"

	"github.com/bouk/monkey"
	"github.com/stretchr/testify/require"

	util "github.com/joshuabezaleel/library-server/pkg"
)

var userRepository = &MockRepository{}
var userService = service{userRepository: userRepository}

func TestCreate(t *testing.T) {
	createdTime, createdTimePatch := util.CreatedTimePatch()
	defer createdTimePatch.Unpatch()

	expectedPwd := "password"
	passwordPatch := monkey.Patch(hashAndSalt, func(password string) string {
		return expectedPwd
	})
	defer passwordPatch.Unpatch()

	ID, IDPatch := util.NewIDPatch()
	defer IDPatch.Unpatch()

	user := &User{
		ID:           ID,
		Username:     "username",
		Password:     expectedPwd,
		RegisteredAt: createdTime,
	}

	errorUser := &User{
		ID:           ID,
		Username:     "error username",
		Password:     expectedPwd,
		RegisteredAt: createdTime,
	}

	tt := []struct {
		name         string
		user         *User
		returnedUser *User
		err          error
	}{
		{
			name:         "success creating a User",
			user:         user,
			returnedUser: user,
			err:          nil,
		},
		{
			name:         "failed creating a User",
			user:         errorUser,
			returnedUser: nil,
			err:          ErrCreateUser,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userRepository.On("Save", tc.user).Return(tc.returnedUser, tc.err)

			newUser, err := userService.Create(tc.user)

			require.Equal(t, tc.err, err)

			if tc.err == nil {
				require.Equal(t, user.ID, newUser.ID)
				require.Equal(t, user.Username, newUser.Username)
			}
		})
	}
}

func TestGet(t *testing.T) {
	user := &User{
		ID: util.NewID(),
	}

	tt := []struct {
		name         string
		ID           string
		returnedUser *User
		err          error
	}{
		{
			name:         "success retrieving a User",
			ID:           user.ID,
			returnedUser: user,
			err:          nil,
		},
		{
			name:         "failed retrieving a User",
			ID:           util.NewID(),
			returnedUser: nil,
			err:          ErrGetUser,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userRepository.On("Get", tc.ID).Return(tc.returnedUser, tc.err)

			returnedUser, err := userService.Get(tc.ID)

			require.Equal(t, tc.err, err)

			if tc.err == nil {
				require.Equal(t, user.ID, returnedUser.ID)
			}
		})
	}
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

	errorUser := &User{
		ID: util.NewID(),
	}

	tt := []struct {
		name         string
		user         *User
		returnedUser *User
		err          error
	}{
		{
			name:         "success updating a User",
			user:         user,
			returnedUser: expectedUser,
			err:          nil,
		},
		{
			name:         "failed updating a User",
			user:         errorUser,
			returnedUser: nil,
			err:          ErrUpdateUser,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userRepository.On("Update", tc.user).Return(tc.returnedUser, tc.err)

			updatedUser, err := userService.Update(tc.user)

			require.Equal(t, tc.err, err)

			if tc.err == nil {
				require.Equal(t, expectedUser.ID, updatedUser.ID)
				require.Equal(t, expectedUser.Username, updatedUser.Username)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	user := &User{
		ID: util.NewID(),
	}

	tt := []struct {
		name string
		ID   string
		err  error
	}{
		{
			name: "success deleting a User",
			ID:   user.ID,
			err:  nil,
		},
		{
			name: "failed deleting a User",
			ID:   util.NewID(),
			err:  ErrDeleteUser,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userRepository.On("Delete", tc.ID).Return(tc.err)

			err := userService.Delete(tc.ID)

			require.Equal(t, tc.err, err)
		})
	}
}

func TestGetUserIDByUsername(t *testing.T) {
	user := &User{
		ID:       util.NewID(),
		Username: "username",
	}

	tt := []struct {
		name           string
		username       string
		returnedUserID string
		err            error
	}{
		{
			name:           "success retrieving user's ID",
			username:       user.Username,
			returnedUserID: user.ID,
			err:            nil,
		},
		{
			name:           "failed retrieving user's ID",
			username:       "random username",
			returnedUserID: "",
			err:            ErrGetUserIDByUsername,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userRepository.On("GetIDByUsername", tc.username).Return(tc.returnedUserID, tc.err)

			userID, err := userService.GetUserIDByUsername(tc.username)

			require.Equal(t, tc.err, err)

			if tc.err == nil {
				require.Equal(t, user.ID, userID)
			}
		})
	}
}

func TestGetRole(t *testing.T) {
	user := &User{
		ID:       util.NewID(),
		Username: "username",
		Role:     "librarian",
	}

	errorUser := &User{
		ID:       util.NewID(),
		Username: "error username",
		Role:     "",
	}

	tt := []struct {
		name string
		user *User
		err  error
	}{
		{
			name: "success retrieving user's role",
			user: user,
			err:  nil,
		},
		{
			name: "failed retrieving user's role",
			user: errorUser,
			err:  ErrGetRole,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			userRepository.On("GetIDByUsername", tc.user.Username).Return(tc.user.ID, nil)

			userRepository.On("GetRole", tc.user.ID).Return(tc.user.Role, tc.err)

			role, err := userService.GetRole(tc.user.Username)

			require.Equal(t, tc.err, err)

			if tc.err == nil {
				require.Equal(t, tc.user.Role, role)
			}
		})
	}
}

func TestGetTotalFine(t *testing.T) {
	user := &User{
		ID:        util.NewID(),
		TotalFine: 7000,
	}

	errorUser := &User{
		ID:        util.NewID(),
		TotalFine: 0,
	}

	tt := []struct {
		name string
		user *User
		err  error
	}{
		{
			name: "success retrieving user's total fine",
			user: user,
			err:  nil,
		},
		{
			name: "failed retrieving user's total fine",
			user: errorUser,
			err:  ErrGetTotalFine,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userRepository.On("GetTotalFine", tc.user.ID).Return(tc.user.TotalFine, tc.err)

			totalFine, err := userService.GetTotalFine(tc.user.ID)

			require.Equal(t, tc.err, err)

			if tc.err == nil {
				require.Equal(t, user.TotalFine, totalFine)
			}
		})
	}
}

func TestAddFine(t *testing.T) {
	user := &User{
		ID:        util.NewID(),
		TotalFine: 2000,
	}

	errorUser := &User{
		ID:        util.NewID(),
		TotalFine: 0,
	}

	var fine uint32 = 7000

	tt := []struct {
		name string
		user *User
		fine uint32
		err  error
	}{
		{
			name: "success adding fine to User",
			user: user,
			fine: fine,
			err:  nil,
		},
		{
			name: "failed adding fine to User",
			user: errorUser,
			fine: fine,
			err:  ErrAddFine,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userRepository.On("GetTotalFine", tc.user.ID).Return(tc.user.TotalFine, nil)

			userRepository.On("AddFine", tc.user.ID, tc.user.TotalFine+tc.fine).Return(tc.err)

			totalFine, err := userService.AddFine(tc.user.ID, tc.fine)

			require.Equal(t, tc.err, err)

			if tc.err == nil {
				require.Equal(t, tc.user.TotalFine+fine, totalFine)
			}
		})
	}
	// userRepository.On("GetTotalFine", user.ID).Return(user.TotalFine, nil)

	// userRepository.On("AddFine", user.ID, user.TotalFine+fine).Return(nil)

	// addedFine, err := userService.AddFine(user.ID, fine)
	// require.Nil(t, err)
	// require.Equal(t, user.TotalFine+fine, addedFine)

	// // Error on GetTotalFine
	// userRepository.On("GetTotalFine", "another username").Return(uint32(0), errors.New("another error"))

	// addedFine, err = userService.AddFine("another username", fine)
	// require.NotNil(t, err)
	// require.Equal(t, uint32(0), addedFine)
}
