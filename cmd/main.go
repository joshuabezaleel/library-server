package main

import (
	_ "github.com/lib/pq"

	"github.com/joshuabezaleel/library-server/persistence"
	"github.com/joshuabezaleel/library-server/pkg/auth"
	"github.com/joshuabezaleel/library-server/pkg/borrowing"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
	"github.com/joshuabezaleel/library-server/server"
)

const (
	deployment string = "PRODUCTION"
)

func main() {
	repository := persistence.NewRepository(deployment)

	// Setting up domain services.
	userService := user.NewUserService(repository.UserRepository)
	authService := auth.NewAuthService(repository.AuthRepository, userService)
	bookService := book.NewBookService(repository.BookRepository)
	bookCopyService := bookcopy.NewBookCopyService(repository.BookCopyRepository, bookService)
	borrowService := borrowing.NewBorrowingService(repository.BorrowRepository, userService, bookCopyService)

	srv := server.NewServer(deployment, authService, bookService, bookCopyService, userService, borrowService)
	srv.Run(deployment)

	repository.DB.Close()
}
