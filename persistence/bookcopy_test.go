package persistence

import (
	"testing"

	util "github.com/joshuabezaleel/library-server/pkg"
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
	"github.com/stretchr/testify/require"
)

func TestBookCopySave(t *testing.T) {
	// Create a new BookCopy and save it.
	bookCopy := &bookcopy.BookCopy{
		ID: util.NewID(),
	}
	bookCopy1, err := repository.BookCopyRepository.Save(bookCopy)

	// Happy path.
	require.Nil(t, err)
	require.Equal(t, bookCopy.ID, bookCopy1.ID)
}

func TestBookCopyGet(t *testing.T) {
	// Create a new BookCopy and save it.
	bookCopy := &bookcopy.BookCopy{
		ID: util.NewID(),
	}
	bookCopy1, err := repository.BookCopyRepository.Save(bookCopy)
	require.Nil(t, err)

	// Get the BookCopy.
	bookCopy2, err := repository.BookCopyRepository.Get(bookCopy.ID)
	require.Nil(t, err)
	require.Equal(t, bookCopy1.ID, bookCopy2.ID)

	// Get invalid BookCopy.
	_, err = repository.BookCopyRepository.Get(util.NewID())
	require.NotNil(t, err)
}

func TestBookCopyUpdate(t *testing.T) {
	// Create a new BookCopy and save it.
	bookCopy := &bookcopy.BookCopy{
		ID:        util.NewID(),
		Condition: "good",
	}
	bookCopy1, err := repository.BookCopyRepository.Save(bookCopy)
	require.Nil(t, err)

	// Update BookCopy's Condition
	bookCopy1.Condition = "bad"
	bookCopy2, err := repository.BookCopyRepository.Update(bookCopy1)
	require.Nil(t, err)
	require.Equal(t, bookCopy1.ID, bookCopy2.ID)
	require.Equal(t, bookCopy2.Condition, "bad")
}

func TestBookCopyDelete(t *testing.T) {
	// Create a new BookCopy and save it.
	bookCopy := &bookcopy.BookCopy{
		ID: util.NewID(),
	}
	_, err := repository.BookCopyRepository.Save(bookCopy)
	require.Nil(t, err)

	// Delete the BookCopy that was just created.
	err = repository.BookCopyRepository.Delete(bookCopy.ID)
	require.Nil(t, err)

	// Unable to retrieve the BookCopy that was just deleted.
	_, err = repository.BookCopyRepository.Get(bookCopy.ID)
	require.NotNil(t, err)
}
