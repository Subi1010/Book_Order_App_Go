package routers

import (
	"book_order_app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(rg *gin.RouterGroup) {
	bookController := controllers.InitializeBookController()
	books := rg.Group("/books")
	{
		books.GET("", bookController.GetBooks)
		books.POST("", bookController.AddBook)
	}
}
