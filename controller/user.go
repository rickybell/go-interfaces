package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rickybell/go-interfaces/app/repositories"
	"github.com/rickybell/go-interfaces/config"
	"github.com/rickybell/go-interfaces/services"
)

func GetUsers(c *gin.Context) {
	users, err := services.NewUserService(repositories.NewUserPostgresSqlGormRepository(config.DB, c)).All()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, &users)
}
