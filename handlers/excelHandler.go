package handlers

import (
	"github.com/ertantorizkyf/gin-gorm-sample/usecases"
	"github.com/gin-gonic/gin"
)

func GenerateExcelGet(c *gin.Context) {
	usecases.GenerateSampleExcel()

	c.JSON(200, gin.H{
		"data": "OK",
	})
}
