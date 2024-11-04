package userAuth

import (
	"net/http"
	"os"
	"strings"
	"time"

	"authRestApis/models"
	"authRestApis/utils"

	"github.com/gin-gonic/gin"
)

func RevokeToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	_, err := utils.ValidateToken(tokenString, os.Getenv("ACCESS_SECRET"))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	models.BlacklistedTokens.Lock()
	models.BlacklistedTokens.TokenMap[tokenString] = time.Now()
	models.BlacklistedTokens.Unlock()

	c.JSON(http.StatusOK, gin.H{
		"message": "Token has been revoked",
	})
}
