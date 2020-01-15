package bookcopy

import (
	"testing"
	"time"

	"github.com/bouk/monkey"
	"github.com/stretchr/testify/require"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
)

var bookCopyRepository = &MockRepository{}
var bookRepository = &book.MockRepository{}

var bookService = book.NewBookService(bookRepository)
var bookCopyService = service{
	bookCopyRepository: bookCopyRepository,
	bookService:        bookService,
}

func TestCreate(t *testing.T) {
	createdTime := time.Now()
	timePatch := monkey.Patch(time.Now, func() time.Time {
		return createdTime
	})
	defer timePatch.Unpatch()

	initialBookCopy := &BookCopy{
		ID:      util.NewID(),
		BookID:  util.NewID(),
		AddedAt: createdTime,
	}
	book := &book.Book{
		ID:       util.NewID(),
		Quantity: 0,
		AddedAt:  createdTime,
	}

	bookCopyRepository.On("Save", initialBookCopy).Return(initialBookCopy, nil)
	bookRepository.On("Get", initialBookCopy.BookID).Return(book, nil)

	book.Quantity++
	bookRepository.On("Update", book).Return(book, nil)

	newBookCopy, err := bookCopyService.Create(initialBookCopy)

	require.Nil(t, err)
	require.Equal(t, initialBookCopy.ID, newBookCopy.ID)
}

func TestGet(t *testing.T) {
	bookCopy := &BookCopy{
		ID: util.NewID(),
	}
	bookCopyRepository.On("Get", bookCopy.ID).Return(bookCopy, nil)

	newBookCopy, err := bookCopyService.Get(bookCopy.ID)

	require.Nil(t, err)
	require.Equal(t, bookCopy.ID, newBookCopy.ID)
}

func TestUpdate(t *testing.T) {
	bookCopy := &BookCopy{
		ID:        util.NewID(),
		Condition: "Repaired",
	}

	expectedBookCopy := &BookCopy{
		ID:        bookCopy.ID,
		Condition: "Available",
	}

	bookCopyRepository.On("Update", bookCopy).Return(expectedBookCopy, nil)

	updatedBookCopy, err := bookCopyService.Update(bookCopy)

	require.Nil(t, err)
	require.Equal(t, bookCopy.ID, updatedBookCopy.ID)
	require.Equal(t, expectedBookCopy.Condition, updatedBookCopy.Condition)
}

func TestDelete(t *testing.T) {
	bookCopy := &BookCopy{
		ID: util.NewID(),
	}

	bookCopyRepository.On("Delete", bookCopy.ID).Return(nil)

	err := bookCopyService.Delete(bookCopy.ID)

	require.Nil(t, err)
}
