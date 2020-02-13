package book

import (
	"errors"
	"time"

	util "github.com/joshuabezaleel/library-server/pkg"
)

// Errors definition.
var (
	ErrCreateBook = errors.New("Error creating Book")
	ErrGetBook    = errors.New("Error retrieving Book")
	ErrUpdateBook = errors.New("Error updating Book")
	ErrDeleteBook = errors.New("Error deleting Book")

	ErrGetSubjectIDs     = errors.New("Error retrieving subject IDs")
	ErrSaveBookSubjects  = errors.New("Error saving Book's subjects")
	ErrGetBookSubjectIDs = errors.New("Error retrieving Book's subjects")
	ErrGetSubjectsByID   = errors.New("Error retrieving subjects")
)

// Service provides basic operations on Book domain model.
type Service interface {
	// CRUD operations.
	Create(book *Book) (*Book, error)
	Get(bookID string) (*Book, error)
	Update(book *Book) (*Book, error)
	Delete(bookID string) error

	// Other operations.
	GetSubjectIDs(subjects []string) ([]int64, error)
	SaveBookSubjects(bookID string, subjectIDs []int64) error
	GetBookSubjectIDs(bookID string) ([]int64, error)
	GetSubjectsByID(subjectIDs []int64) ([]string, error)
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
	var newBook *Book

	// Create a new instance of Book.
	if book.ID == "" {
		newBook = NewBook(util.NewID(), book.Title, book.Publisher, book.YearPublished, book.CallNumber, book.CoverPicture, book.ISBN, book.Collation, book.Edition, book.Description, book.LOCClassification, book.Author, book.Quantity, time.Now())
	} else {
		newBook = NewBook(book.ID, book.Title, book.Publisher, book.YearPublished, book.CallNumber, book.CoverPicture, book.ISBN, book.Collation, book.Edition, book.Description, book.LOCClassification, book.Author, book.Quantity, time.Now())
	}

	newBook, err := s.bookRepository.Save(newBook)
	if err != nil {
		return nil, ErrCreateBook
	}

	// Retrieve the IDs of the particular Book subjects that want to be created.
	subjectIDs, err := s.GetSubjectIDs(book.Subject)
	if err != nil {
		return nil, ErrGetSubjectIDs
	}

	// Save the relation between this BookID with all of the subjectIDs
	err = s.SaveBookSubjects(newBook.ID, subjectIDs)
	if err != nil {
		return nil, ErrSaveBookSubjects
	}

	createdBook, err := s.Get(newBook.ID)
	if err != nil {
		return nil, ErrGetBook
	}

	return createdBook, nil
}

func (s *service) Get(bookID string) (*Book, error) {
	book, err := s.bookRepository.Get(bookID)
	if err != nil {
		return nil, ErrGetBook
	}

	// Retrieve the IDs of the particular Book subjects that want to be retrieved.
	subjectIDs, err := s.GetBookSubjectIDs(bookID)
	if err != nil {
		return nil, ErrGetBookSubjectIDs
	}

	// Retrieve the Subjects by the IDs.
	subjects, err := s.GetSubjectsByID(subjectIDs)
	if err != nil {
		return nil, ErrGetSubjectsByID
	}

	book.Subject = subjects

	return book, nil
}

func (s *service) Update(book *Book) (*Book, error) {
	return s.bookRepository.Update(book)
}

func (s *service) Delete(bookID string) error {
	return s.bookRepository.Delete(bookID)
}

func (s *service) GetSubjectIDs(subjects []string) ([]int64, error) {
	return s.bookRepository.GetSubjectIDs(subjects)
}

func (s *service) SaveBookSubjects(bookID string, subjectIDs []int64) error {
	return s.bookRepository.SaveBookSubjects(bookID, subjectIDs)
}

func (s *service) GetBookSubjectIDs(bookID string) ([]int64, error) {
	return s.bookRepository.GetBookSubjectIDs(bookID)
}

func (s *service) GetSubjectsByID(subjectIDs []int64) ([]string, error) {
	return s.bookRepository.GetSubjectsByID(subjectIDs)
}
