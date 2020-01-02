package persistence

import (
	"testing"

	"github.com/jmoiron/sqlx"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

// func TestGetPassword(t *testing.T) {
// 	// Create a User with a password and save it.
// 	user := &user.User{
// 		ID:       util.NewID(),
// 		Username: "usernamefortesting2",
// 		Password: "passwordfortesting2",
// 	}
// 	user1, err := repository.UserRepository.Save(user)
// 	require.Nil(t, err)

// 	// Get the User's password
// 	password, err := repository.AuthRepository.GetPassword(user.Username)
// 	require.Nil(t, err)
// 	require.Equal(t, password, user1.Password)

// 	repository.CleanUp()
// }

func TestGetPasswordMock(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		require.Nil(t, err)
	}
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")

	tt := []struct {
		name     string
		username string
		password string
	}{
		{
			name:     "valid username and password",
			username: "username",
			password: "password",
		},
	}

	// username := "username"
	// password := "password"

	// rows := sqlmock.NewRows([]string{"password"}).
	// 	AddRow("password").
	// 	AddRow("")

	// mock.ExpectQuery("SELECT password FROM users WHERE username=?").
	// 	WithArgs(username).
	// 	WillReturnRows(rows)

	authRepository := NewAuthRepository(db)

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			rows := sqlmock.NewRows([]string{"password"}).AddRow(tc.password)

			mock.ExpectQuery("SELECT password FROM users WHERE username=?").
				WithArgs(tc.username).
				WillReturnRows(rows)

			password, err := authRepository.GetPassword(tc.username)
			require.Nil(t, err)
			require.Equal(t, password, tc.password)
		})
	}

	// rows := sqlmock.NewRows([]string{"password"}).AddRow("password")

	// mock.ExpectQuery("SELECT password FROM users WHERE username=?").
	// 	WithArgs("zaky")

	_, err = authRepository.GetPassword("joshua")
	require.NotNil(t, err)
	// require.Equal(t, err.Error(), "")
	// require.Equal(t, password, tc.password)

}
