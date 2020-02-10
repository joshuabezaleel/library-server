package integrationtests

import (
	"os"
	"testing"

	"github.com/joho/godotenv"

	"github.com/joshuabezaleel/library-server/persistence"
	"github.com/joshuabezaleel/library-server/server"
)

var repository *persistence.Repository
var srv *server.Server

const (
	deployment = "TESTING"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../build/.env")
	if err != nil {
		panic(err)
	}
	repository = persistence.NewRepository(deployment)
	defer repository.DB.Close()

	repository.EnsureTableExists()

	// Setting up domain services.
	// userService := user.NewUserService(repository.UserRepository)
	// authService := auth.NewAuthService(repository.AuthRepository, userService)
	// bookService := book.NewBookService(repository.BookRepository)
	// bookCopyService := bookcopy.NewBookCopyService(repository.BookCopyRepository, bookService)
	// borrowService := borrowing.NewBorrowingService(repository.BorrowRepository, userService, bookCopyService)

	// srv = server.NewServer(deployment, authService, bookService, bookCopyService, userService, borrowService)

	// go srv.Run(deployment)

	code := m.Run()

	repository.CleanUp()

	os.Exit(code)
}
