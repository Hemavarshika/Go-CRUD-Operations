package main

import (
	"crud/config"
	"crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	router := gin.New()
	routes.UserRoutes(router)
	router.Run(":8080")
}
