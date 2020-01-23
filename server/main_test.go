package server

import (
	"os"
	"testing"

	"github.com/joshuabezaleel/library-server/pkg/auth"
	"github.com/joshuabezaleel/library-server/pkg/borrowing"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
)

var (
	authTestingHandler     authHandler
	bookTestingHandler     bookHandler
	borrowTestingHandler   borrowingHandler
	bookCopyTestingHandler bookCopyHandler
	userTestingHandler     userHandler

	authService     *auth.MockService
	borrowService   *borrowing.MockService
	bookService     *book.MockService
	bookCopyService *bookcopy.MockService
	userService     *user.MockService
)

func TestMain(m *testing.M) {
	// Initiating mock services.
	authService = &auth.MockService{}
	borrowService = &borrowing.MockService{}
	bookService = &book.MockService{}
	bookCopyService = &bookcopy.MockService{}
	userService = &user.MockService{}

	// Initiating handlers with dependency to mock service.
	authTestingHandler = authHandler{authService}
	borrowTestingHandler = borrowingHandler{borrowService, authService}
	bookTestingHandler = bookHandler{bookService, authService}
	bookCopyTestingHandler = bookCopyHandler{bookCopyService, authService}
	userTestingHandler = userHandler{userService, authService}

	code := m.Run()

	os.Exit(code)
}
