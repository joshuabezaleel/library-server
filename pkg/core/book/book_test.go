package book

import (
	"testing"

	"github.com/stretchr/testify/require"

	util "github.com/joshuabezaleel/library-server/pkg"
)

var bookRepository = &MockRepository{}
var bookService = service{bookRepository: bookRepository}

func TestCreate(t *testing.T) {
	createdTime, createdTimePatch := util.CreatedTimePatch()
	defer createdTimePatch.Unpatch()

	ID, IDPatch := util.NewIDPatch()
	defer IDPatch.Unpatch()

	subjects := []string{"Mathematics", "Physics"}
	subjectIDs := []int64{1, 2}
	authors := []string{"author1", "author2"}
	authorIDs := []int64{1, 2}

	book := &Book{
		ID:      ID,
		Title:   "book",
		Subject: subjects,
		Author:  authors,
		AddedAt: createdTime,
	}

	errorBook := &Book{
		ID:      ID,
		Title:   "errorBook",
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
			bookRepository.On("GetSubjectIDs", tc.book.Subject).Return(subjectIDs, nil)
			bookRepository.On("SaveBookSubjects", tc.book.ID, subjectIDs).Return(nil)
			bookRepository.On("SaveAuthors", tc.book.Author).Return(nil)
			bookRepository.On("GetAuthorIDs", tc.book.Author).Return(authorIDs, nil)
			bookRepository.On("SaveBookAuthors", tc.book.ID, authorIDs).Return(nil)

			newBook, err := bookService.Create(tc.book)

			require.Equal(t, tc.err, err)

			if tc.err == nil {
				require.Equal(t, book.ID, newBook.ID)
				require.Equal(t, book.Title, newBook.Title)
			}
		})
	}
}

func TestGet(t *testing.T) {
	subjects := []string{"Mathematics", "Physics"}
	subjectIDs := []int64{1, 2}
	authors := []string{"author1", "author2"}
	authorIDs := []int64{1, 2}

	book := &Book{
		ID:      util.NewID(),
		Title:   "book",
		Subject: subjects,
		Author:  authors,
	}

	errorBook := &Book{
		ID:      util.NewID(),
		Title:   "errorBook",
		Subject: subjects,
		Author:  authors,
	}

	tt := []struct {
		name          string
		book          *Book
		returnPayload *Book
		err           error
	}{
		{
			name:          "success retrieving a Book",
			book:          book,
			returnPayload: book,
			err:           nil,
		},
		{
			name:          "failed retrieving a Book",
			book:          errorBook,
			returnPayload: nil,
			err:           ErrGetBook,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bookRepository.On("Get", tc.book.ID).Return(tc.returnPayload, tc.err)
			bookRepository.On("GetBookSubjectIDs", tc.book.ID).Return(subjectIDs, nil)
			bookRepository.On("GetSubjectsByID", subjectIDs).Return(subjects, nil)
			bookRepository.On("GetBookAuthorIDs", tc.book.ID).Return(authorIDs, nil)
			bookRepository.On("GetAuthorsByID", authorIDs).Return(authors, nil)

			newBook, err := bookService.Get(tc.book.ID)

			require.Equal(t, tc.err, err)

			if tc.err == nil {
				require.Equal(t, book.ID, newBook.ID)
				require.Equal(t, book.Title, newBook.Title)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	book := &Book{
		ID:    util.NewID(),
		Title: "book",
	}

	expectedBook := &Book{
		ID:    book.ID,
		Title: "edited book",
	}

	errorBook := &Book{
		ID:    util.NewID(),
		Title: "error book",
	}

	tt := []struct {
		name          string
		book          *Book
		returnPayload *Book
		err           error
	}{
		{
			name:          "success updating a Book",
			book:          book,
			returnPayload: expectedBook,
			err:           nil,
		},
		{
			name:          "failed updating a Book",
			book:          errorBook,
			returnPayload: nil,
			err:           ErrUpdateBook,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bookRepository.On("Update", tc.book).Return(tc.returnPayload, tc.err)

			updatedBook, err := bookService.Update(tc.book)

			require.Equal(t, tc.err, err)

			if tc.err == nil {
				require.Equal(t, tc.book.ID, updatedBook.ID)
				require.Equal(t, expectedBook.Title, updatedBook.Title)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	book := &Book{
		ID: util.NewID(),
	}

	errorBook := &Book{
		ID: util.NewID(),
	}

	tt := []struct {
		name string
		ID   string
		err  error
	}{
		{
			name: "success deleting a Book",
			ID:   book.ID,
			err:  nil,
		},
		{
			name: "failed deleting a Book",
			ID:   errorBook.ID,
			err:  ErrDeleteBook,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bookRepository.On("Delete", tc.ID).Return(tc.err)

			err := bookService.Delete(tc.ID)

			require.Equal(t, tc.err, err)
		})
	}
}
