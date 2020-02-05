package persistence

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Importing postgre SQL driver

	"github.com/joshuabezaleel/library-server/pkg/auth"
	"github.com/joshuabezaleel/library-server/pkg/borrowing"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
)

var tableCreationQueries = []string{bookTable, bookCopyTable, borrowTable, userTable}

const (
	bookTable = `CREATE TABLE IF NOT EXISTS books (
			id VARCHAR(27),
			title VARCHAR,
			publisher VARCHAR,
			year_published INT,
			call_number VARCHAR UNIQUE,
			cover_picture VARCHAR,
			isbn VARCHAR UNIQUE,
			book_collation TEXT,
			edition VARCHAR,
			description TEXT,
			loc_classification VARCHAR(2),
			subject VARCHAR[],
			author VARCHAR[],
			quantity INT,
			added_at TIMESTAMP WITHOUT TIME ZONE,
			CONSTRAINT books_pkey PRIMARY KEY (id)
			)`
	bookCopyTable = `CREATE TABLE IF NOT EXISTS bookcopies (
			id VARCHAR(27),
			barcode VARCHAR UNIQUE,
			book_id VARCHAR(27),
			condition VARCHAR,
			added_at TIMESTAMP WITHOUT TIME ZONE,
			CONSTRAINT bookcopies_pkey PRIMARY KEY (id)
			)`
	borrowTable = `CREATE TABLE IF NOT EXISTS borrows (
			id VARCHAR(27),
			user_id VARCHAR(27),
			bookcopy_id VARCHAR(27) UNIQUE, 
			fine INT,
			borrowed_at TIMESTAMP WITHOUT TIME ZONE,
			due_date TIMESTAMP WITHOUT TIME ZONE,
			returned_at TIMESTAMP WITHOUT TIME ZONE,
			CONSTRAINT borrows_pkey PRIMARY KEY (id)
			)`
	userTable = `CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(27),
			student_id VARCHAR(8) UNIQUE,
			role VARCHAR,
			username VARCHAR UNIQUE,
			email VARCHAR UNIQUE,
			password VARCHAR,
			total_fine INT,
			registered_at TIMESTAMP WITHOUT TIME ZONE,
			CONSTRAINT users_pkey PRIMARY KEY (id)
			)`
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
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	basepath = strings.TrimSuffix(basepath, "/persistence")

	err := godotenv.Load(basepath + "/build/.env")
	if err != nil {
		panic(err)
	}

	var dbName string
	if env == "PRODUCTION" {
		dbName = os.Getenv("DB_NAME")
	} else if env == "TESTING" {
		dbName = os.Getenv("DB_TESTING_NAME")
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

// // EnsureDatabaseExists runs the query for creating the database
// // in case it does not exists.
// func (repo *Repository) EnsureDatabaseExists() {
// 	if _, err := repo.DB.Exec(`CREATE DATABASE IF NOT EXISTS library-server-test`); err != nil {
// 		log.Println(err)
// 	}
// }

// EnsureTableExists runs the query for creating tables
// if it do not exist.
func (repo *Repository) EnsureTableExists() {
	for _, query := range tableCreationQueries {
		if _, err := repo.DB.Exec(query); err != nil {
			panic(err)
		}
	}
}

// CleanUp make sure that all of the data from all of the
// tables are deleted.
func (repo *Repository) CleanUp() {
	repo.DB.Exec("TRUNCATE books")
	repo.DB.Exec("TRUNCATE bookcopies")
	repo.DB.Exec("TRUNCATE borrows")
	repo.DB.Exec("TRUNCATE users")
}
