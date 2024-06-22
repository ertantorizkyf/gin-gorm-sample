package middlewares

import (
	"log"
	"strings"

	"github.com/ertantorizkyf/gin-gorm-sample/helpers"
	"github.com/gin-gonic/gin"
)

func AuthorizeUser(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	authToken := strings.Split(authHeader, " ")

	if authHeader == "" || len(authToken) < 2 {
		message := "USER UNAUTHORIZED!"
		log.Printf("[ERR] %s", message)
		c.JSON(500, gin.H{
			"message": message,
		})

		c.Abort()
	}

	err := helpers.VerifyToken(authToken[1])
	if err != nil {
		message := "USER UNAUTHORIZED!"
		log.Printf("[ERR] %s", message)
		c.JSON(500, gin.H{
			"message": message,
		})

		c.Abort()
	}

	c.Next()
}
