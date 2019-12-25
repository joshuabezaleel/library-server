package persistence

import (
	"testing"

	"github.com/stretchr/testify/require"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
)

func TestBookSave(t *testing.T) {
	// Create a new Book and save it.
	book := &book.Book{
		ID: util.NewID(),
	}
	newBook, err := repository.BookRepository.Save(book)

	// Happy path.
	require.Nil(t, err)
	require.Equal(t, book.ID, newBook.ID)
}

func TestBookGet(t *testing.T) {
	// Create a new Book and save it.
	book := &book.Book{
		ID: util.NewID(),
	}
	book1, err := repository.BookRepository.Save(book)
	require.Nil(t, err)

	// Get the Book.
	book2, err := repository.BookRepository.Get(book.ID)
	require.Nil(t, err)
	require.Equal(t, book2.ID, book1.ID)

	// Get invalid Book.
	_, err = repository.BookRepository.Get(util.NewID())
	require.NotNil(t, err)
}

func TestBookUpdate(t *testing.T) {
	// Create a new Book and save it.
	book := &book.Book{
		ID:    util.NewID(),
		Title: "title",
	}
	book1, err := repository.BookRepository.Save(book)
	require.Nil(t, err)

	// Update the Book's title.
	book1.Title = "edited title"
	book2, err := repository.BookRepository.Update(book1)
	require.Nil(t, err)
	require.Equal(t, book1.ID, book2.ID)
	require.Equal(t, book2.Title, "edited title")
}

func TestBookDelete(t *testing.T) {
	// Create a new Book and save it.
	book := &book.Book{
		ID:    util.NewID(),
		Title: "title",
	}
	_, err := repository.BookRepository.Save(book)
	require.Nil(t, err)

	// Delete the Book that was just created.
	err = repository.BookRepository.Delete(book.ID)
	require.Nil(t, err)

	// Unable to retrieve the Book that was just deleted.
	_, err = repository.BookRepository.Get(book.ID)
	require.NotNil(t, err)
}
