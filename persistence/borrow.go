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

func (repo *borrowRepository) Borrow(borrow *borrowing.Borrow) (*borrowing.Borrow, error) {
	_, err := repo.DB.NamedExec("INSERT INTO borrows (id, user_id, bookcopy_id, fine, borrowed_at, due_date, returned_at) VALUES (:id, :user_id, :bookcopy_id, :fine, :borrowed_at, :due_date, :returned_at)", borrow)

	if err != nil {
		return nil, err
	}

	return borrow, nil
}

func (repo *borrowRepository) Get(borrowID string) (*borrowing.Borrow, error) {
	borrow := borrowing.Borrow{}

	err := repo.DB.QueryRowx("SELECT * FROM borrows WHERE id=$1", borrowID).StructScan(&borrow)
	if err != nil {
		return nil, err
	}

	return &borrow, nil
}

func (repo *borrowRepository) GetByUserIDAndBookCopyID(userID string, bookCopyID string) (*borrowing.Borrow, error) {
	borrow := borrowing.Borrow{}

	err := repo.DB.QueryRowx("SELECT * FROM borrows WHERE user_id=$1 AND bookcopy_id=$2", userID, bookCopyID).StructScan(&borrow)
	if err != nil {
		return nil, err
	}

	return &borrow, nil
}

func (repo *borrowRepository) CheckBorrowed(bookCopyID string) (bool, error) {
	var isBorrowed bool

	err := repo.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM borrows WHERE bookcopy_id=$1", bookCopyID).Scan(&isBorrowed)
	if err != nil {
		return false, err
	}

	return isBorrowed, nil
}

func (repo *borrowRepository) Return(borrow *borrowing.Borrow) (*borrowing.Borrow, error) {
	_, err := repo.DB.NamedExec("UPDATE borrows SET fine=:fine, returned_at=:returned_at WHERE id=:id", borrow)
	if err != nil {
		return nil, err
	}

	returnedBorrow, err := repo.Get(borrow.ID)
	if err != nil {
		return nil, err
	}

	return returnedBorrow, nil
}
