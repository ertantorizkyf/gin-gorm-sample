package routers

import (
	"github.com/ertantorizkyf/gin-gorm-sample/handlers"
	"github.com/gin-gonic/gin"
)

func InitExcelRouter(router *gin.Engine) *gin.Engine {
	articlesRouter := router.Group("/excel")
	{
		articlesRouter.GET("", handlers.GenerateExcelGet)
	}

	return router
}
