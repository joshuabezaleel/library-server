package persistence

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	// "github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
)

// func TestBookSave(t *testing.T) {
// 	// Create a new Book and save it.
// 	book := &book.Book{
// 		ID:         util.NewID(),
// 		CallNumber: util.NewID(),
// 	}
// 	newBook, err := repository.BookRepository.Save(book)

// 	// Happy path.
// 	require.Nil(t, err)
// 	require.Equal(t, book.ID, newBook.ID)

// 	repository.CleanUp()
// }
func TestBookSave(t *testing.T) {
	tt := []struct {
		name string
		book *book.Book
	}{
		{
			name: "valid book",
			book: &book.Book{
				ID:    util.NewID(),
				Title: "testTitle",
			},
		},
	}

	// bookRepository := NewBookRepository(db)

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := sqlmock.NewResult(1, 1)
			// rows := sqlmock.NewRows([]string{"id", "title"}).
			// 	AddRow(tc.book.ID, tc.book.Title)

			Mock.ExpectExec("INSERT INTO books").
				WithArgs(tc.book.ID, tc.book.Title, tc.book.Publisher, tc.book.YearPublished, tc.book.CallNumber, tc.book.CoverPicture, tc.book.ISBN, tc.book.Collation, tc.book.Edition, tc.book.Description, tc.book.LOCClassification, tc.book.Subject, tc.book.Author, tc.book.Quantity, tc.book.AddedAt).
				WillReturnResult(result)

			newBook, err := BookTestingRepository.Save(tc.book)

			require.Nil(t, err)
			require.Equal(t, newBook.ID, tc.book.ID)
		})
	}

	anotherBook := &book.Book{
		ID: util.NewID(),
	}
	_, err := BookTestingRepository.Save(anotherBook)
	require.NotNil(t, err)
}

func TestBookGet(t *testing.T) {
	// mockDB, mock, err := sqlmock.New()
	// if err != nil {
	// 	require.Nil(t, err)
	// }
	// defer mockDB.Close()
	// db := sqlx.NewDb(mockDB, "sqlmock")

	tt := []struct {
		name string
		book *book.Book
	}{
		{
			name: "valid book",
			book: &book.Book{
				ID:    util.NewID(),
				Title: "testTitle",
			},
		},
	}

	// bookRepository := NewBookRepository(db)

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			rows := sqlmock.NewRows([]string{"id", "title"}).
				AddRow(tc.book.ID, tc.book.Title)

			Mock.ExpectQuery("SELECT (.+) FROM books WHERE id=?").
				WithArgs(tc.book.ID).
				WillReturnRows(rows)

			newBook, err := BookTestingRepository.Get(tc.book.ID)
			require.Nil(t, err)
			require.Equal(t, newBook.ID, tc.book.ID)
		})

	}
}

// func TestBookGet(t *testing.T) {
// 	// Create a new Book and save it.
// 	book := &book.Book{
// 		ID: util.NewID(),
// 	}
// 	book1, err := repository.BookRepository.Save(book)
// 	require.Nil(t, err)

// 	// Get the Book.
// 	book2, err := repository.BookRepository.Get(book.ID)
// 	require.Nil(t, err)
// 	require.Equal(t, book2.ID, book1.ID)

// 	// Get invalid Book.
// 	_, err = repository.BookRepository.Get(util.NewID())
// 	require.NotNil(t, err)

// 	repository.CleanUp()
// }

// func TestBookUpdate(t *testing.T) {
// 	// Create a new Book and save it.
// 	book := &book.Book{
// 		ID:    util.NewID(),
// 		Title: "title",
// 	}
// 	book1, err := repository.BookRepository.Save(book)
// 	require.Nil(t, err)

// 	// Update the Book's title.
// 	book1.Title = "edited title"
// 	book2, err := repository.BookRepository.Update(book1)
// 	require.Nil(t, err)
// 	require.Equal(t, book1.ID, book2.ID)
// 	require.Equal(t, book2.Title, "edited title")

// 	repository.CleanUp()
// }

// func TestBookDelete(t *testing.T) {
// 	// Create a new Book and save it.
// 	book := &book.Book{
// 		ID:    util.NewID(),
// 		Title: "title",
// 	}
// 	_, err := repository.BookRepository.Save(book)
// 	require.Nil(t, err)

// 	// Delete the Book that was just created.
// 	err = repository.BookRepository.Delete(book.ID)
// 	require.Nil(t, err)

// 	// Unable to retrieve the Book that was just deleted.
// 	_, err = repository.BookRepository.Get(book.ID)
// 	require.NotNil(t, err)

// 	repository.CleanUp()
// }
