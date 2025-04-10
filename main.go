package main

import (
    "time"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    AuthorController "ApiBook/src/Author/Infraestructure/Controller"
    AuthorService "ApiBook/src/Author/Application"
    AuthorDb "ApiBook/src/Author/Infraestructure/Database"
    AuthorRoutes "ApiBook/src/Author/Infraestructure/Routers"
    BookController "ApiBook/src/Book/Infraestructure/Controller"
    BookService "ApiBook/src/Book/Application"
    BookDb "ApiBook/src/Book/Infraestructure/Database"
    BookRoutes "ApiBook/src/Book/Infraestructure/Routes"
    "ApiBook/src/core"
)

func main() {
    // ================== CONFIGURACIÓN INICIAL ==================
    db, err := core.ConnectDB()
    if err != nil {
        log.Fatalf("Error al conectar a la base de datos: %v", err)
    }
    defer db.Close()

    // ================== INICIALIZACIÓN DE DEPENDENCIAS ==================
    // Autor
    authorRepo := AuthorDb.NewMySQLAuthorRepository(db)
    authorService := AuthorService.NewAuthorService(authorRepo)
    authorController := AuthorController.NewAuthorController(authorService)

    // Libro
    bookRepo := BookDb.NewsqlBookRepository(db)
    bookService := BookService.NewBookService(bookRepo)
    bookController := BookController.NewBookController(bookService)

    // ================== CONFIGURACIÓN DEL SERVIDOR ==================
    router := gin.Default()
    
    // 1. Deshabilitar redirecciones automáticas
    router.RedirectTrailingSlash = false
    
    // 2. Configuración de CORS (Detallada)
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
        ExposeHeaders:    []string{"Content-Length", "Content-Type"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

    // 3. Registrar rutas
    AuthorRoutes.RegisterAuthorRoutes(router, authorController)
    BookRoutes.RegisterBookRoutes(router, bookController)

    // ================== INICIAR SERVIDOR ==================
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Error al iniciar el servidor: %v", err)
    }
}