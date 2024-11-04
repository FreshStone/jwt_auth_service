package main

import(
	"authRestApis/routes"
)

func InitializeRoutes() {
	routes.AuthRoutes(router)
}