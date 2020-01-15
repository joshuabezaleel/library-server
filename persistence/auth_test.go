package persistence

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestGetPassword(t *testing.T) {
	tt := []struct {
		name     string
		username string
		password string
		err      bool
	}{
		{
			name:     "valid username and password",
			username: "username",
			password: "password",
			err:      false,
		},
		{
			name:     "invalid username",
			username: "anotherUsername",
			password: "",
			err:      true,
		},
	}

	// Asserting a valid username and password
	validUsername := tt[0].username
	validPassword := tt[0].password
	rows := sqlmock.NewRows([]string{"password"}).AddRow(validPassword)

	Mock.ExpectQuery("SELECT password FROM users WHERE username=?").
		WithArgs(validUsername).
		WillReturnRows(rows)

	// Tests
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			password, err := AuthTestingRepository.GetPassword(tc.username)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
			require.Equal(t, password, tc.password)
		})
	}
}
