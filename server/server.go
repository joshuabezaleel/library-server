package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joshuabezaleel/library-server/pkg/auth"
	"github.com/joshuabezaleel/library-server/pkg/borrowing"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
	"github.com/joshuabezaleel/library-server/pkg/core/user"

	"github.com/gorilla/mux"
)

// Server holds dependencies for the HTTP server.
type Server struct {
	authService     auth.Service
	bookService     book.Service
	bookCopyService bookcopy.Service
	userService     user.Service
	borrowService   borrowing.Service

	Router *mux.Router
}

// NewServer returns a new HTTP server
// with all of the necessary dependencies.
func NewServer(authService auth.Service, bookService book.Service, bookCopyService bookcopy.Service, userService user.Service, borrowService borrowing.Service) *Server {
	server := &Server{
		authService:     authService,
		bookService:     bookService,
		bookCopyService: bookCopyService,
		userService:     userService,
		borrowService:   borrowService,
	}

	authHandler := authHandler{authService}
	bookHandler := bookHandler{bookService, authService}
	bookCopyHandler := bookCopyHandler{bookCopyService, authService}
	userHandler := userHandler{userService, authService}
	borrowHandler := borrowingHandler{borrowService, authService}

	router := mux.NewRouter()

	authHandler.registerRouter(router)
	bookHandler.registerRouter(router)
	bookCopyHandler.registerRouter(router)
	userHandler.registerRouter(router)
	borrowHandler.registerRouter(router)

	server.Router = router

	return server
}

// Run runs the HTTP server with the port and the router.
func (srv *Server) Run() {
	var port string
	port = ":" + os.Getenv("SERVER_PORT")
	fmt.Println("Server is running ...")

	err := http.ListenAndServe(port, srv.Router)
	if err != nil {
		panic(err)
	}
}

// Stop stops the HTTP server that is currently running.
func (srv *Server) Stop() {

}
