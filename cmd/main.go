package main

import (
	"log"
	"rest-project/internal/models"
	"rest-project/internal/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:myPassword@localhost:5444/bookdatabase?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatal("Error on migrating to the DB", err)
	}

	r := gin.Default()
	routes.SetupRoutes(r, db)
	r.Run(":8080")
}
