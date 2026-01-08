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

	// Load Database
	config.ConnectToDB()

	// Init Routes
	routes.InitRoutes(engine)

	// Run Engine
	engine.Run()
}
