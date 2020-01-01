package book

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	mockBookRepository := &MockRepository{}

	book := &Book{
		ID:    "123",
		Title: "Kamasutra",
	}
	mockBookRepository.On("Get", "123").Return(book, nil)

	bookService := service{bookRepository: mockBookRepository}
	newBook, err := bookService.Get("123")
	mockBookRepository.AssertExpectations(t)
	require.Nil(t, err)
	require.Equal(t, newBook.Title, "Kamasutra")
}
