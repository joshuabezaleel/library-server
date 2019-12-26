package persistence

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

const (
	deployment = "TESTING"
)

var repository *Repository

func TestMain(m *testing.M) {
	err := godotenv.Load("../build/.env")
	if err != nil {
		panic(err)
	}
	repository = NewRepository(deployment)
	defer repository.DB.Close()

	// repository.EnsureDatabaseExists()
	repository.EnsureTableExists()

	code := m.Run()

	repository.CleanUp()

	os.Exit(code)
}
