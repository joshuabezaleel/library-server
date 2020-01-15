package persistence

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
)

func TestBookCopySave(t *testing.T) {
	tt := []struct {
		name     string
		bookCopy *bookcopy.BookCopy
	}{
		{
			name: "valid book",
			bookCopy: &bookcopy.BookCopy{
				ID:        util.NewID(),
				Condition: "New",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := sqlmock.NewResult(1, 1)

			Mock.ExpectExec("INSERT INTO bookcopies").
				WithArgs(tc.bookCopy.ID, tc.bookCopy.Barcode, tc.bookCopy.BookID, tc.bookCopy.Condition, tc.bookCopy.AddedAt).
				WillReturnResult(result)

			newBookCopy, err := BookCopyTestingRepository.Save(tc.bookCopy)

			require.Nil(t, err)
			require.Equal(t, tc.bookCopy.ID, newBookCopy.ID)
		})
	}

	// Covering error
	anotherBookCopy := &bookcopy.BookCopy{
		ID: util.NewID(),
	}
	_, err := BookCopyTestingRepository.Save(anotherBookCopy)
	require.NotNil(t, err)
}

func TestBookCopyGet(t *testing.T) {
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

func TestBookCopyUpdate(t *testing.T) {

}

func TestBookCopyDelete(t *testing.T) {

}

// func TestBookCopySave(t *testing.T) {
// 	// Create a new BookCopy and save it.
// 	bookCopy := &bookcopy.BookCopy{
// 		ID: util.NewID(),
// 	}
// 	bookCopy1, err := repository.BookCopyRepository.Save(bookCopy)

// 	// Happy path.
// 	require.Nil(t, err)
// 	require.Equal(t, bookCopy.ID, bookCopy1.ID)

// 	repository.CleanUp()
// }

// func TestBookCopyGet(t *testing.T) {
// 	// Create a new BookCopy and save it.
// 	bookCopy := &bookcopy.BookCopy{
// 		ID: util.NewID(),
// 	}
// 	bookCopy1, err := repository.BookCopyRepository.Save(bookCopy)
// 	require.Nil(t, err)

// 	// Get the BookCopy.
// 	bookCopy2, err := repository.BookCopyRepository.Get(bookCopy.ID)
// 	require.Nil(t, err)
// 	require.Equal(t, bookCopy1.ID, bookCopy2.ID)

// 	// Get invalid BookCopy.
// 	_, err = repository.BookCopyRepository.Get(util.NewID())
// 	require.NotNil(t, err)

// 	repository.CleanUp()
// }

// func TestBookCopyUpdate(t *testing.T) {
// 	// Create a new BookCopy and save it.
// 	bookCopy := &bookcopy.BookCopy{
// 		ID:        util.NewID(),
// 		Condition: "good",
// 	}
// 	bookCopy1, err := repository.BookCopyRepository.Save(bookCopy)
// 	require.Nil(t, err)

// 	// Update BookCopy's Condition
// 	bookCopy1.Condition = "bad"
// 	bookCopy2, err := repository.BookCopyRepository.Update(bookCopy1)
// 	require.Nil(t, err)
// 	require.Equal(t, bookCopy1.ID, bookCopy2.ID)
// 	require.Equal(t, bookCopy2.Condition, "bad")

// 	repository.CleanUp()
// }

// func TestBookCopyDelete(t *testing.T) {
// 	// Create a new BookCopy and save it.
// 	bookCopy := &bookcopy.BookCopy{
// 		ID: util.NewID(),
// 	}
// 	_, err := repository.BookCopyRepository.Save(bookCopy)
// 	require.Nil(t, err)

// 	// Delete the BookCopy that was just created.
// 	err = repository.BookCopyRepository.Delete(bookCopy.ID)
// 	require.Nil(t, err)

// 	// Unable to retrieve the BookCopy that was just deleted.
// 	_, err = repository.BookCopyRepository.Get(bookCopy.ID)
// 	require.NotNil(t, err)

// 	repository.CleanUp()
// }
