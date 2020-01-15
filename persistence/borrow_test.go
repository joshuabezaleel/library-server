package persistence

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
