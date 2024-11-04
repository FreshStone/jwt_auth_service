package userAuth

import (
	"encoding/json"
	"io"
	"net/http"

	dataStore "authRestApis/models"
	et "authRestApis/models/entities"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var (
		req et.User
		err error
	)
	body, _ := io.ReadAll(c.Request.Body)
	err = json.Unmarshal(body, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "invalid request data",
			// "message": "Please send user_email and user_password fields",
		})
		c.Abort()
		return
	}

	dataStore.UserData[req.Email] = &et.User{Email: req.Email, Password: req.Password}

	c.JSON(http.StatusOK, gin.H{
		// "user_email": req.Email,
		"message":    "User created successfully",
	})
}
