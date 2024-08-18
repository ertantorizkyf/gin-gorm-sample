package routers

import (
	"github.com/ertantorizkyf/gin-gorm-sample/handlers"
	"github.com/ertantorizkyf/gin-gorm-sample/usecases"
	"github.com/gin-gonic/gin"
)

func InitExcelRouter(router *gin.Engine) *gin.Engine {
	excelUsecase := usecases.NewExcelUsecase()
	excelHandler := handlers.NewExcelHandler(excelUsecase)

	excelRouter := router.Group("/excel")
	{
		excelRouter.GET("", excelHandler.GenerateExcelGet)
	}

	return router
}
