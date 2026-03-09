package controllers

import (
	"guardpost-gin/src/internal/requests"
	"guardpost-gin/src/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterController(c *gin.Context) {
	var registerRequest requests.RegisterRequest

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := repositories.CreateUser(&registerRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Register successful"})
}
