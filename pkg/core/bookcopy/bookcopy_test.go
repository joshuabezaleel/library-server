package bookcopy

import (
	"testing"

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
	subjects := []string{"Mathematics", "Physics"}
	subjectIDs := []int64{1, 2}
	authors := []string{"author1", "author2"}
	authorIDs := []int64{1, 2}

	createdTime, createdTimePatch := util.CreatedTimePatch()
	defer createdTimePatch.Unpatch()

	book := &book.Book{
		ID:       util.NewID(),
		Subject:  subjects,
		Author:   authors,
		Quantity: 0,
		AddedAt:  createdTime,
	}

	ID, IDPatch := util.NewIDPatch()
	defer IDPatch.Unpatch()

	bookCopy := &BookCopy{
		ID:        ID,
		Condition: "Available",
		BookID:    book.ID,
		AddedAt:   createdTime,
	}

	errorBookCopy := &BookCopy{
		ID:        ID,
		Condition: "Repaired",
		BookID:    book.ID,
		AddedAt:   createdTime,
	}

	tt := []struct {
		name             string
		bookCopy         *BookCopy
		returnedBookCopy *BookCopy
		err              error
	}{
		{
			name:             "success creating a Book Copy",
			bookCopy:         bookCopy,
			returnedBookCopy: bookCopy,
			err:              nil,
		},
		{
			name:             "failed creating a Book Copy",
			bookCopy:         errorBookCopy,
			returnedBookCopy: nil,
			err:              ErrCreateBookCopy,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bookCopyRepository.On("Save", tc.bookCopy).Return(tc.returnedBookCopy, tc.err)
			bookRepository.On("Get", book.ID).Return(book, nil)
			bookRepository.On("GetBookSubjectIDs", book.ID).Return(subjectIDs, nil)
			bookRepository.On("GetSubjectsByID", subjectIDs).Return(subjects, nil)
			bookRepository.On("GetBookAuthorIDs", book.ID).Return(authorIDs, nil)
			bookRepository.On("GetAuthorsByID", authorIDs).Return(authors, nil)

			book.Quantity++
			bookRepository.On("Update", book).Return(book, nil)

			returnedBookCopy, err := bookCopyService.Create(tc.bookCopy)

			require.Equal(t, tc.err, err)

			if tc.err == nil {
				require.Equal(t, tc.bookCopy.ID, returnedBookCopy.ID)
				require.Equal(t, tc.bookCopy.Condition, returnedBookCopy.Condition)
			}
		})
	}
}

func TestGet(t *testing.T) {
	bookCopy := &BookCopy{
		ID: util.NewID(),
	}

	tt := []struct {
		name             string
		ID               string
		returnedBookCopy *BookCopy
		err              error
	}{
		{
			name:             "success retrieving a Book Copy",
			ID:               bookCopy.ID,
			returnedBookCopy: bookCopy,
			err:              nil,
		},
		{
			name:             "failed retrieving a Book Copy",
			ID:               util.NewID(),
			returnedBookCopy: nil,
			err:              ErrGetBookCopy,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bookCopyRepository.On("Get", tc.ID).Return(tc.returnedBookCopy, tc.err)

			returnedBookCopy, err := bookCopyService.Get(tc.ID)

			require.Equal(t, tc.err, err)

			if tc.err == nil {
				require.Equal(t, bookCopy.ID, returnedBookCopy.ID)
			}
		})
	}
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

	errorBookCopy := &BookCopy{
		ID: util.NewID(),
	}

	tt := []struct {
		name             string
		bookCopy         *BookCopy
		returnedBookCopy *BookCopy
		err              error
	}{
		{
			name:             "success updating a Book Copy",
			bookCopy:         bookCopy,
			returnedBookCopy: expectedBookCopy,
			err:              nil,
		},
		{
			name:             "failed updating a Book Copy",
			bookCopy:         errorBookCopy,
			returnedBookCopy: nil,
			err:              ErrUpdateBookCopy,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bookCopyRepository.On("Update", tc.bookCopy).Return(tc.returnedBookCopy, tc.err)

			updatedBookCopy, err := bookCopyService.Update(tc.bookCopy)

			require.Equal(t, tc.err, err)

			if tc.err == nil {
				require.Equal(t, expectedBookCopy.ID, updatedBookCopy.ID)
				require.Equal(t, expectedBookCopy.Condition, updatedBookCopy.Condition)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	bookCopy := &BookCopy{
		ID: util.NewID(),
	}

	tt := []struct {
		name string
		ID   string
		err  error
	}{
		{
			name: "success deleting a Book Copy",
			ID:   bookCopy.ID,
			err:  nil,
		},
		{
			name: "failed deleting a Book Copy",
			ID:   util.NewID(),
			err:  ErrDeleteBookCopy,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bookCopyRepository.On("Delete", tc.ID).Return(tc.err)

			err := bookCopyService.Delete(tc.ID)

			require.Equal(t, tc.err, err)
		})
	}
}
