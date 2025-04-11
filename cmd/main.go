package main

import (
	"rest-project/internal/db"
	"rest-project/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r, db.DB)
	r.Run(":8080")
}
