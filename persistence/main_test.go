package persistence

import (
	"database/sql"
	"os"
	"testing"

	"github.com/joshuabezaleel/library-server/pkg/auth"
	"github.com/joshuabezaleel/library-server/pkg/borrowing"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
	"github.com/joshuabezaleel/library-server/pkg/core/user"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

// const (
// 	deployment = "TESTING"
// )

var (
	DB                        *sqlx.DB
	Mock                      sqlmock.Sqlmock
	AuthTestingRepository     auth.Repository
	BookTestingRepository     book.Repository
	BookCopyTestingRepository bookcopy.Repository
	BorrowTestingRepository   borrowing.Repository
	UserTestingRepository     user.Repository
)

// var repository *Repository

func TestMain(m *testing.M) {
	var mockDB *sql.DB
	var err error
	mockDB, Mock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()
	DB = sqlx.NewDb(mockDB, "sqlmock")

	// Repositories initialization with mock db
	AuthTestingRepository = NewAuthRepository(DB)
	BookTestingRepository = NewBookRepository(DB)
	BookCopyTestingRepository = NewBookCopyRepository(DB)
	BorrowTestingRepository = NewBorrowRepository(DB)
	UserTestingRepository = NewUserRepository(DB)

	code := m.Run()

	err = Mock.ExpectationsWereMet()
	if err != nil {
		panic(err)
	}

	os.Exit(code)
	// err := godotenv.Load("../build/.env")
	// if err != nil {
	// 	panic(err)
	// }
	// repository = NewRepository(deployment)
	// defer repository.DB.Close()

	// // repository.EnsureDatabaseExists()
	// repository.EnsureTableExists()

	// code := m.Run()

	// repository.CleanUp()

	// os.Exit(code)
}
