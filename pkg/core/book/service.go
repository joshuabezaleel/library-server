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

	ErrSaveAuthors      = errors.New("Error saving authors")
	ErrGetAuthorIDs     = errors.New("Error retrieving author IDs")
	ErrSaveBookAuthors  = errors.New("Error saving Book's authors")
	ErrGetBookAuthorIDs = errors.New("Error retrieving Book's authors")
	ErrGetAuthorsByID   = errors.New("Error retrieving authors")
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

	SaveAuthors(authors []string) error
	GetAuthorIDs(authors []string) ([]int64, error)
	SaveBookAuthors(bookID string, authorIDs []int64) error
	GetBookAuthorIDs(bookID string) ([]int64, error)
	GetAuthorsByID(authorIDs []int64) ([]string, error)
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
	newBook = NewBook(util.NewID(), book.Title, book.Publisher, book.YearPublished, book.CallNumber, book.CoverPicture, book.ISBN, book.Collation, book.Edition, book.Description, book.LOCClassification, book.Subject, book.Author, book.Quantity, time.Now())

	newBook, err := s.bookRepository.Save(newBook)
	if err != nil {
		return nil, ErrCreateBook
	}

	// Retrieve the subjectIDs of the particular Book that want to be created.
	subjectIDs, err := s.GetSubjectIDs(book.Subject)
	if err != nil {
		return nil, ErrGetSubjectIDs
	}

	// Save the relation between this BookID with all of the subjectIDs.
	err = s.SaveBookSubjects(newBook.ID, subjectIDs)
	if err != nil {
		return nil, ErrSaveBookSubjects
	}

	// Save Authors of the Books.
	err = s.SaveAuthors(book.Author)
	if err != nil {
		return nil, ErrSaveAuthors
	}

	// Retrieve the authorIds of the particular Book that want to be created.
	authorIDs, err := s.GetAuthorIDs(book.Author)
	if err != nil {
		return nil, ErrGetAuthorIDs
	}

	// Save the relation between this BookID with all of the authorIDs.
	err = s.SaveBookAuthors(newBook.ID, authorIDs)
	if err != nil {
		return nil, ErrSaveBookAuthors
	}

	return newBook, nil
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
	book, err := s.bookRepository.Update(book)
	if err != nil {
		return nil, ErrUpdateBook
	}

	return book, nil
}

func (s *service) Delete(bookID string) error {
	err := s.bookRepository.Delete(bookID)
	if err != nil {
		return ErrDeleteBook
	}

	return nil
}

func (s *service) GetSubjectIDs(subjects []string) ([]int64, error) {
	subjectIDs, err := s.bookRepository.GetSubjectIDs(subjects)
	if err != nil {
		return nil, ErrGetSubjectIDs
	}
	return subjectIDs, nil
}

func (s *service) SaveBookSubjects(bookID string, subjectIDs []int64) error {
	err := s.bookRepository.SaveBookSubjects(bookID, subjectIDs)
	if err != nil {
		return ErrSaveBookSubjects
	}

	return nil
}

func (s *service) GetBookSubjectIDs(bookID string) ([]int64, error) {
	bookSubjectIDs, err := s.bookRepository.GetBookSubjectIDs(bookID)
	if err != nil {
		return nil, ErrGetBookSubjectIDs
	}

	return bookSubjectIDs, nil
}

func (s *service) GetSubjectsByID(subjectIDs []int64) ([]string, error) {
	subjects, err := s.bookRepository.GetSubjectsByID(subjectIDs)
	if err != nil {
		return nil, ErrGetSubjectsByID
	}

	return subjects, nil
}

func (s *service) SaveAuthors(authors []string) error {
	err := s.bookRepository.SaveAuthors(authors)
	if err != nil {
		return ErrSaveAuthors
	}

	return nil
}

func (s *service) GetAuthorIDs(authors []string) ([]int64, error) {
	authorIDs, err := s.bookRepository.GetAuthorIDs(authors)
	if err != nil {
		return nil, ErrGetAuthorIDs
	}
	return authorIDs, nil
}

func (s *service) SaveBookAuthors(bookID string, authorIDs []int64) error {
	err := s.bookRepository.SaveBookAuthors(bookID, authorIDs)
	if err != nil {
		return ErrSaveBookAuthors
	}

	return nil
}

func (s *service) GetBookAuthorIDs(bookID string) ([]int64, error) {
	bookAuthorIDs, err := s.bookRepository.GetBookAuthorIDs(bookID)
	if err != nil {
		return nil, ErrGetBookAuthorIDs
	}

	return bookAuthorIDs, nil
}

func (s *service) GetAuthorsByID(authorIDs []int64) ([]string, error) {
	authors, err := s.bookRepository.GetAuthorsByID(authorIDs)
	if err != nil {
		return nil, ErrGetAuthorsByID
	}

	return authors, nil
}
