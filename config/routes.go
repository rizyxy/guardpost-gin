package config

import (
	"guardpost-gin/src/middleware"
	"guardpost-gin/src/models"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-yaml"
)

func LoadRouteFile(filename string) (*models.Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config models.Config
	err = yaml.Unmarshal(data, &config)

	return &config, err
}

func LoadRoutes(config *models.Config, engine *gin.Engine) {

	limiter := middleware.NewIPRateLimiter(10, 10)

	for _, route := range config.Routes {
		targetURL, _ := url.Parse(route.Downstream)
		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		handlers := []gin.HandlerFunc{}

		handlers = append(handlers, middleware.LimitRate(limiter))

		if route.Protected {
			handlers = append(handlers, middleware.IsAuthenticated())
		}

		handlers = append(handlers, func(c *gin.Context) {
			proxy.ServeHTTP(c.Writer, c.Request)
		})

		engine.Handle(route.Method, route.Path, handlers...)
	}
}
