package persistence

import (
	"github.com/jmoiron/sqlx"

	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
)

type bookCopyRepository struct {
	DB *sqlx.DB
}

// NewBookCopyRepository returns initialized implementations of the repository for
// BookCopy domain model.
func NewBookCopyRepository(DB *sqlx.DB) bookcopy.Repository {
	return &bookCopyRepository{
		DB: DB,
	}
}

func (repo *bookCopyRepository) Save(bookCopy *bookcopy.BookCopy) (*bookcopy.BookCopy, error) {
	_, err := repo.DB.NamedExec("INSERT INTO bookcopies (id, barcode, book_id, condition, added_at) VALUES (:id, :barcode, :book_id, :condition, :added_at)", bookCopy)

	if err != nil {
		return nil, err
	}

	return bookCopy, nil
}

func (repo *bookCopyRepository) Get(bookCopyID string) (*bookcopy.BookCopy, error) {
	bookCopy := bookcopy.BookCopy{}

	err := repo.DB.QueryRowx("SELECT * FROM bookcopies WHERE id=$1", bookCopyID).StructScan(&bookCopy)
	if err != nil {
		return nil, err
	}

	return &bookCopy, nil
}

func (repo *bookCopyRepository) Update(bookCopy *bookcopy.BookCopy) (*bookcopy.BookCopy, error) {
	_, err := repo.DB.NamedExec("UPDATE bookcopies SET barcode=:barcode, book_id=:book_id, condition=:condition WHERE id=:id", bookCopy)

	if err != nil {
		return nil, err
	}

	// log.Println(bookCopy.ID)

	// var isExist bool

	// err = repo.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM bookcopies WHERE id=$1)", bookCopy.ID).Scan(&isExist)
	// if err != nil {
	// 	log.Println("C")
	// 	return nil, err
	// }
	// log.Printf("isExist= %v\n", isExist)

	updatedBookCopy, err := repo.Get(bookCopy.ID)
	if err != nil {
		return nil, err
	}

	return updatedBookCopy, nil
}

func (repo *bookCopyRepository) Delete(bookCopyID string) error {
	_, err := repo.DB.Exec("DELETE FROM bookcopies WHERE id=$1", bookCopyID)

	if err != nil {
		return err
	}

	return nil
}
