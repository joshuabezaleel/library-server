package persistence

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/joshuabezaleel/library-server/pkg/auth"
	"github.com/joshuabezaleel/library-server/pkg/borrowing"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
)

const (
	connectionHost     = "localhost"
	connectionPort     = 8081
	connectionUsername = "postgres"
	connectionPassword = "postgres"
	dbName             = "library-server"
)

// Repository holds dependencies for the current persistence layer.
type Repository struct {
	AuthRepository     auth.Repository
	BookRepository     book.Repository
	BookCopyRepository bookcopy.Repository
	UserRepository     user.Repository
	BorrowRepository   borrowing.Repository

	DB *sqlx.DB
}

// NewRepository returns a new Repository
// with all of the necessary dependencies.
func NewRepository() *Repository {
	connectionString := fmt.Sprintf("host = %s port=%d user=%s password=%s dbname=%s sslmode=disable", connectionHost, connectionPort, connectionUsername, connectionUsername, dbName)

	DB, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	authRepository := NewAuthRepository(DB)
	bookRepository := NewBookRepository(DB)
	bookCopyRepository := NewBookCopyRepository(DB)
	userRepository := NewUserRepository(DB)
	borrowRepository := NewBorrowRepository(DB)

	repository := &Repository{
		AuthRepository:     authRepository,
		BookRepository:     bookRepository,
		BookCopyRepository: bookCopyRepository,
		UserRepository:     userRepository,
		BorrowRepository:   borrowRepository,
		DB:                 DB,
	}

	return repository
}
