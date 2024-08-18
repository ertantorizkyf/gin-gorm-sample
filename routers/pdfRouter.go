package routers

import (
	"github.com/ertantorizkyf/gin-gorm-sample/handlers"
	"github.com/ertantorizkyf/gin-gorm-sample/usecases"
	"github.com/gin-gonic/gin"
)

func InitPDFRouter(router *gin.Engine) *gin.Engine {
	pdfUsecase := usecases.NewPdfUsecase()
	pdfHandler := handlers.NewPdfHandler(pdfUsecase)

	pdfRouter := router.Group("/pdf")
	{
		pdfRouter.GET("", pdfHandler.GeneratePDFGet)
	}

	return router
}
