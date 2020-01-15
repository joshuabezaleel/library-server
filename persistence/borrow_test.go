package persistence

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/borrowing"
)

func TestBorrowBorrow(t *testing.T) {
	tt := []struct {
		name   string
		borrow *borrowing.Borrow
		err    bool
	}{
		{
			name: "valid borrow",
			borrow: &borrowing.Borrow{
				ID:         util.NewID(),
				UserID:     util.NewID(),
				BookCopyID: util.NewID(),
			},
			err: false,
		},
		{
			name: "invalid borrow",
			borrow: &borrowing.Borrow{
				ID:         util.NewID(),
				UserID:     util.NewID(),
				BookCopyID: util.NewID(),
			},
			err: true,
		},
	}

	// Assert a valid Borrow.
	validBorrow := tt[0].borrow

	result := sqlmock.NewResult(1, 1)

	Mock.ExpectExec("INSERT INTO borrows").
		WithArgs(validBorrow.ID, validBorrow.UserID, validBorrow.BookCopyID, validBorrow.Fine, validBorrow.BorrowedAt, validBorrow.DueDate, validBorrow.ReturnedAt).
		WillReturnResult(result)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			newBorrow, err := BorrowTestingRepository.Borrow(tc.borrow)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
			require.Equal(t, tc.borrow.UserID, newBorrow.UserID)
			require.Equal(t, tc.borrow.BookCopyID, newBorrow.BookCopyID)
		})
	}
}

func TestBorrowGet(t *testing.T) {
	tt := []struct {
		name   string
		borrow *borrowing.Borrow
		err    bool
	}{
		{
			name: "get a valid borrow",
			borrow: &borrowing.Borrow{
				ID:         util.NewID(),
				UserID:     util.NewID(),
				BookCopyID: util.NewID(),
			},
			err: false,
		},
		{
			name: "get an invalid borrow",
			borrow: &borrowing.Borrow{
				ID:         util.NewID(),
				UserID:     util.NewID(),
				BookCopyID: util.NewID(),
			},
			err: true,
		},
	}

	// Assert a get for a valid Borrow.
	validBorrow := tt[0].borrow

	rows := sqlmock.NewRows([]string{"id", "user_id", "bookcopy_id"}).
		AddRow(validBorrow.ID, validBorrow.UserID, validBorrow.BookCopyID)

	Mock.ExpectQuery("SELECT (.+) FROM borrows").
		WithArgs(validBorrow.ID).
		WillReturnRows(rows)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			newBorrow, err := BorrowTestingRepository.Get(tc.borrow.ID)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
			require.Equal(t, tc.borrow.ID, newBorrow.ID)
			require.Equal(t, tc.borrow.UserID, newBorrow.UserID)
			require.Equal(t, tc.borrow.BookCopyID, newBorrow.BookCopyID)
		})
	}
}

func TestBorrowGetByUserIDAndBookCopyID(t *testing.T) {
	tt := []struct {
		name   string
		borrow *borrowing.Borrow
		err    bool
	}{
		{
			name: "get a valid borrow",
			borrow: &borrowing.Borrow{
				ID:         util.NewID(),
				UserID:     util.NewID(),
				BookCopyID: util.NewID(),
			},
			err: false,
		},
		{
			name: "get an invalid borrow",
			borrow: &borrowing.Borrow{
				ID:         util.NewID(),
				UserID:     util.NewID(),
				BookCopyID: util.NewID(),
			},
			err: true,
		},
	}

	// Assert a get for a valid Borrow.
	validBorrow := tt[0].borrow

	rows := sqlmock.NewRows([]string{"id", "user_id", "bookcopy_id"}).
		AddRow(validBorrow.ID, validBorrow.UserID, validBorrow.BookCopyID)

	Mock.ExpectQuery("SELECT (.+) FROM borrows").
		WithArgs(validBorrow.UserID, validBorrow.BookCopyID).
		WillReturnRows(rows)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			newBorrow, err := BorrowTestingRepository.GetByUserIDAndBookCopyID(tc.borrow.UserID, tc.borrow.BookCopyID)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
			require.Equal(t, tc.borrow.ID, newBorrow.ID)
			require.Equal(t, tc.borrow.UserID, newBorrow.UserID)
			require.Equal(t, tc.borrow.BookCopyID, newBorrow.BookCopyID)
		})
	}
}

func TestBorrowCheckBorrowed(t *testing.T) {
	tt := []struct {
		name       string
		bookCopyID string
		borrowed   bool
		err        bool
	}{
		{
			name:       "borrowed book",
			bookCopyID: util.NewID(),
			borrowed:   true,
			err:        false,
		},
		{
			name:       "unborrowed book",
			bookCopyID: util.NewID(),
			borrowed:   false,
			err:        true,
		},
	}

	// Assert a valid checkBorrowed.
	validBookCopyID := tt[0].bookCopyID
	validBorrowed := tt[0].borrowed

	rows := sqlmock.NewRows([]string{"is_borrowed"}).
		AddRow(validBorrowed)

	Mock.ExpectQuery("SELECT (.+)").
		WithArgs(validBookCopyID).
		WillReturnRows(rows)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			borrowed, err := BorrowTestingRepository.CheckBorrowed(tc.bookCopyID)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
			require.Equal(t, tc.borrowed, borrowed)
		})
	}
}

func TestBorrowReturn(t *testing.T) {
	tt := []struct {
		name   string
		borrow *borrowing.Borrow
		err    bool
	}{
		{
			name: "return a valid borrow",
			borrow: &borrowing.Borrow{
				ID:         util.NewID(),
				UserID:     util.NewID(),
				BookCopyID: util.NewID(),
			},
			err: false,
		},
		{
			name: "return an invalid borrow",
			borrow: &borrowing.Borrow{
				ID:         util.NewID(),
				UserID:     util.NewID(),
				BookCopyID: util.NewID(),
			},
			err: true,
		},
	}

	// Assert a return for a valid Borrow.
	validBorrow := tt[0].borrow

	result := sqlmock.NewResult(1, 1)

	Mock.ExpectExec("UPDATE borrows SET").
		WithArgs(validBorrow.Fine, validBorrow.ReturnedAt, validBorrow.ID).
		WillReturnResult(result)

	rows := sqlmock.NewRows([]string{"id", "user_id", "bookcopy_id"}).
		AddRow(validBorrow.ID, validBorrow.UserID, validBorrow.BookCopyID)

	Mock.ExpectQuery("SELECT (.+) FROM borrows WHERE id=?").
		WithArgs(validBorrow.ID).
		WillReturnRows(rows)

	// Tests.
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			returnedBorrow, err := BorrowTestingRepository.Return(tc.borrow)

			if tc.err {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
			require.Equal(t, tc.borrow.ID, returnedBorrow.ID)
			require.Equal(t, tc.borrow.UserID, returnedBorrow.UserID)
			require.Equal(t, tc.borrow.BookCopyID, returnedBorrow.BookCopyID)
		})
	}
}

// func TestBorrow(t *testing.T) {
// 	// Create a new Borrow instance and Borrow it.
// 	borrow := &borrowing.Borrow{
// 		ID: util.NewID(),
// 	}
// 	borrow1, err := repository.BorrowRepository.Borrow(borrow)

// 	// Happy path.
// 	require.Nil(t, err)
// 	require.Equal(t, borrow.ID, borrow1.ID)

// 	repository.CleanUp()
// }

// func TestBorrowGetByID(t *testing.T) {
// 	// Create a new Borrow instance and Borrow it.
// 	borrow := &borrowing.Borrow{
// 		ID: util.NewID(),
// 	}
// 	borrow1, err := repository.BorrowRepository.Borrow(borrow)
// 	require.Nil(t, err)

// 	// Get by ID.
// 	borrow2, err := repository.BorrowRepository.Get(borrow.ID)
// 	require.Nil(t, err)
// 	require.Equal(t, borrow2.ID, borrow1.ID)

// 	// Get invalid Borrow.
// 	_, err = repository.BorrowRepository.Get(util.NewID())
// 	require.NotNil(t, err)

// 	repository.CleanUp()
// }

// func TestBorrowGetByUserIDAndBookCopyID(t *testing.T) {
// 	// Create a new Borrow instance with UserID and BookCopyID and Borrow it.
// 	borrow := &borrowing.Borrow{
// 		ID:         util.NewID(),
// 		UserID:     util.NewID(),
// 		BookCopyID: util.NewID(),
// 	}
// 	borrow1, err := repository.BorrowRepository.Borrow(borrow)
// 	require.Nil(t, err)

// 	// Get by UserID and BookCopyID.
// 	borrow2, err := repository.BorrowRepository.GetByUserIDAndBookCopyID(borrow.UserID, borrow.BookCopyID)
// 	require.Nil(t, err)
// 	require.Equal(t, borrow2.ID, borrow1.ID)

// 	// Get invalid Borrow.
// 	_, err = repository.BorrowRepository.GetByUserIDAndBookCopyID(util.NewID(), util.NewID())
// 	require.NotNil(t, err)

// 	repository.CleanUp()
// }

// func TestCheckBorrowed(t *testing.T) {
// 	// Create a new Borrow instance and Borrow it.
// 	borrow := &borrowing.Borrow{
// 		ID:         util.NewID(),
// 		UserID:     util.NewID(),
// 		BookCopyID: util.NewID(),
// 	}
// 	borrow1, err := repository.BorrowRepository.Borrow(borrow)
// 	require.Nil(t, err)

// 	// Check if the BookCopy is being borrowed.
// 	borrowed, err := repository.BorrowRepository.CheckBorrowed(borrow1.BookCopyID)
// 	t.Log(borrowed)
// 	require.Nil(t, err)
// 	require.True(t, borrowed)

// 	// Check for another BookCopy that is not being borrowed.
// 	borrowed1, err := repository.BorrowRepository.CheckBorrowed(util.NewID())
// 	require.Nil(t, err)
// 	require.False(t, borrowed1)

// 	repository.CleanUp()
// }

// func TestReturn(t *testing.T) {
// 	// Create a new Borrow instance and Borrow it.
// 	borrow := &borrowing.Borrow{
// 		ID:         util.NewID(),
// 		BookCopyID: util.NewID(),
// 	}
// 	borrow1, err := repository.BorrowRepository.Borrow(borrow)
// 	require.Nil(t, err)

// 	// Return the borrowed BookCopy.
// 	returnedBorrow, err := repository.BorrowRepository.Return(borrow1)
// 	require.Nil(t, err)
// 	require.Equal(t, returnedBorrow.ID, borrow.ID)

// 	repository.CleanUp()
// }
