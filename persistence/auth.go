package persistence

import (
	"github.com/jmoiron/sqlx"

	"github.com/joshuabezaleel/library-server/pkg/auth"
)

type authRepository struct {
	DB *sqlx.DB
}

// NewAuthRepository returns initialized implementaions of
// Authentication service repository.
func NewAuthRepository(DB *sqlx.DB) auth.Repository {
	return &authRepository{
		DB: DB,
	}
}

func (repo *authRepository) GetPassword(username string) (string, error) {
	var password string

	res := repo.DB.QueryRow("SELECT password FROM users WHERE username=$1", username)
	err := res.Scan(&password)
	if err != nil {
		return "", err
	}

	return password, nil
}
