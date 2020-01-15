package persistence

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
)

func TestBookSave(t *testing.T) {
	tt := []struct {
		name string
		book *book.Book
		err  bool
	}{
		{
			name: "save a valid book",
			book: &book.Book{
				ID:    util.NewID(),
				Title: "testTitle",
			},
			err: false,
		},
		{
			name: "save an invalid book",
			book: &book.Book{
				ID:    util.NewID(),
				Title: "anotherTestTitle",
			},
			err: true,
		},
	}

	// Asssert a save for a valid Book.
	validBook := tt[0].book

	result := sqlmock.NewResult(1, 1)

	Mock.ExpectExec("INSERT INTO books").
		WithArgs(validBook.ID, validBook.Title, validBook.Publisher, validBook.YearPublished, validBook.CallNumber, validBook.CoverPicture, validBook.ISBN, validBook.Collation, validBook.Edition, validBook.Description, validBook.LOCClassification, validBook.Subject, validBook.Author, validBook.Quantity, validBook.AddedAt).
		WillReturnResult(result)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			newBook, err := BookTestingRepository.Save(tc.book)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
			require.Equal(t, tc.book.ID, newBook.ID)
		})
	}
}

func TestBookGet(t *testing.T) {
	tt := []struct {
		name string
		book *book.Book
		err  bool
	}{
		{
			name: "get a valid book",
			book: &book.Book{
				ID:    util.NewID(),
				Title: "testTitle",
			},
			err: false,
		},
		{
			name: "get an invalid book",
			book: &book.Book{
				ID:    util.NewID(),
				Title: "anotherTestTitle",
			},
			err: true,
		},
	}

	// Assert a get for a valid Book.
	validBook := tt[0].book

	rows := sqlmock.NewRows([]string{"id", "title"}).
		AddRow(validBook.ID, validBook.Title)

	Mock.ExpectQuery("SELECT (.+) FROM books WHERE id=?").
		WithArgs(validBook.ID).
		WillReturnRows(rows)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			newBook, err := BookTestingRepository.Get(tc.book.ID)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
			require.Equal(t, tc.book.ID, newBook.ID)
		})
	}
}

func TestBookUpdate(t *testing.T) {
	tt := []struct {
		name string
		book *book.Book
		err  bool
	}{
		{
			name: "update a valid book",
			book: &book.Book{
				ID:    util.NewID(),
				Title: "testTitle",
			},
			err: false,
		},
		{
			name: "update an invalid book",
			book: &book.Book{
				ID:    util.NewID(),
				Title: "anotherTestTitle",
			},
			err: true,
		},
	}

	// Assert an update for a valid Book.
	validBook := tt[0].book

	result := sqlmock.NewResult(1, 1)

	Mock.ExpectExec("UPDATE books SET").
		WithArgs(validBook.Title, validBook.Publisher, validBook.YearPublished, validBook.CallNumber, validBook.CoverPicture, validBook.ISBN, validBook.Collation, validBook.Edition, validBook.Description, validBook.LOCClassification, validBook.Subject, validBook.Author, validBook.Quantity, validBook.ID).
		WillReturnResult(result)

	rows := sqlmock.NewRows([]string{"id", "title"}).
		AddRow(validBook.ID, validBook.Title)

	Mock.ExpectQuery("SELECT (.+) FROM books WHERE id=?").
		WithArgs(validBook.ID).
		WillReturnRows(rows)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			updatedBook, err := BookTestingRepository.Update(tc.book)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
			require.Equal(t, tc.book.ID, updatedBook.ID)
			require.Equal(t, tc.book.Title, updatedBook.Title)
		})
	}
}

func TestBookDelete(t *testing.T) {
	tt := []struct {
		name string
		book *book.Book
		err  bool
	}{
		{
			name: "delete a valid book",
			book: &book.Book{
				ID:    util.NewID(),
				Title: "testTitle",
			},
			err: false,
		},
		{
			name: "delete an invalid book",
			book: &book.Book{
				ID:    util.NewID(),
				Title: "anotherTestTitle",
			},
			err: true,
		},
	}

	// Assert a delete for a valid Book.
	validBook := tt[0].book

	result := sqlmock.NewResult(1, 1)

	Mock.ExpectExec("DELETE FROM books").
		WithArgs(validBook.ID).
		WillReturnResult(result)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := BookTestingRepository.Delete(tc.book.ID)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
		})
	}
}

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
