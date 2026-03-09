package routes

import (
	"guardpost-gin/src/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(engine *gin.Engine) {

	router := engine.Group("/auth")

	router.POST("/login", controllers.LoginController)
	router.POST("/register", controllers.RegisterController)
	router.POST("/refresh-token", controllers.RefreshTokenController)
}
