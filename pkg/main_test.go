package pkg_test

import (
	"os"
	"testing"

	"github.com/joshuabezaleel/library-server/pkg/core/book"
)

var bookRepository book.MockRepository

func TestMain(m *testing.M) {
	//  := &user.MockRepository{}

	code := m.Run()

	os.Exit(code)
}
