package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/imbivek08/web-service/controllers"
	"github.com/imbivek08/web-service/middlewares"
)

func CartRoutes(r *gin.Engine) {
	orderGroup := r.Group("/cart")
	orderGroup.Use(middlewares.AuthMiddleware())
	{
		orderGroup.POST("/", controllers.AddToCartHandler())
		orderGroup.GET("/", controllers.ViewCartHandler())
	}
}
