package routes

import (
	controller "guardpost-gin/internal/controllers/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(engine *gin.Engine) {

	router := engine.Group("/auth")

	router.POST("/login", controller.LoginController)
	router.POST("/register", controller.RegisterController)
	router.POST("/refresh-token", controller.RefreshTokenController)
}
