package book

// Repository provides access to the Book store.
type Repository interface {
	// CRUD operations.
	Save(book *Book) (*Book, error)
	Get(bookID string) (*Book, error)
	Update(book *Book) (*Book, error)
	Delete(bookID string) error

	// Other operations.
	GetSubjectIDs(subjects []string) ([]int64, error)
	SaveBookSubjects(bookID string, subjectIDs []int64) error
	GetBookSubjectIDs(bookID string) ([]int64, error)
	GetSubjectsByID(subjectIDs []int64) ([]string, error)
}
