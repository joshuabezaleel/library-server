package pkg_test

import (
	"os"
	"testing"

	"github.com/joshuabezaleel/library-server/pkg/auth"
	"github.com/joshuabezaleel/library-server/pkg/borrowing"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
)

var authRepository auth.MockRepository
var borrowRepository borrowing.MockRepository
var bookRepository book.MockRepository
var bookCopyRepository bookcopy.MockRepository
var userRepository user.MockRepository
var A string = "test"

func TestMain(m *testing.M) {
	// mockBookRepository =
	code := m.Run()

	os.Exit(code)
}
