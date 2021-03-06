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
		WithArgs(validBook.ID, validBook.Title, validBook.Publisher, validBook.YearPublished, validBook.CallNumber, validBook.CoverPicture, validBook.ISBN, validBook.Collation, validBook.Edition, validBook.Description, validBook.LOCClassification, validBook.Quantity, validBook.AddedAt).
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
		WithArgs(validBook.Title, validBook.Publisher, validBook.YearPublished, validBook.CallNumber, validBook.CoverPicture, validBook.ISBN, validBook.Collation, validBook.Edition, validBook.Description, validBook.LOCClassification, validBook.Quantity, validBook.ID).
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
