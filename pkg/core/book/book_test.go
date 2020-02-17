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
	authors := []string{"author1", "author2"}
	authorIDs := []int64{1, 2}

	book := &Book{
		ID:      ID,
		Title:   "book title",
		Subject: subjects,
		Author:  authors,
		AddedAt: createdTime,
	}

	errorBook := &Book{
		ID:      ID,
		Title:   "error book",
		Subject: subjects,
		Author:  authors,
		AddedAt: createdTime,
	}

	tt := []struct {
		name          string
		book          *Book
		returnPayload *Book
		err           error
	}{
		{
			name:          "success creating a Book",
			book:          book,
			returnPayload: book,
			err:           nil,
		},
		{
			name:          "failed creating a Book",
			book:          errorBook,
			returnPayload: nil,
			err:           ErrCreateBook,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bookRepository.On("Save", tc.book).Return(tc.returnPayload, tc.err)
			bookRepository.On("GetSubjectIDs", book.Subject).Return(subjectIDs, nil)
			bookRepository.On("SaveBookSubjects", book.ID, subjectIDs).Return(nil)
			bookRepository.On("SaveAuthors", book.Author).Return(nil)
			bookRepository.On("GetAuthorIDs", book.Author).Return(authorIDs, nil)
			bookRepository.On("SaveBookAuthors", book.ID, authorIDs).Return(nil)

			newBook, err := bookService.Create(book)

			require.Nil(t, err)

			if err == nil {
				require.Equal(t, book.ID, newBook.ID)
				require.Equal(t, book.Title, newBook.Title)
			}
		})
	}
}

func TestGet(t *testing.T) {
	ID := util.NewID()
	IDPatch := monkey.Patch(util.NewID, func() string {
		return ID
	})
	defer IDPatch.Unpatch()

	subjects := []string{"Mathematics", "Physics"}
	subjectIDs := []int64{1, 2}
	authors := []string{"author1", "author2"}
	authorIDs := []int64{1, 2}

	book := &Book{
		ID:      ID,
		Title:   "book title",
		Subject: subjects,
		Author:  authors,
	}

	bookRepository.On("Get", book.ID).Return(book, nil)
	bookRepository.On("GetBookSubjectIDs", book.ID).Return(subjectIDs, nil)
	bookRepository.On("GetSubjectsByID", subjectIDs).Return(subjects, nil)
	bookRepository.On("GetBookAuthorIDs", book.ID).Return(authorIDs, nil)
	bookRepository.On("GetAuthorsByID", authorIDs).Return(authors, nil)

	newBook, err := bookService.Get(book.ID)

	require.Nil(t, err)
	require.Equal(t, book.ID, newBook.ID)
	require.Equal(t, book.Title, newBook.Title)
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
