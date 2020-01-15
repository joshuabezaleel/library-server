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

	// Assserting a valid Book
	validBook := tt[0].book

	result := sqlmock.NewResult(1, 1)

	Mock.ExpectExec("INSERT INTO books").
		WithArgs(validBook.ID, validBook.Title, validBook.Publisher, validBook.YearPublished, validBook.CallNumber, validBook.CoverPicture, validBook.ISBN, validBook.Collation, validBook.Edition, validBook.Description, validBook.LOCClassification, validBook.Subject, validBook.Author, validBook.Quantity, validBook.AddedAt).
		WillReturnResult(result)

	// Tests
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
	}{
		{
			name: "get a valid book",
			book: &book.Book{
				ID:    util.NewID(),
				Title: "testTitle",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			rows := sqlmock.NewRows([]string{"id", "title"}).
				AddRow(tc.book.ID, tc.book.Title)

			Mock.ExpectQuery("SELECT (.+) FROM books WHERE id=?").
				WithArgs(tc.book.ID).
				WillReturnRows(rows)

			newBook, err := BookTestingRepository.Get(tc.book.ID)
			require.Nil(t, err)
			require.Equal(t, tc.book.ID, newBook.ID)
		})
	}

	// Covering error
	_, err := BookTestingRepository.Get(util.NewID())
	require.NotNil(t, err)
}

func TestBookUpdate(t *testing.T) {
	tt := []struct {
		name string
		book *book.Book
	}{
		{
			name: "update a valid book",
			book: &book.Book{
				ID:    util.NewID(),
				Title: "testTitle",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := sqlmock.NewResult(1, 1)

			Mock.ExpectExec("UPDATE books SET").
				WithArgs(tc.book.Title, tc.book.Publisher, tc.book.YearPublished, tc.book.CallNumber, tc.book.CoverPicture, tc.book.ISBN, tc.book.Collation, tc.book.Edition, tc.book.Description, tc.book.LOCClassification, tc.book.Subject, tc.book.Author, tc.book.Quantity, tc.book.ID).
				WillReturnResult(result)

			rows := sqlmock.NewRows([]string{"id", "title"}).
				AddRow(tc.book.ID, tc.book.Title)

			Mock.ExpectQuery("SELECT (.+) FROM books WHERE id=?").
				WithArgs(tc.book.ID).
				WillReturnRows(rows)

			updatedBook, err := BookTestingRepository.Update(tc.book)

			require.Nil(t, err)
			require.Equal(t, tc.book.ID, updatedBook.ID)
			require.Equal(t, tc.book.Title, updatedBook.Title)
		})
	}

	// Covering error
	anotherBook := &book.Book{
		ID: util.NewID(),
	}
	_, err := BookTestingRepository.Update(anotherBook)
	require.NotNil(t, err)
}

func TestBookDelete(t *testing.T) {
	tt := []struct {
		name string
		book *book.Book
	}{
		{
			name: "delete a valid book",
			book: &book.Book{
				ID:    util.NewID(),
				Title: "testTitle",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := sqlmock.NewResult(1, 1)

			Mock.ExpectExec("DELETE FROM books").
				WithArgs(tc.book.ID).
				WillReturnResult(result)

			err := BookTestingRepository.Delete(tc.book.ID)

			require.Nil(t, err)
		})
	}

	// Covering error
	anotherBook := &book.Book{
		ID: util.NewID(),
	}
	err := BookTestingRepository.Delete(anotherBook.ID)
	require.NotNil(t, err)
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
