package routers

import (
	"book_order_app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(rg *gin.RouterGroup) {
	orderController := controllers.InitializeOrderController()
	orders := rg.Group("/orders")
	{
		orders.GET("", orderController.GetOrders)
		orders.POST("", orderController.PlaceOrder)
	}
}
