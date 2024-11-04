package main

import (
	"authRestApis/configs"
	"authRestApis/models"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	configs.LoadEnv()
	models.InitializeDb()
}

func main() {
	router = gin.New()

	InitializeRoutes()
	serverIP := os.Getenv("SERVER_IP")
	ginAppPort := os.Getenv("GIN_APP_PORT")
	GinServerAddress := fmt.Sprintf("%s:%s", serverIP, ginAppPort)

	router.Run(GinServerAddress)
}
