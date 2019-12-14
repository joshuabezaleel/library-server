package borrowing

// Repository provides access to the Borrowing store.
type Repository interface {
	Borrow(userID string, bookCopyID string) error
	Return(userID string, bookCopyID string) error
}
