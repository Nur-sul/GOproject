package routes

import (
	"rest-project/internal/delivery"
	"rest-project/internal/repository"
	service "rest-project/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Инициализация репозитория
	bookRepo := repository.NewBookRepository(db)

	// Инициализация сервиса
	bookService := service.NewBookService(bookRepo)

	// Инициализация обработчика
	bookHandler := delivery.NewBookHandler(bookService)

	// Роуты
	books := r.Group("api/v1/books")
	{
		books.GET("/", bookHandler.GetAllBooks)
		books.POST("/", bookHandler.CreateBook)
		books.GET("/:id", bookHandler.GetBook)
		books.PUT("/:id", bookHandler.UpdateBook)
		books.DELETE("/:id", bookHandler.DeleteBook)
	}
}
