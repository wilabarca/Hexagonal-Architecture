package controller

import (
	application "ApiBook/src/Book/Application"
	entities "ApiBook/src/Book/Domain/Entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	service *application.BookService
}

func NewBookController(service *application.BookService) *BookController {
	return &BookController{service: service}
}

// Crear un libro
func (c *BookController) CreateBook(ctx *gin.Context) {
	var book entities.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	err := c.service.CreateBook(&book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Devolver el libro creado, incluyendo el ID generado
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Libro Creado",
		"book":    book, // Devuelves el libro con su ID generado
	})
}

// Obtener todos los libros
func (c *BookController) GetAllBooks(ctx *gin.Context) {
	books, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, books)
}

// Obtener un libro por ID
func (c *BookController) GetBookByID(ctx *gin.Context) {
	id := ctx.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de libro inválido"})
		return
	}

	book, err := c.service.GetByID(int64(num))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if book == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Libro no encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, book) // Devolvemos el libro encontrado
}

// Actualizar un libro
func (c *BookController) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var book entities.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	book.ID = int64(bookID)

	err = c.service.UpdateBook(&book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Devolver el libro actualizado
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Libro actualizado",
		"book":    book, // Devuelves el libro actualizado
	})
}

// Eliminar un libro
func (c *BookController) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de libro inválido"})
		return
	}

	err = c.service.DeleteBook(int64(num))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Libro eliminado"})
}
