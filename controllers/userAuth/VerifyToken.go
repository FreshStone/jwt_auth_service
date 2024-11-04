package userAuth

import (
	"net/http"
	"os"
	"strings"

	"authRestApis/utils"

	"github.com/gin-gonic/gin"
)

func VerifyToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	userEmail, err := utils.ValidateToken(tokenString, os.Getenv("ACCESS_SECRET"))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Token is valid",
		"user_email": userEmail,
	})
}
