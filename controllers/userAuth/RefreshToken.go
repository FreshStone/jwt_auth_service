package userAuth

import (
	"authRestApis/models"
	"authRestApis/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
	}
	userEmail, err := utils.ValidateToken(req.RefreshToken, os.Getenv("REFRESH_SECRET"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid refresh token",
		})
		c.Abort()
		return
	}

	accessToken, refreshToken, err := utils.GenerateToken(userEmail)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error": "failed to generate tokens",
			})
		c.Abort()
		return
	}

	models.UserData[userEmail].AccessToken = accessToken
	models.UserData[userEmail].RefreshToken = refreshToken

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}
