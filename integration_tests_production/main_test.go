package integrationtestsproduction

import (
	"os"
	"testing"

	"github.com/joho/godotenv"

	"github.com/joshuabezaleel/library-server/persistence"
	"github.com/joshuabezaleel/library-server/pkg/auth"
	"github.com/joshuabezaleel/library-server/pkg/borrowing"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
	"github.com/joshuabezaleel/library-server/server"
)

var (
	repository         *persistence.Repository
	srv                *server.Server
	userStudent        *user.User
	userLibrarian      *user.User
	userStudentToken   string
	userLibrarianToken string
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../build/.env")
	if err != nil {
		panic(err)
	}
	repository = persistence.NewRepository("PRODUCTION")
	defer repository.DB.Close()

	repository.EnsureTableExists()

	// Setting up domain services.
	userService := user.NewUserService(repository.UserRepository)
	authService := auth.NewAuthService(repository.AuthRepository, userService)
	bookService := book.NewBookService(repository.BookRepository)
	bookCopyService := bookcopy.NewBookCopyService(repository.BookCopyRepository, bookService)
	borrowService := borrowing.NewBorrowingService(repository.BorrowRepository, userService, bookCopyService)

	srv = server.NewServer(authService, bookService, bookCopyService, userService, borrowService)

	go srv.Run()

	// Create users with different roles and login.
	userStudent = createUserStudent("globalUserStudent")
	userLibrarian = createUserLibrarian("globalUserLibrarian")

	userStudentToken = "Bearer " + login(userStudent)
	userLibrarianToken = "Bearer " + login(userLibrarian)

	code := m.Run()

	repository.CleanUp()

	os.Exit(code)
}
