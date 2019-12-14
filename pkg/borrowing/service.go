package borrowing

// Service provides basic operations on Borrowing domain model.
type Service interface {
	Borrow(userID string, bookCopyID string) error
	Return(userID string, bokoCopyID string) error
}

type service struct {
	borrowingRepository Repository
}

// NewBorrowingService creates an instance of the service for the Borrowing domain model
// with all of the necessary dependencies.
func NewBorrowingService(borrowingRepository Repository) Service {
	return &service{
		borrowingRepository: borrowingRepository,
	}
}

func (s *service) Borrow(userID string, bookCopyID string) error {
	return nil
}

func (s *service) Return(userID string, bookCopyID string) error {
	return nil
}
