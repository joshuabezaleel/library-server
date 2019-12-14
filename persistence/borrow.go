package persistence

import (
	"github.com/jmoiron/sqlx"

	"github.com/joshuabezaleel/library-server/pkg/borrowing"
)

type borrowRepository struct {
	DB *sqlx.DB
}

// NewBorrowRepository returns initialized implementation of the repository for
// Borrow domain model.
func NewBorrowRepository(DB *sqlx.DB) borrowing.Repository {
	return &borrowRepository{
		DB: DB,
	}
}

func (repo *borrowRepository) Borrow(userID string, bookCopyID string) error {
	return nil
}

func (repo *borrowRepository) Return(userID string, bookCopyID string) error {
	return nil
}
