package persistence

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
)

func TestBookCopySave(t *testing.T) {
	tt := []struct {
		name     string
		bookCopy *bookcopy.BookCopy
		err      bool
	}{
		{
			name: "save a valid book copy",
			bookCopy: &bookcopy.BookCopy{
				ID:        util.NewID(),
				Condition: "New",
			},
			err: false,
		},
		{
			name: "save an invalid book copy",
			bookCopy: &bookcopy.BookCopy{
				ID:        util.NewID(),
				Condition: "New",
			},
			err: true,
		},
	}

	// Assert a save for a valid BookCopy.
	validBookCopy := tt[0].bookCopy

	result := sqlmock.NewResult(1, 1)

	Mock.ExpectExec("INSERT INTO bookcopies").
		WithArgs(validBookCopy.ID, validBookCopy.Barcode, validBookCopy.BookID, validBookCopy.Condition, validBookCopy.AddedAt).
		WillReturnResult(result)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			newBookCopy, err := BookCopyTestingRepository.Save(tc.bookCopy)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
			require.Equal(t, tc.bookCopy.ID, newBookCopy.ID)
		})
	}
}

func TestBookCopyGet(t *testing.T) {
	tt := []struct {
		name     string
		bookCopy *bookcopy.BookCopy
		err      bool
	}{
		{
			name: "get a valid book copy",
			bookCopy: &bookcopy.BookCopy{
				ID:        util.NewID(),
				Condition: "New",
			},
			err: false,
		},
		{
			name: "get an invalid book copy",
			bookCopy: &bookcopy.BookCopy{
				ID:        util.NewID(),
				Condition: "New",
			},
			err: true,
		},
	}

	// Assert a get for a valid book copy.
	validBookCopy := tt[0].bookCopy

	rows := sqlmock.NewRows([]string{"id", "condition"}).
		AddRow(validBookCopy.ID, validBookCopy.Condition)

	Mock.ExpectQuery("SELECT (.+) FROM bookcopies").
		WithArgs(validBookCopy.ID).
		WillReturnRows(rows)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			newBookCopy, err := BookCopyTestingRepository.Get(tc.bookCopy.ID)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
			require.Equal(t, tc.bookCopy.ID, newBookCopy.ID)
		})
	}
}

func TestBookCopyUpdate(t *testing.T) {
	tt := []struct {
		name     string
		bookCopy *bookcopy.BookCopy
		err      bool
	}{
		{
			name: "update a valid book copy",
			bookCopy: &bookcopy.BookCopy{
				ID:        util.NewID(),
				Condition: "New",
			},
			err: false,
		},
		{
			name: "update an invalid book copy",
			bookCopy: &bookcopy.BookCopy{
				ID:        util.NewID(),
				Condition: "New",
			},
			err: true,
		},
	}

	// Assert an update for a valid Book Copy.
	validBookCopy := tt[0].bookCopy

	result := sqlmock.NewResult(1, 1)

	Mock.ExpectExec("UPDATE bookcopies SET").
		WithArgs(validBookCopy.Barcode, validBookCopy.BookID, validBookCopy.Condition, validBookCopy.ID).
		WillReturnResult(result)

	rows := sqlmock.NewRows([]string{"id", "condition"}).
		AddRow(validBookCopy.ID, validBookCopy.Condition)

	Mock.ExpectQuery("SELECT (.+) FROM bookcopies WHERE id=?").
		WithArgs(validBookCopy.ID).
		WillReturnRows(rows)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			updatedBookCopy, err := BookCopyTestingRepository.Update(tc.bookCopy)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
			require.Equal(t, tc.bookCopy.ID, updatedBookCopy.ID)
			require.Equal(t, tc.bookCopy.Condition, updatedBookCopy.Condition)
		})
	}
}

func TestBookCopyDelete(t *testing.T) {
	tt := []struct {
		name     string
		bookCopy *bookcopy.BookCopy
		err      bool
	}{
		{
			name: "delete a valid book copy",
			bookCopy: &bookcopy.BookCopy{
				ID:        util.NewID(),
				Condition: "New",
			},
			err: false,
		},
		{
			name: "delete an invalid book copy",
			bookCopy: &bookcopy.BookCopy{
				ID:        util.NewID(),
				Condition: "New",
			},
			err: true,
		},
	}

	// Assert a delete for a valid Book Copy.
	validBookCopy := tt[0].bookCopy

	result := sqlmock.NewResult(1, 1)

	Mock.ExpectExec("DELETE FROM bookcopies").
		WithArgs(validBookCopy.ID).
		WillReturnResult(result)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := BookCopyTestingRepository.Delete(tc.bookCopy.ID)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
		})
	}
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
