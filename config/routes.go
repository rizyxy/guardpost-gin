package config

import (
	"guardpost-gin/internal/middleware"
	"guardpost-gin/internal/models"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(config *models.Config, engine *gin.Engine) {
	for _, route := range config.Routes {
		targetURL, _ := url.Parse(route.Downstream)
		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		handlers := []gin.HandlerFunc{}

		// Authentication Middleware
		if route.Protected {
			handlers = append(handlers, middleware.IsAuthenticated())
		}

		// Proxy Handler
		handlers = append(handlers, func(c *gin.Context) {
			proxy.ServeHTTP(c.Writer, c.Request)
		})

		engine.Handle(route.Method, route.Path, handlers...)
	}
}
