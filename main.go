package main

import (
	"log"

	_ "book_order_app/docs"
	"book_order_app/routers"

	"github.com/gin-gonic/gin"
)

// @title Book Order API
// @version 1.0
// @description API for managing books and orders
// @host localhost:8080
// @BasePath /api/v1
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {

	// Initialize Gin router
	r := gin.Default()
	routers.RegisterRoutes(r)

	// Start server
	log.Println("Server starting on :8080")
	log.Println("Swagger documentation available at http://localhost:8080/swagger/index.html")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
