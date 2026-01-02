package routers

import (
	"book_order_app/controllers"
	"book_order_app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(rg *gin.RouterGroup) {
	bookController := controllers.InitializeBookController()
	books := rg.Group("/books")
	{
		books.GET("", bookController.GetBooks)
		books.POST("", middleware.AuthMiddleware(), middleware.RequireRole("admin"), bookController.AddBook)
	}
}
