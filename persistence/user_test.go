package persistence

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
)

func TestUserSave(t *testing.T) {
	tt := []struct {
		name string
		user *user.User
		err  bool
	}{
		{
			name: "save a valid user",
			user: &user.User{
				ID:       util.NewID(),
				Username: "testUsername",
			},
			err: false,
		},
		{
			name: "save an invalid user",
			user: &user.User{
				ID:       util.NewID(),
				Username: "anotherTestUsername",
			},
			err: true,
		},
	}

	// Asssert a save for a valid User.
	validUser := tt[0].user

	result := sqlmock.NewResult(1, 1)

	Mock.ExpectExec("INSERT INTO users").
		WithArgs(validUser.ID, validUser.StudentID, validUser.Role, validUser.Username, validUser.Email, validUser.Password, validUser.TotalFine, validUser.RegisteredAt).
		WillReturnResult(result)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			newUser, err := UserTestingRepository.Save(tc.user)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
			require.Equal(t, tc.user.ID, newUser.ID)
			require.Equal(t, tc.user.Username, newUser.Username)
		})
	}
}

func TestUserGet(t *testing.T) {
	tt := []struct {
		name string
		user *user.User
		err  bool
	}{
		{
			name: "get a valid user",
			user: &user.User{
				ID:       util.NewID(),
				Username: "testUsername",
			},
			err: false,
		},
		{
			name: "get an invalid user",
			user: &user.User{
				ID:       util.NewID(),
				Username: "anotherTestUsername",
			},
			err: true,
		},
	}

	// Assert a get for a valid User.
	validUser := tt[0].user

	rows := sqlmock.NewRows([]string{"id", "username"}).
		AddRow(validUser.ID, validUser.Username)

	Mock.ExpectQuery("SELECT (.+) FROM users WHERE id=?").
		WithArgs(validUser.ID).
		WillReturnRows(rows)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			newUser, err := UserTestingRepository.Get(tc.user.ID)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
			require.Equal(t, tc.user.ID, newUser.ID)
		})
	}
}

func TestUserUpdate(t *testing.T) {
	tt := []struct {
		name string
		user *user.User
		err  bool
	}{
		{
			name: "update a valid user",
			user: &user.User{
				ID:       util.NewID(),
				Username: "testUsername",
			},
			err: false,
		},
		{
			name: "update an invalid user",
			user: &user.User{
				ID:       util.NewID(),
				Username: "anotherTestUsername",
			},
			err: true,
		},
	}

	// Assert an update for a valid User.
	validUser := tt[0].user

	result := sqlmock.NewResult(1, 1)

	Mock.ExpectExec("UPDATE users SET").
		WithArgs(validUser.StudentID, validUser.Role, validUser.Username, validUser.Email, validUser.Password, validUser.TotalFine, validUser.ID).
		WillReturnResult(result)

	rows := sqlmock.NewRows([]string{"id", "username"}).
		AddRow(validUser.ID, validUser.Username)

	Mock.ExpectQuery("SELECT (.+) FROM users WHERE id=?").
		WithArgs(validUser.ID).
		WillReturnRows(rows)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			updatedUser, err := UserTestingRepository.Update(tc.user)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
			require.Equal(t, tc.user.ID, updatedUser.ID)
			require.Equal(t, tc.user.Username, updatedUser.Username)
		})
	}
}

func TestUserDelete(t *testing.T) {
	tt := []struct {
		name string
		user *user.User
		err  bool
	}{
		{
			name: "delete a valid user",
			user: &user.User{
				ID:       util.NewID(),
				Username: "testUsername",
			},
			err: false,
		},
		{
			name: "delete an invalid user",
			user: &user.User{
				ID:       util.NewID(),
				Username: "anotherTestUsername",
			},
			err: true,
		},
	}

	// Assert a delete for a valid User.
	validUser := tt[0].user

	result := sqlmock.NewResult(1, 1)

	Mock.ExpectExec("DELETE FROM users").
		WithArgs(validUser.ID).
		WillReturnResult(result)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := UserTestingRepository.Delete(tc.user.ID)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
		})
	}
}

func TestUserGetIDByUsername(t *testing.T) {
	tt := []struct {
		name string
		user *user.User
		err  bool
	}{
		{
			name: "get a valid user",
			user: &user.User{
				ID:       util.NewID(),
				Username: "testUsername",
			},
			err: false,
		},
		{
			name: "get an invalid user",
			user: &user.User{
				ID:       util.NewID(),
				Username: "anotherTestUsername",
			},
			err: true,
		},
	}

	// Assert a get ID for a valid User's username.
	validUser := tt[0].user

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(validUser.ID)

	Mock.ExpectQuery("SELECT id FROM users WHERE username=?").
		WithArgs(validUser.Username).
		WillReturnRows(rows)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			newUserID, err := UserTestingRepository.GetIDByUsername(tc.user.Username)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
			require.Equal(t, tc.user.ID, newUserID)
		})
	}
}

func TestUserGetRole(t *testing.T) {
	tt := []struct {
		name string
		user *user.User
		err  bool
	}{
		{
			name: "get a valid user's role",
			user: &user.User{
				ID:   util.NewID(),
				Role: "librarian",
			},
			err: false,
		},
		{
			name: "get an invalid user's role",
			user: &user.User{
				ID:   util.NewID(),
				Role: "anotherTestingRole",
			},
			err: true,
		},
	}

	// Assert a get for a valid User's role.
	validUser := tt[0].user

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(validUser.Role)

	Mock.ExpectQuery("SELECT role FROM users WHERE id=?").
		WithArgs(validUser.ID).
		WillReturnRows(rows)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			newUserRole, err := UserTestingRepository.GetRole(tc.user.ID)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
			require.Equal(t, tc.user.Role, newUserRole)
		})
	}
}

func TestUserAddFine(t *testing.T) {
	tt := []struct {
		name   string
		userID string
		fine   uint32
		err    bool
	}{
		{
			name:   "update a valid user's fine",
			userID: util.NewID(),
			fine:   uint32(7000),
			err:    false,
		},
		{
			name:   "update an invalid user's fine",
			userID: util.NewID(),
			fine:   uint32(7000),
			err:    true,
		},
	}

	// Assert an update for a valid User's fine.
	validUserID := tt[0].userID
	validFine := tt[0].fine

	result := sqlmock.NewResult(1, 1)

	Mock.ExpectExec("UPDATE users (.+)").
		WithArgs(validFine, validUserID).
		WillReturnResult(result)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := UserTestingRepository.AddFine(tc.userID, tc.fine)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
		})
	}
}

func GetTotalFine(t *testing.T) {
	tt := []struct {
		name      string
		userID    string
		totalFine uint32
		err       bool
	}{
		{
			name:      "get a valid user's fine",
			userID:    util.NewID(),
			totalFine: uint32(15000),
			err:       false,
		},
		{
			name:      "get an invalid user's fine",
			userID:    util.NewID(),
			totalFine: uint32(15000),
			err:       true,
		},
	}

	// Assert a get for a valid User's total fine.
	validUserID := tt[0].userID
	validFine := tt[0].totalFine

	rows := sqlmock.NewRows([]string{"total_fine"}).
		AddRow(validFine)

	Mock.ExpectQuery("SELECT total_fine FROM users WHERE id=?").
		WithArgs(validUserID).
		WillReturnRows(rows)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			newUserTotalFine, err := UserTestingRepository.GetTotalFine(tc.userID)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
			require.Equal(t, tc.totalFine, newUserTotalFine)
		})
	}
}

// func TestUserSave(t *testing.T) {
// 	// Create a new User and save it.
// 	user := &user.User{
// 		ID:       util.NewID(),
// 		Username: "username",
// 	}
// 	user1, err := repository.UserRepository.Save(user)

// 	// Happy path.
// 	require.Nil(t, err)
// 	require.Equal(t, user1.ID, user.ID)

// 	repository.CleanUp()
// }

// func TestUserGet(t *testing.T) {
// 	// Create a new User and save it.
// 	user := &user.User{
// 		ID: util.NewID(),
// 	}
// 	user1, err := repository.UserRepository.Save(user)
// 	require.Nil(t, err)

// 	// Get the User.
// 	user2, err := repository.UserRepository.Get(user.ID)
// 	require.Nil(t, err)
// 	require.Equal(t, user2.ID, user1.ID)

// 	// Get invalid User.
// 	_, err = repository.UserRepository.Get(util.NewID())
// 	require.NotNil(t, err)

// 	repository.CleanUp()
// }

// func TestUserUpdate(t *testing.T) {
// 	// Create a new User and save it.
// 	user := &user.User{
// 		ID:       util.NewID(),
// 		Username: "username",
// 	}
// 	user1, err := repository.UserRepository.Save(user)
// 	require.Nil(t, err)

// 	// Update the User's username.
// 	user1.Username = "edited username"
// 	user2, err := repository.UserRepository.Update(user1)
// 	require.Nil(t, err)
// 	require.Equal(t, user1.ID, user2.ID)
// 	require.Equal(t, user.Username, "edited username")

// 	repository.CleanUp()
// }

// func TestUserDelete(t *testing.T) {
// 	// Create a new User and save it.
// 	user := &user.User{
// 		ID: util.NewID(),
// 	}
// 	_, err := repository.UserRepository.Save(user)
// 	require.Nil(t, err)

// 	// Delete the User that was just created.
// 	err = repository.UserRepository.Delete(user.ID)
// 	require.Nil(t, err)

// 	// Unable to retrieve the User that was just deleted.
// 	_, err = repository.UserRepository.Get(user.ID)
// 	require.NotNil(t, err)

// 	repository.CleanUp()
// }

// func TestUserGetIDByUsername(t *testing.T) {
// 	// Create a new User and save it.
// 	user := &user.User{
// 		ID:       util.NewID(),
// 		Username: "username",
// 	}
// 	_, err := repository.UserRepository.Save(user)
// 	require.Nil(t, err)

// 	// Get User ID by the username.
// 	user2ID, err := repository.UserRepository.GetIDByUsername(user.Username)
// 	require.Nil(t, err)
// 	require.Equal(t, user2ID, user.ID)

// 	repository.CleanUp()
// }

// func TestUserCheckLibrarian(t *testing.T) {
// 	// Create a new User with the role "librarian" and save it.
// 	user := &user.User{
// 		ID:   util.NewID(),
// 		Role: "librarian",
// 	}
// 	user1, err := repository.UserRepository.Save(user)
// 	require.Nil(t, err)

// 	// Check the User's role.
// 	role, err := repository.UserRepository.CheckLibrarian(user1.ID)
// 	require.Nil(t, err)
// 	require.Equal(t, role, "librarian")

// 	repository.CleanUp()
// }

// func TestUserAddFine(t *testing.T) {
// 	// Create a new User and save it.
// 	user := &user.User{
// 		ID: util.NewID(),
// 	}
// 	user1, err := repository.UserRepository.Save(user)
// 	require.Nil(t, err)

// 	// Add fine to the user.
// 	var fine uint32 = 7000
// 	err = repository.UserRepository.AddFine(user.ID, fine)
// 	require.Nil(t, err)

// 	// Get the User and check the fine amount.
// 	user2, err := repository.UserRepository.Get(user.ID)
// 	require.Nil(t, err)
// 	require.Equal(t, user2.ID, user1.ID)
// 	require.Equal(t, user2.TotalFine, fine)

// 	repository.CleanUp()
// }

// func TestUserGetTotalFine(t *testing.T) {

// }
