package bookcopy

import (
	"time"

	"github.com/joshuabezaleel/library-server/pkg/core/book"

	"github.com/segmentio/ksuid"
)

// Service provides basic operations on BookCopy domain model.
type Service interface {
	// CRUD operations.
	Create(bookCopy *BookCopy) (*BookCopy, error)
	Get(bookID string) (*BookCopy, error)
	Update(bookCopy *BookCopy) (*BookCopy, error)
	Delete(bookID string) error

	// Other operations.
}

type service struct {
	bookCopyRepository Repository
	bookService        book.Service
}

// NewBookCopyService creates an instance of the service for the BookCopy domain model
// with all of the necessary dependencies.
func NewBookCopyService(bookCopyRepository Repository, bookService book.Service) Service {
	return &service{
		bookCopyRepository: bookCopyRepository,
		bookService:        bookService,
	}
}

func (s *service) Create(bookCopy *BookCopy) (*BookCopy, error) {
	newBookCopy := NewBookCopy(newBookCopyID(), bookCopy.Barcode, bookCopy.BookID, bookCopy.Condition, time.Now())

	book, err := s.bookService.Get(bookCopy.BookID)
	if err != nil {
		return nil, err
	}

	book.Quantity++

	_, err = s.bookService.Update(book)
	if err != nil {
		return nil, err
	}

	return s.bookCopyRepository.Save(newBookCopy)
}

func (s *service) Get(bookCopyID string) (*BookCopy, error) {
	return s.bookCopyRepository.Get(bookCopyID)
}

func (s *service) Update(bookCopy *BookCopy) (*BookCopy, error) {
	return s.bookCopyRepository.Update(bookCopy)
}

func (s *service) Delete(bookCopyID string) error {
	return s.bookCopyRepository.Delete(bookCopyID)
}

func newBookCopyID() string {
	return ksuid.New().String()
}
