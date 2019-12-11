package book

import (
	"time"

	"github.com/segmentio/ksuid"
)

// Service provides basic operations on Book domain model.
type Service interface {
	// CRUD operations.
	Create(book *Book) (*Book, error)
	Get(bookID string) (*Book, error)
	Update(book *Book) (*Book, error)
	Delete(bookID string) error

	// Other operations.
}

type service struct {
	bookRepository Repository
}

// NewBookService creates an instance of the service for the Book domain model
// with all of the necessary dependencies.
func NewBookService(bookRepository Repository) Service {
	return &service{
		bookRepository: bookRepository,
	}
}

func (s *service) Create(book *Book) (*Book, error) {
	newBook := NewBook(newBookID(), book.Title, book.Publisher, book.YearPublished, book.CallNumber, book.CoverPicture, book.ISBN, book.Collation, book.Edition, book.Description, book.LOCClassification, book.Subject, book.Author, book.Quantity, time.Now())

	return s.bookRepository.Save(newBook)
}

func (s *service) Get(bookID string) (*Book, error) {
	return s.bookRepository.Get(bookID)
}

func (s *service) Update(book *Book) (*Book, error) {
	return s.bookRepository.Update(book)
}

func (s *service) Delete(bookID string) error {
	return s.bookRepository.Delete(bookID)
}

func newBookID() string {
	return ksuid.New().String()
}
