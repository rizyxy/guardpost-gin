package controller

import (
	"guardpost-gin/internal/repository"
	requests "guardpost-gin/internal/requests/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterController(c *gin.Context) {
	var registerRequest requests.RegisterRequest

	// Bind JSON Request
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Register User to DB
	err := repository.CreateUser(&registerRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Register successful"})
}
