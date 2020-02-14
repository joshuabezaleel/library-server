package book

import (
	"testing"
	"time"

	"github.com/bouk/monkey"
	"github.com/stretchr/testify/require"

	util "github.com/joshuabezaleel/library-server/pkg"
)

var bookRepository = &MockRepository{}
var bookService = service{bookRepository: bookRepository}

func TestCreate(t *testing.T) {
	createdTime := time.Now()
	timePatch := monkey.Patch(time.Now, func() time.Time {
		return createdTime
	})
	defer timePatch.Unpatch()

	ID := util.NewID()
	IDPatch := monkey.Patch(util.NewID, func() string {
		return ID
	})
	defer IDPatch.Unpatch()

	subjects := []string{"Mathematics", "Physics"}
	subjectIDs := []int64{1, 2}
	book := &Book{
		ID:      ID,
		Title:   "book title",
		AddedAt: createdTime,
		Subject: subjects,
	}

	bookRepository.On("Save", book).Return(book, nil)
	bookRepository.On("GetSubjectIDs", book.Subject).Return(subjectIDs, nil)
	bookRepository.On("SaveBookSubjects", book.ID, subjectIDs).Return(nil)
	bookRepository.On("Get", book.ID).Return(book, nil)
	bookRepository.On("GetBookSubjectIDs", book.ID).Return(subjectIDs, nil)
	bookRepository.On("GetSubjectsByID", subjectIDs).Return(subjects, nil)

	newBook, err := bookService.Create(book)

	require.Nil(t, err)
	require.Equal(t, book.ID, newBook.ID)
	require.Equal(t, book.Title, newBook.Title)
}

func TestGet(t *testing.T) {
	book := &Book{
		ID: util.NewID(),
	}
	bookRepository.On("Get", book.ID).Return(book, nil)

	newBook, err := bookService.Get(book.ID)

	require.Nil(t, err)
	require.Equal(t, book.ID, newBook.ID)
}

func TestUpdate(t *testing.T) {
	book := &Book{
		ID:    util.NewID(),
		Title: "title",
	}

	expectedBook := &Book{
		ID:    book.ID,
		Title: "edited title",
	}

	bookRepository.On("Update", book).Return(expectedBook, nil)

	updatedBook, err := bookService.Update(book)

	require.Nil(t, err)
	require.Equal(t, book.ID, updatedBook.ID)
	require.Equal(t, expectedBook.Title, updatedBook.Title)
}

func TestDelete(t *testing.T) {
	book := &Book{
		ID: util.NewID(),
	}

	bookRepository.On("Delete", book.ID).Return(nil)

	err := bookService.Delete(book.ID)

	require.Nil(t, err)
}
