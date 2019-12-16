package borrowing

import (
	"time"

	"github.com/joshuabezaleel/library-server/pkg/core/user"

	"github.com/segmentio/ksuid"
)

const (
	finePerDay = 2000
)

// Service provides basic operations on Borrowing domain model.
type Service interface {
	Borrow(userID string, bookCopyID string) (*Borrow, error)
	Get(borrowID string) (*Borrow, error)
	GetByUserIDAndBookCopyID(userID string, bookCopyID string) (*Borrow, error)
	Return(userID string, bookCopyID string) (*Borrow, error)
}

type service struct {
	borrowingRepository Repository
	userService         user.Service
}

// NewBorrowingService creates an instance of the service for the Borrowing domain model
// with all of the necessary dependencies.
func NewBorrowingService(borrowingRepository Repository, userService user.Service) Service {
	return &service{
		borrowingRepository: borrowingRepository,
		userService:         userService,
	}
}

func (s *service) Borrow(userID string, bookCopyID string) (*Borrow, error) {
	newBorrow := NewBorrow(newBorrowID(), userID, bookCopyID, 0, time.Now(), time.Now().AddDate(0, 0, 7), time.Time{})

	return s.borrowingRepository.Borrow(newBorrow)
}

func (s *service) Get(borrowID string) (*Borrow, error) {
	return s.borrowingRepository.Get(borrowID)
}

func (s *service) GetByUserIDAndBookCopyID(userID string, bookCopyID string) (*Borrow, error) {
	return s.borrowingRepository.GetByUserIDAndBookCopyID(userID, bookCopyID)
}

func (s *service) Return(userID string, bookCopyID string) (*Borrow, error) {
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

func newBorrowID() string {
	return ksuid.New().String()
}
