package main

import (
	AuthorControlller "ApiBook/src/Author/Infraestructure/Controller"
	AuthorService "ApiBook/src/Author/Application"
	AuthorDb "ApiBook/src/Author/Infraestructure/Database"
	AuthorRoutes "ApiBook/src/Author/Infraestructure/Routers"

	"ApiBook/src/core"

	BookController "ApiBook/src/Book/Infraestructure/Controller"
	BookService "ApiBook/src/Book/Application"
	BookDb "ApiBook/src/Book/Infraestructure/Database"
	BookRoutes "ApiBook/src/Book/Infraestructure/Routes"


	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := core.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect")
	    return
	}

	defer db.Close()

	AuthorRepo :=  AuthorDb.NewsqlAuthorRepository(db)
	AuthorService :=  AuthorService.NewAuthorService(AuthorRepo)
	AuthorControlller := AuthorControlller.NewAuthorController(AuthorService)

	BookRepo := BookDb.NewsqlBookRepository(db)
	BookService := BookService.NewBookService(BookRepo)
	BookController := BookController.NewBookController(BookService)


	router := gin.Default()


	AuthorRoutes.RegisterAuthorRoutes(router , AuthorControlller)
	BookRoutes.RegisterBookRoutes(router, BookController)


	
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
        println(err)
		
	}

	
}