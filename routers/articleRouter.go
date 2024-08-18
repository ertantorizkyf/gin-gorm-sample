package routers

import (
	"github.com/ertantorizkyf/gin-gorm-sample/handlers"
	"github.com/ertantorizkyf/gin-gorm-sample/repositories"
	"github.com/ertantorizkyf/gin-gorm-sample/usecases"
	"github.com/gin-gonic/gin"
)

func InitArticleRouter(router *gin.Engine) *gin.Engine {
	articleRepository := repositories.NewArticleRepository()
	articleUsecase := usecases.NewArticleUsecase(articleRepository)
	articleHandler := handlers.NewArticleHandler(articleUsecase)

	articlesRouter := router.Group("/articles")
	{
		articlesRouter.GET("", articleHandler.ArticleIndex)
		articlesRouter.GET(":id", articleHandler.ArticleDetail)
		articlesRouter.POST("", articleHandler.ArticleCreate)
		articlesRouter.PUT(":id", articleHandler.ArticleUpdate)
		articlesRouter.DELETE(":id", articleHandler.ArticleDelete)

		articlesJsonRouter := articlesRouter.Group("/json")
		{
			articlesJsonRouter.GET("/structured", handlers.ArticleJsonStructuredIndex)
			articlesJsonRouter.GET("/unstructured", handlers.ArticleJsonUnstructuredIndex)
		}
	}

	return router
}
