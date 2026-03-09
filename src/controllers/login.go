package controllers

import (
	"guardpost-gin/src/internal/requests"
	"guardpost-gin/src/internal/utils"
	"guardpost-gin/src/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginController(c *gin.Context) {
	var loginRequest requests.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := repositories.FindUserByEmail(loginRequest.Email)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := utils.VerifyPassword(loginRequest.Password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	accessToken, refreshToken, err := utils.GenerateAccessAndRefreshToken()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "data": gin.H{"access_token": accessToken, "refresh_token": refreshToken}})
}
