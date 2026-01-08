package routes

import "github.com/gin-gonic/gin"

func InitRoutes(engine *gin.Engine) {
	AuthRoutes(engine)
}
