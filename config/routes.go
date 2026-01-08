package config

import (
	"guardpost-gin/internal/models"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(config *models.Config, engine *gin.Engine) {
	for _, route := range config.Routes {
		targetURL, _ := url.Parse(route.Downstream)
		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		// Register the route dynamically
		engine.Handle(route.Method, route.Path, func(c *gin.Context) {
			// Optional: Add gateway logic here (Auth, Logging, Rate Limiting)

			// Forward the request
			proxy.ServeHTTP(c.Writer, c.Request)
		})
	}
}
