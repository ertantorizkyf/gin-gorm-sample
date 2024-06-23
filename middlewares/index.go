package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RejectAuthorization(c *gin.Context) {
	message := "USER UNAUTHORIZED!"
	log.Printf("[ERR] %s", message)
	c.JSON(401, gin.H{
		"message": message,
	})
	c.Abort()
}
