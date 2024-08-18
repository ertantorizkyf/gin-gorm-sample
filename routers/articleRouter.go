package routers

import (
	"github.com/ertantorizkyf/gin-gorm-sample/controllers"
	"github.com/ertantorizkyf/gin-gorm-sample/handlers"
	"github.com/gin-gonic/gin"
)

func InitArticleRouter(router *gin.Engine) *gin.Engine {
	articlesRouter := router.Group("/articles")
	{
		articlesRouter.GET("", controllers.ArticleIndex)
		articlesRouter.GET(":id", controllers.ArticleDetail)
		articlesRouter.POST("", controllers.ArticleCreate)
		articlesRouter.PUT(":id", controllers.ArticleUpdate)
		articlesRouter.DELETE(":id", controllers.ArticleDelete)

		articlesJsonRouter := articlesRouter.Group("/json")
		{
			articlesJsonRouter.GET("/structured", handlers.ArticleJsonStructuredIndex)
			articlesJsonRouter.GET("/unstructured", handlers.ArticleJsonUnstructuredIndex)
		}
	}

	return router
}
