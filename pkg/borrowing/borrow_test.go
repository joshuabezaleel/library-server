package borrowing

import (
	"errors"
	"testing"
	"time"

	"github.com/bouk/monkey"
	"github.com/stretchr/testify/require"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
)

var userRepository = &user.MockRepository{}
var bookRepository = &book.MockRepository{}
var bookCopyRepository = &bookcopy.MockRepository{}
var borrowRepository = &MockRepository{}

var userService = user.NewUserService(userRepository)
var bookService = book.NewBookService(bookRepository)
var bookCopyService = bookcopy.NewBookCopyService(bookCopyRepository, bookService)
var borrowService = NewBorrowingService(borrowRepository, userService, bookCopyService)

func TestBorrow(t *testing.T) {
	createdTime := time.Now()
	timePatch := monkey.Patch(time.Now, func() time.Time {
		return createdTime
	})
	defer timePatch.Unpatch()

	dueDate := createdTime.AddDate(0, 0, 7)

	user := &user.User{
		ID:       util.NewID(),
		Username: "testUsername",
	}
	userRepository.On("GetIDByUsername", user.Username).Return(user.ID, nil)

	bookCopy := &bookcopy.BookCopy{
		ID:     util.NewID(),
		BookID: util.NewID(),
	}
	bookCopyRepository.On("Get", bookCopy.ID).Return(bookCopy, nil)

	borrowRepository.On("CheckBorrowed", bookCopy.ID).Return(false, nil)

	borrowID := util.NewID()
	borrowIDPatch := monkey.Patch(util.NewID, func() string {
		return borrowID
	})
	defer borrowIDPatch.Unpatch()

	borrow := &Borrow{
		ID:         borrowID,
		UserID:     user.ID,
		BookCopyID: bookCopy.ID,
		BorrowedAt: createdTime,
		DueDate:    dueDate,
	}
	borrowRepository.On("Borrow", borrow).Return(borrow, nil)
	newBorrow, err := borrowService.Borrow(user.Username, bookCopy.ID)

	require.Nil(t, err)
	require.Equal(t, user.ID, newBorrow.UserID)
	require.Equal(t, bookCopy.ID, newBorrow.BookCopyID)

	// Check for Book that is not borrowed.
	anotherBookCopy := &bookcopy.BookCopy{
		ID:     util.NewID(),
		BookID: util.NewID(),
	}
	bookCopyRepository.On("Get", anotherBookCopy.ID).Return(anotherBookCopy, nil)

	borrowRepository.On("CheckBorrowed", anotherBookCopy.ID).Return(true, nil)

	anotherBorrow, err := borrowService.Borrow(user.Username, anotherBookCopy.ID)

	require.Nil(t, anotherBorrow)
	require.Equal(t, err, errors.New("Book "+anotherBookCopy.ID+" is currently being borrowed"))
}

func TestGet(t *testing.T) {
	borrow := &Borrow{
		ID: util.NewID(),
	}
	borrowRepository.On("Get", borrow.ID).Return(borrow, nil)

	newBorrow, err := borrowService.Get(borrow.ID)

	require.Nil(t, err)
	require.Equal(t, borrow.ID, newBorrow.ID)
}

func TestGetByUserIDAndBookCopyID(t *testing.T) {
	borrow := &Borrow{
		ID:         util.NewID(),
		UserID:     util.NewID(),
		BookCopyID: util.NewID(),
	}
	borrowRepository.On("GetByUserIDAndBookCopyID", borrow.UserID, borrow.BookCopyID).Return(borrow, nil)

	newBorrow, err := borrowService.GetByUserIDAndBookCopyID(borrow.UserID, borrow.BookCopyID)

	require.Nil(t, err)
	require.Equal(t, borrow.ID, newBorrow.ID)
}

func TestCheckBorrowed(t *testing.T) {
	borrow := &Borrow{
		ID:         util.NewID(),
		BookCopyID: util.NewID(),
	}

	borrowRepository.On("CheckBorrowed", borrow.BookCopyID).Return(true, nil)

	isBorrowed, err := borrowService.CheckBorrowed(borrow.BookCopyID)

	require.Nil(t, err)
	require.True(t, isBorrowed)
}

func TestReturn(t *testing.T) {
	user := &user.User{
		ID:        util.NewID(),
		Username:  "username",
		TotalFine: uint32(0),
	}
	userRepository.On("GetIDByUsername", user.Username).Return(user.ID, nil)

	bookCopy := &bookcopy.BookCopy{
		ID: util.NewID(),
	}

	borrow := &Borrow{
		ID:         util.NewID(),
		UserID:     user.ID,
		BookCopyID: bookCopy.ID,
		DueDate:    time.Now().AddDate(0, 0, -7),
	}

	borrowRepository.On("GetByUserIDAndBookCopyID", user.ID, bookCopy.ID).Return(borrow, nil)

	expectedFine := uint32(14000)

	userRepository.On("GetTotalFine", user.ID).Return(user.TotalFine, nil)
	userRepository.On("AddFine", user.ID, expectedFine).Return(nil)

	borrowRepository.On("Return", borrow).Return(borrow, nil)

	returnedBorrow, err := borrowService.Return(user.Username, bookCopy.ID)

	require.Nil(t, err)
	require.Equal(t, borrow.ID, returnedBorrow.ID)
}
