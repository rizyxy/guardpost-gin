package controllers

import (
	"guardpost-gin/src/internal/requests"
	"guardpost-gin/src/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RefreshTokenController(c *gin.Context) {
	var refreshRequest *requests.RefreshRequest

	if err := c.ShouldBindJSON(&refreshRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !utils.VerifyJWT(refreshRequest.RefreshToken) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	accessToken, refreshToken, err := utils.GenerateAccessAndRefreshToken()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token refreshed successfully", "data": gin.H{"access_token": accessToken, "refresh_token": refreshToken}})
}
