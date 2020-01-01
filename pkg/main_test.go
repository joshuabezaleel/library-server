package pkg

import (
	"os"
	"testing"

	"github.com/joshuabezaleel/library-server/pkg/core/user"
)

func TestMain(m *testing.M) {
	mockUserRepository := &user.MockRepository{}

	code := m.Run()

	os.Exit(code)
}
