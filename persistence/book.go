package persistence

import (
	"github.com/jmoiron/sqlx"

	"github.com/joshuabezaleel/library-server/pkg/core/book"
)

type bookRepository struct {
	DB *sqlx.DB
}

// NewBookRepository returns initialized implementations of the repository for
// Book domain model.
func NewBookRepository(DB *sqlx.DB) book.Repository {
	return &bookRepository{
		DB: DB,
	}
}

func (repo *bookRepository) Save(book *book.Book) (*book.Book, error) {
	_, err := repo.DB.NamedExec("INSERT INTO books (id, title, publisher, year_published, call_number, cover_picture, isbn, book_collation, edition, description, loc_classification, quantity, added_at) VALUES (:id, :title, :publisher, :year_published, :call_number, :cover_picture, :isbn, :book_collation, :edition, :description, :loc_classification, :quantity, :added_at)", book)

	if err != nil {
		return nil, err
	}

	return book, nil
}

func (repo *bookRepository) Get(bookID string) (*book.Book, error) {
	book := book.Book{}

	err := repo.DB.QueryRowx("SELECT * FROM books WHERE id=$1", bookID).StructScan(&book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (repo *bookRepository) Update(book *book.Book) (*book.Book, error) {
	_, err := repo.DB.NamedExec("UPDATE books SET title=:title, publisher=:publisher, year_published=:year_published, call_number=:call_number, cover_picture=:cover_picture, isbn=:isbn, book_collation=:book_collation, edition=:edition, description=:description, loc_classification=:loc_classification, author=:author, quantity=:quantity WHERE id=:id", book)
	if err != nil {
		return nil, err
	}

	updatedBook, err := repo.Get(book.ID)
	if err != nil {
		return nil, err
	}

	return updatedBook, nil
}

func (repo *bookRepository) Delete(bookID string) error {
	_, err := repo.DB.Exec("DELETE FROM books WHERE id=$1", bookID)

	if err != nil {
		return err
	}

	return nil
}

func (repo *bookRepository) GetSubjectIDs(subjects []string) ([]int64, error) {
	var subjectID int64
	var subjectIDs []int64

	for _, subject := range subjects {
		err := repo.DB.QueryRow("SELECT id FROM subjects WHERE subject=$1", subject).Scan(&subjectID)

		if err != nil {
			return nil, err
		}

		subjectIDs = append(subjectIDs, subjectID)
	}

	return subjectIDs, nil
}

func (repo *bookRepository) SaveBookSubjects(bookID string, subjectIDs []int64) error {
	for _, subjectID := range subjectIDs {
		_, err := repo.DB.Exec("INSERT INTO books_subjects (book_id, subject_id) VALUES ($1,$2)", bookID, subjectID)

		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *bookRepository) GetBookSubjectIDs(bookID string) ([]int64, error) {
	var subjectID int64
	var subjectIDs []int64

	rows, err := repo.DB.Query("SELECT subject_id FROM books_subjects WHERE book_id=$1", bookID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&subjectID)
		if err != nil {
			return nil, err
		}

		subjectIDs = append(subjectIDs, subjectID)
	}

	return subjectIDs, nil
}

func (repo *bookRepository) GetSubjectsByID(subjectIDs []int64) ([]string, error) {
	var subject string
	var subjects []string

	for _, subjectID := range subjectIDs {
		err := repo.DB.QueryRow("SELECT subject FROM subjects WHERE id=$1", subjectID).Scan(&subject)
		if err != nil {
			return nil, err
		}

		subjects = append(subjects, subject)
	}

	return subjects, nil
}
