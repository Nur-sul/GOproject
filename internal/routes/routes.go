package routes

import (
	"rest-project/internal/auth"
	"rest-project/internal/delivery"
	"rest-project/internal/middleware"
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

	authRoutes := r.Group("api/v1/auth")
	{
		authRoutes.POST("/login", auth.Login)
		authRoutes.POST("/register", auth.Register)
	}

	protected := r.Group("api/v1")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/me", auth.Me)

	}

	books := r.Group("api/v1/books")
	books.Use(middleware.AuthRequired())
	{
		books.GET("/", bookHandler.GetAllBooks)
		books.POST("/", bookHandler.CreateBook)
		books.GET("/:id", bookHandler.GetBook)
		books.PUT("/:id", bookHandler.UpdateBook)
		books.DELETE("/:id", bookHandler.DeleteBook)
	}

}
