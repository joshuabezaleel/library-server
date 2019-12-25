package persistence

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Importing postgre SQL driver

	"github.com/joshuabezaleel/library-server/pkg/auth"
	"github.com/joshuabezaleel/library-server/pkg/borrowing"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
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
func NewRepository(env string) *Repository {
	// err := godotenv.Load("build/.env")
	// if err != nil {
	// 	panic(err)
	// }

	var dbName string
	if env == "testing" {
		dbName = os.Getenv("DB_NAME")
	} else if env == "production" {
		dbName = os.Getenv("DB_NAME")
	}

	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	connectionString := fmt.Sprintf("host = %s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), dbPort, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), dbName)

	DB, err := sqlx.Open(os.Getenv("DB_DRIVER"), connectionString)
	if err != nil {
		panic(err)
	}

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
