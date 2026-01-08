package controller

import (
	"guardpost-gin/internal/models"
	"guardpost-gin/internal/repository"
	requests "guardpost-gin/internal/requests/auth"
	"guardpost-gin/internal/utils"
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

	// Hash Password
	hashedPassword, err := utils.HashPassword(registerRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Create User Instance
	user := models.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: hashedPassword,
	}

	// Register User to DB
	err = repository.CreateUser(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Register successful"})
}
