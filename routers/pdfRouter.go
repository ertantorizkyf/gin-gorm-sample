package routers

import (
	"github.com/ertantorizkyf/gin-gorm-sample/handlers"
	"github.com/gin-gonic/gin"
)

func InitPDFRouter(router *gin.Engine) *gin.Engine {
	pdfRouter := router.Group("/pdf")
	{
		pdfRouter.GET("", handlers.GeneratePDFGet)
	}

	return router
}
