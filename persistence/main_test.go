package persistence

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var repository *Repository

func TestMain(m *testing.M) {
	err := godotenv.Load("../build/.env")
	if err != nil {
		panic(err)
	}
	repository = NewRepository("testing")
	defer repository.DB.Close()

	code := m.Run()

	os.Exit(code)
}
