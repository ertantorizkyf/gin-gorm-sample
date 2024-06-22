package main

import (
	"github.com/ertantorizkyf/gin-gorm-sample/controllers"
	"github.com/ertantorizkyf/gin-gorm-sample/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()

	articlesRouter := router.Group("/articles")
	{
		articlesRouter.GET("", controllers.ArticleIndex)
		articlesRouter.GET(":id", controllers.ArticleDetail)
		articlesRouter.POST("", controllers.ArticleCreate)
		articlesRouter.PUT(":id", controllers.ArticleUpdate)
		articlesRouter.DELETE(":id", controllers.ArticleDelete)

		articlesJsonRouter := articlesRouter.Group("/json")
		{
			articlesJsonRouter.GET("/structured", controllers.ArticleJsonStructuredIndex)
			articlesJsonRouter.GET("/unstructured", controllers.ArticleJsonUnstructuredIndex)
		}
	}

	redisRouter := router.Group("/redis")
	{
		redisRouter.GET("", controllers.RedisGet)
		redisRouter.POST("", controllers.RedisSet)
	}

	router.Run()
}
