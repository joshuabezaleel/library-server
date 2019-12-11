package book

// Repository provides access to the Book store.
type Repository interface {
	// CRUD operations.
	Save(book *Book) (*Book, error)
	Get(bookID string) (*Book, error)
	Update(book *Book) (*Book, error)
	Delete(bookID string) error

	// Other operations.
}
