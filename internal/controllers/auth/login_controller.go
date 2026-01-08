package controller

import (
	"guardpost-gin/internal/repository"
	requests "guardpost-gin/internal/requests/auth"
	"guardpost-gin/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginController(c *gin.Context) {
	var loginRequest requests.LoginRequest

	// Bind JSON Request
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Find User by Email
	user := repository.FindUserByEmail(loginRequest.Email)

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid credentials"})
		return
	}

	// Verify Password
	if err := utils.VerifyPassword(loginRequest.Password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate Token
	accessToken, refreshToken, err := utils.GenerateToken()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "data": gin.H{"access_token": accessToken, "refresh_token": refreshToken}})
}
