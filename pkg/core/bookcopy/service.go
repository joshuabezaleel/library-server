package bookcopy

import (
	"errors"
	"time"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
)

// Errors definition.
var (
	ErrCreateBookCopy = errors.New("Error creating Book Copy")
	ErrGetBookCopy    = errors.New("Error retrieving Book Copy")
	ErrUpdateBookCopy = errors.New("Error updating Book Copy")
	ErrDeleteBookCopy = errors.New("Error deleting Book Copy")
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
	var newBookCopy *BookCopy

	if bookCopy.ID == "" {
		newBookCopy = NewBookCopy(util.NewID(), bookCopy.Barcode, bookCopy.BookID, bookCopy.Condition, time.Now())
	} else {
		newBookCopy = NewBookCopy(bookCopy.ID, bookCopy.Barcode, bookCopy.BookID, bookCopy.Condition, time.Now())
	}

	newBookCopy, err := s.bookCopyRepository.Save(newBookCopy)
	if err != nil {
		return nil, ErrCreateBookCopy
	}

	book, err := s.bookService.Get(bookCopy.BookID)
	if err != nil {
		return nil, ErrGetBookCopy
	}

	book.Quantity++

	_, err = s.bookService.Update(book)
	if err != nil {
		return nil, ErrUpdateBookCopy
	}

	return newBookCopy, nil
}

func (s *service) Get(bookCopyID string) (*BookCopy, error) {
	bookCopy, err := s.bookCopyRepository.Get(bookCopyID)
	if err != nil {
		return nil, ErrGetBookCopy
	}

	return bookCopy, nil
}

func (s *service) Update(bookCopy *BookCopy) (*BookCopy, error) {
	bookCopy, err := s.bookCopyRepository.Update(bookCopy)
	if err != nil {
		return nil, ErrUpdateBookCopy
	}

	return bookCopy, nil
}

func (s *service) Delete(bookCopyID string) error {
	err := s.bookCopyRepository.Delete(bookCopyID)
	if err != nil {
		return ErrDeleteBookCopy
	}

	return nil
}
