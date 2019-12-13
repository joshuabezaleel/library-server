package bookcopy

// Repository provides access to the BookCopy store.
type Repository interface {
	// CRUD operations.
	Save(bookCopy *BookCopy) (*BookCopy, error)
	Get(bookCopyID string) (*BookCopy, error)
	Update(bookCopy *BookCopy) (*BookCopy, error)
	Delete(bookCopyID string) error

	// Other operations.
}
