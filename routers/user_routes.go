package routers

import (
	"book_order_app/controllers"
	"book_order_app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userController := controllers.InitializeUserController()
	users := rg.Group("/users")
	{
		users.POST("/login", userController.LoginUser)
		users.POST("/register", userController.RegisterUser)

		// Protected routes
		users.GET("/profile", middleware.AuthMiddleware(), userController.GetProfile)
	}
}
