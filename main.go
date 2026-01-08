package main

import (
	"guardpost-gin/config"
	"guardpost-gin/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init Gin Engine
	engine := gin.Default()

	// Load Dotenv
	config.LoadDotEnv()

	// Load Config
	c, _ := config.LoadConfig("routes.yaml")

	// Load Routes
	config.LoadRoutes(c, engine)

	// Load Database
	config.ConnectToDB()

	// Init Internal Routes
	routes.InitRoutes(engine)

	// Run Engine
	engine.Run()
}
