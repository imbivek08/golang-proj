package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/imbivek08/web-service/controllers"
	"github.com/imbivek08/web-service/middlewares"
)

func ProductRoute(r *gin.Engine) {

	publicGroup := r.Group("/product")
	{
		publicGroup.GET("/:id", controllers.GetProduct)
		publicGroup.GET("/", controllers.GetAllProducts)
	}
	privateGroup := r.Group("/admin")
	privateGroup.Use(middlewares.AuthMiddleware(), middlewares.AdminOnly())
	{
		privateGroup.POST("/product", controllers.CreateProduct)
		privateGroup.DELETE("/product/:id", controllers.DeleteProduct)
	}

}
