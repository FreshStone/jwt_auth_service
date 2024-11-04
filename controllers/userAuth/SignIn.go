package userAuth

import (
	"encoding/json"
	"io"
	"net/http"

	dataStore "authRestApis/models"
	et "authRestApis/models/entities"
	"authRestApis/utils"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	var (
		req                       et.User
		err                       error
		accessToken, refreshToken string
	)
	body, _ := io.ReadAll(c.Request.Body)
	err = json.Unmarshal(body, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request",
			// "message": "Please send user_email and user_password fields",
		})
		c.Abort()
		return
	}

	user, ok := dataStore.UserData[req.Email]
	if !ok || (user.Password != req.Password) {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "Authentication failed",
			// "message": "Invalid user_email or user_password",
		})
		c.Abort()
		return
	}

	// dataStore.UserData[req.Email] = &et.User{Email: req.Email, Password: req.Password}

	if accessToken, refreshToken, err = utils.GenerateToken(req.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate tokens",
			// "message": "Unable to create token at the moment",
		})
		c.Abort()
		return
	}

	dataStore.UserData[req.Email].AccessToken = accessToken
	dataStore.UserData[req.Email].RefreshToken = refreshToken

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"message":       "User signed in successfully",
	})
}
