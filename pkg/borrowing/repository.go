package borrowing

// Repository provides access to the Borrowing store.
type Repository interface {
	// CRUD operations.
	Get(borrowID string) (*Borrow, error)

	// Other operations.
	Borrow(userID string, bookCopyID string) error
	Return(userID string, bookCopyID string) error
}
