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

}

func TestBorrowGetByUserIDAndBookCopyID(t *testing.T) {

}

func TestBorrowCheckBorrowed(t *testing.T) {

}

func TestBorrowReturn(t *testing.T) {

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
