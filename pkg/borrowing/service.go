package borrowing

import (
	"errors"
	"time"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
)

const (
	finePerDay = 2000
)

// Service provides basic operations on Borrowing domain model.
type Service interface {
	Borrow(username string, bookCopyID string) (*Borrow, error)
	Get(borrowID string) (*Borrow, error)
	GetByUserIDAndBookCopyID(userID string, bookCopyID string) (*Borrow, error)
	CheckBorrowed(bookCopyID string) (bool, error)
	Return(username string, bookCopyID string) (*Borrow, error)
}

type service struct {
	borrowingRepository Repository
	userService         user.Service
	bookCopyService     bookcopy.Service
}

// NewBorrowingService creates an instance of the service for the Borrowing domain model
// with all of the necessary dependencies.
func NewBorrowingService(borrowingRepository Repository, userService user.Service, bookCopyService bookcopy.Service) Service {
	return &service{
		borrowingRepository: borrowingRepository,
		userService:         userService,
		bookCopyService:     bookCopyService,
	}
}

func (s *service) Borrow(username string, bookCopyID string) (*Borrow, error) {
	userID, err := s.userService.GetUserIDByUsername(username)
	if err != nil {
		return nil, err
	}

	// Check if Book Copy with the particular ID exists.
	if _, err := s.bookCopyService.Get(bookCopyID); err != nil {
		return nil, err
	}

	// Check if the particular Book Copy is being borrowed.
	isBorrowed, err := s.CheckBorrowed(bookCopyID)
	if err != nil {
		return nil, err
	}

	if isBorrowed {
		return nil, errors.New("Book " + bookCopyID + " is currently being borrowed")
	}

	newBorrow := NewBorrow(util.NewID(), userID, bookCopyID, 0, time.Now(), time.Now().AddDate(0, 0, 7), time.Time{})

	return s.borrowingRepository.Borrow(newBorrow)
}

func (s *service) Get(borrowID string) (*Borrow, error) {
	return s.borrowingRepository.Get(borrowID)
}

func (s *service) GetByUserIDAndBookCopyID(userID string, bookCopyID string) (*Borrow, error) {
	return s.borrowingRepository.GetByUserIDAndBookCopyID(userID, bookCopyID)
}

func (s *service) CheckBorrowed(bookCopyID string) (bool, error) {
	return s.borrowingRepository.CheckBorrowed(bookCopyID)
}

func (s *service) Return(username string, bookCopyID string) (*Borrow, error) {
	userID, err := s.userService.GetUserIDByUsername(username)
	if err != nil {
		return nil, err
	}

	borrow, err := s.GetByUserIDAndBookCopyID(userID, bookCopyID)
	if err != nil {
		return nil, err
	}

	borrow.ReturnedAt = time.Now()

	if borrow.ReturnedAt.After(borrow.DueDate) {
		diff := int(borrow.ReturnedAt.Sub(borrow.DueDate).Hours() / 24)
		borrow.Fine = uint32(diff * finePerDay)

		_, err = s.userService.AddFine(userID, borrow.Fine)
		if err != nil {
			return nil, err
		}
	}

	return s.borrowingRepository.Return(borrow)
}
