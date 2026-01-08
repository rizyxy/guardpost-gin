package controller

import (
	requests "guardpost-gin/internal/requests/auth"
	"guardpost-gin/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RefreshTokenController(c *gin.Context) {
	var refreshRequest *requests.RefreshRequest

	// Bind JSON Request
	if err := c.ShouldBindJSON(&refreshRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify Refresh Token
	if !utils.VerifyJWT(refreshRequest.RefreshToken) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Generate Token
	accessToken, refreshToken, err := utils.GenerateToken()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token refreshed successfully", "data": gin.H{"access_token": accessToken, "refresh_token": refreshToken}})
}
