package routes

import "github.com/gin-gonic/gin"

func MainRoute(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Main Hello, World!")
	})
}
