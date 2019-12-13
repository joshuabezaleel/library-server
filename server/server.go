package server

import (
	"github.com/gorilla/mux"
	"github.com/joshuabezaleel/library-server/pkg/core/book"
	"github.com/joshuabezaleel/library-server/pkg/core/bookcopy"
	"github.com/joshuabezaleel/library-server/pkg/core/user"
)

// Server holds dependencies for the HTTP server.
type Server struct {
	bookService     book.Service
	bookCopyService bookcopy.Service
	userService     user.Service

	Router *mux.Router
}

// NewServer returns a new HTTP server
// with all of the necessary dependencies.
func NewServer(bookService book.Service, bookCopyService bookcopy.Service, userService user.Service) *Server {
	server := &Server{
		bookService:     bookService,
		bookCopyService: bookCopyService,
		userService:     userService,
	}

	bookHandler := bookHandler{bookService}
	bookCopyHandler := bookCopyHandler{bookCopyService}
	userHandler := userHandler{userService}

	router := mux.NewRouter()
	bookHandler.registerRouter(router)
	bookCopyHandler.registerRouter(router)
	userHandler.registerRouter(router)

	server.Router = router

	return server
}
