package application

import (
	entities "ApiBook/src/Author/Domain/Entities"
	repositories "ApiBook/src/Author/Domain/Repositories"
	
)

// Servicio para manejar operaciones relacionadas con los autores
type AuthorService struct {
	repository repositories.AuthorRepository
}

// Constructor para AuthorService
func NewAuthorService(repo repositories.AuthorRepository) *AuthorService {
	return &AuthorService{
		repository: repo,
	}
}

// Crear un nuevo autor
func (s *AuthorService) CreateAuthor(author *entities.Author) error {
	return s.repository.CreateAuthor(author)
}

// Obtener un autor por su ID
func (s *AuthorService) GetAuthorByID(id int16) (*entities.Author, error) {
	return s.repository.GetAuthorByID(id)
}

// Obtener todos los autores
func (s *AuthorService) GetAllAuthors() ([]entities.Author, error) {
	return s.repository.GetAllAuthor()
}

// Actualizar la información de un autor
func (s *AuthorService) UpdateAuthor(author *entities.Author) error {
	return s.repository.UpdateAuthor(author)
}

// Eliminar un autor por su ID
func (s *AuthorService) DeleteAuthor(id int16) error {
	return s.repository.DeleteAuthor(id)
}