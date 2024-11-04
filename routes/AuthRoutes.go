package routes

import (
	"github.com/gin-gonic/gin"
	"authRestApis/controllers/userAuth"
)


func AuthRoutes(router *gin.Engine){
	auth := router.Group("/auth")
	
	auth.POST("/sign-up", userAuth.SignUp)
	auth.POST("/sign-in", userAuth.SignIn)
	auth.GET("/verify-token", userAuth.VerifyToken)
	auth.POST("/revoke-token", userAuth.RevokeToken)
	auth.POST("/refresh-token", userAuth.RefreshToken)

}	