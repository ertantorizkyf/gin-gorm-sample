package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.NoRoute(func(c *gin.Context) {
		errMethod := c.Request.Method
		errPath := c.Request.URL.Path
		errMessage := fmt.Sprintf("Path [%s] %s not found!", errMethod, errPath)

		c.JSON(404, gin.H{"message": errMessage})
	})

	return router
}
