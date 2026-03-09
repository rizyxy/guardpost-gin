package main

import (
	"guardpost-gin/config"
	"guardpost-gin/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	config.LoadDotEnv()

	c, _ := config.LoadRouteFile("routes.yaml")

	config.LoadRoutes(c, engine)

	config.ConnectToDB()

	routes.InitRoutes(engine)

	engine.Run()
}
