package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/imbivek08/web-service/controllers"
)

func AuthRutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
}
