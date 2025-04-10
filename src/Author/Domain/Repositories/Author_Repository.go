package repositories

// Importación de los paquetes necesarios
import (
	entities "ApiBook/src/Author/Domain/Entities"    // Paquete que contiene las entidades relacionadas con el autor.
	
)

// AuthorRepository es una interfaz que define las operaciones disponibles para manipular los autores en el repositorio.
type AuthorRepository interface {
	// CreateAuthor guarda un nuevo autor en el repositorio.
	// Recibe un puntero a un objeto de tipo Author y devuelve un error en caso de que ocurra algún problema durante la creación.
	CreateAuthor(author *entities.Author) error
	
	// GetAuthorByID busca un autor en el repositorio utilizando su ID.
	// Recibe un ID de tipo int16 y devuelve un puntero al autor correspondiente y un error si no se encuentra.
	GetAuthorByID(id int16) (*entities.Author, error)
	
	// UpdateAuthor actualiza los datos de un autor en el repositorio.
	// Recibe un puntero a un objeto Author con los nuevos datos y devuelve un error en caso de que ocurra algún problema.
	UpdateAuthor(author *entities.Author) error
	
	// DeleteAuthor elimina un autor del repositorio según su ID.
	// Recibe un ID de tipo int16 y devuelve un error si no se puede eliminar el autor.
	DeleteAuthor(id int16) error
	
	// GetAllAuthor obtiene todos los autores del repositorio.
	// Devuelve un slice de autores y un error si ocurre algún problema.
	GetAllAuthor() ([]entities.Author, error)
	
	
}
