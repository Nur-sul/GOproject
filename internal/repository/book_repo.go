package repository

import (
	"rest-project/internal/models"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	db *gorm.DB
}

// NewBookRepository - Constructor
func NewBookRepository(db *gorm.DB) *BookRepositoryImpl {
	return &BookRepositoryImpl{db: db}
}

// GetAll - Retrieve all books
func (b BookRepositoryImpl) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := b.db.Find(&books).Error
	return books, err
}

// GetById - Retrieve a book by ID
func (b BookRepositoryImpl) GetById(id int) (*models.Book, error) {
	var book models.Book
	err := b.db.First(&book, id).Error
	return &book, err
}

// Create - Add a new book
func (b BookRepositoryImpl) Create(book *models.Book) error {
	return b.db.Create(book).Error
}

// Update - Modify an existing book
func (b BookRepositoryImpl) Update(id int, book *models.BookEdit) error {
	return b.db.Model(&models.Book{}).Where("id = ?", id).Omit("id").Updates(book).Error
}

// Delete - Remove a book by ID
func (b BookRepositoryImpl) Delete(bookID int) error {
	return b.db.Delete(&models.Book{}, bookID).Error
}
