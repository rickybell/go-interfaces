package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rickybell/go-interfaces/config"
	"github.com/rickybell/go-interfaces/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	routes.MainRoute(router)
	router.Run(":8080")
}
