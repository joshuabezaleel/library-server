package borrowing

// Repository provides access to the Borrowing store.
type Repository interface {
	Borrow(borrow *Borrow) (*Borrow, error)
	Get(borrowID string) (*Borrow, error)
	GetByUserIDAndBookCopyID(userID string, bookCopyID string) (*Borrow, error)
	CheckBorrowed(bookCopyID string) (bool, error)
	Return(borrow *Borrow) (*Borrow, error)
}
