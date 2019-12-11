package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/joshuabezaleel/library-server/persistence"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
	"github.com/joshuabezaleel/library-server/server"
)

const (
	serverPort         = ":8082"
	connectionHost     = "localhost"
	connectionPort     = 8081
	connectionUsername = "postgres"
	connectionPassword = "postgres"
	dbName             = "library-server"
)

func main() {
	connectionString := fmt.Sprintf("host = %s port=%d user=%s password=%s dbname=%s sslmode=disable", connectionHost, connectionPort, connectionUsername, connectionUsername, dbName)
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Setting up domain repositories.
	bookRepository := persistence.NewBookRepository(db)
	userRepository := persistence.NewUserRepository(db)

	// Setting up domain services.
	bookService := book.NewBookService(bookRepository)
	userService := user.NewUserService(userRepository)

	srv := server.NewServer(bookService, userService)
	fmt.Println("Server is running...")

	err = http.ListenAndServe(serverPort, srv.Router)
	if err != nil {
		log.Println(err)
	}
}
