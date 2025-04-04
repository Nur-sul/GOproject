package service

import (
	"rest-project/internal/models"
)

// Интерфейс репозитория книг
type BookRepository interface {
	GetAll() ([]models.Book, error)
	GetById(id int) (*models.Book, error)
	Create(book *models.Book) error
	Update(id int, book *models.BookEdit) error
	Delete(bookID int) error
}

// Структура сервиса книг
type BookService struct {
	repo BookRepository
}

// Конструктор BookService
func NewBookService(bookRepo BookRepository) *BookService {
	return &BookService{repo: bookRepo}
}

// Получение всех книг
func (s *BookService) GetAllBooks() ([]models.Book, error) {
	return s.repo.GetAll()
}

// Получение книги по ID
func (s *BookService) GetBookByID(id int) (*models.Book, error) {
	return s.repo.GetById(id)
}

// Создание новой книги
func (s *BookService) Create(title, author string) (*models.Book, error) {
	book := &models.Book{
		Title:  title,
		Author: author,
	}
	err := s.repo.Create(book)
	return book, err
}

// Обновление данных книги
func (s *BookService) Update(id int, book *models.BookEdit) (*models.Book, error) {
	err := s.repo.Update(id, book)
	if err != nil {
		return nil, err
	}
	return s.GetBookByID(id)
}

// Удаление книги
func (s *BookService) DeleteBook(bookID int) error {
	return s.repo.Delete(bookID)
}
