package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/imbivek08/web-service/controllers"
	"github.com/imbivek08/web-service/middlewares"
)

func OrderRoutes(r *gin.Engine) {
	orderGroup := r.Group("/order")
	orderGroup.Use(middlewares.AuthMiddleware())
	{
		orderGroup.POST("/", controllers.CreateOrder)
		orderGroup.DELETE("/delete", controllers.DeleteOrder)
		orderGroup.PUT("/update", controllers.UpdateOrder)
	}
}
