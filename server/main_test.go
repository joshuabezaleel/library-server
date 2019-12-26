package server

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
)

// var repository *Repository
var srv *Server

func TestMain(m *testing.M) {
	err := godotenv.Load("../build/.env")
	if err != nil {
		panic(err)
	}
	repository := persistence.NewRepository("testing")
	defer repository.DB.Close()

	// Setting up domain services.
	userService := user.NewUserService(repository.UserRepository)
	authService := auth.NewAuthService(repository.AuthRepository, userService)
	bookService := book.NewBookService(repository.BookRepository)
	bookCopyService := bookcopy.NewBookCopyService(repository.BookCopyRepository, bookService)
	borrowService := borrowing.NewBorrowingService(repository.BorrowRepository, userService, bookCopyService)

	srv = NewServer(authService, bookService, bookCopyService, userService, borrowService)
	// srv.Router.SkipClean(true)
	go srv.Run(":" + os.Getenv("SERVER_PORT"))

	code := m.Run()

	os.Exit(code)
}
