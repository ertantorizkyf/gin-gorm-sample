package middlewares

import (
	"strings"

	"github.com/ertantorizkyf/gin-gorm-sample/helpers"
	"github.com/gin-gonic/gin"
)

func AuthorizeUser(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		RejectAuthorization(c)

		return
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) < 2 {
		RejectAuthorization(c)

		return
	}

	err := helpers.VerifyToken(authToken[1])
	if err != nil {
		RejectAuthorization(c)

		return
	}

	c.Next()
}
