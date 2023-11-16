package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rickybell/go-interfaces/controller"
)

func UserRoute(router *gin.Engine) {
	router.GET("/users", controller.GetUsers)
	// router.POST("/users", controller.CreateUserController)
	// router.DELETE("/users/:id", controller.DeleteUserController)
	// router.PUT("/users/:id", controller.UpdateUserController)
}
