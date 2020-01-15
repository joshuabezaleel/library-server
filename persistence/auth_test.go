package persistence

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestGetPassword(t *testing.T) {
	// mockDB, mock, err := sqlmock.New()
	// if err != nil {
	// 	require.Nil(t, err)
	// }
	// defer mockDB.Close()
	// db := sqlx.NewDb(mockDB, "sqlmock")

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

	// authRepository := NewAuthRepository(db)

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			rows := sqlmock.NewRows([]string{"password"}).AddRow(tc.password)

			Mock.ExpectQuery("SELECT password FROM users WHERE username=?").
				WithArgs(tc.username).
				WillReturnRows(rows)

			password, err := AuthTestingRepository.GetPassword(tc.username)
			require.Nil(t, err)
			require.Equal(t, password, tc.password)
		})
	}

	_, err := AuthTestingRepository.GetPassword("joshua")
	require.NotNil(t, err)
}
