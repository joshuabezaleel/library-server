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

	book := &Book{
		ID:      util.NewID(),
		AddedAt: createdTime,
	}
	bookRepository.On("Save", book).Return(book, nil)

	newBook, err := bookService.Create(book)

	require.Nil(t, err)
	require.Equal(t, book.ID, newBook.ID)
}

func TestGet(t *testing.T) {
	book := &Book{
		ID: util.NewID(),
	}
	bookRepository.On("Get", book.ID).Return(book, nil)

	newBook, err := bookService.Get(book.ID)

	require.Nil(t, err)
	require.Equal(t, newBook.ID, book.ID)
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
	require.Equal(t, updatedBook.ID, book.ID)
	require.Equal(t, updatedBook.Title, expectedBook.Title)
}

func TestDelete(t *testing.T) {
	book := &Book{
		ID:    util.NewID(),
		Title: "title",
	}

	bookRepository.On("Delete", book.ID).Return(nil)

	err := bookService.Delete(book.ID)
	require.Nil(t, err)
}
