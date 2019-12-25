package main

// func TestMain(m *testing.M) {
// 	repository := persistence.NewRepository("testing")

// 	// Setting up domain services.
// 	userService := user.NewUserService(repository.UserRepository)
// 	authService := auth.NewAuthService(repository.AuthRepository, userService)
// 	bookService := book.NewBookService(repository.BookRepository)
// 	bookCopyService := bookcopy.NewBookCopyService(repository.BookCopyRepository, bookService)
// 	borrowService := borrowing.NewBorrowingService(repository.BorrowRepository, userService, bookCopyService)

// 	srv := server.NewServer(authService, bookService, bookCopyService, userService, borrowService)
// 	log.Printf("hai")

// 	go srv.Run(":8083")
// 	code := m.Run()

// 	repository.DB.Close()

// 	os.Exit(code)
// }
