package handlers

import (
	"github.com/ertantorizkyf/gin-gorm-sample/usecases"
	"github.com/gin-gonic/gin"
)

func GeneratePDFGet(c *gin.Context) {
	usecases.GenerateSamplePDF()

	c.JSON(200, gin.H{
		"data": "OK",
	})
}
