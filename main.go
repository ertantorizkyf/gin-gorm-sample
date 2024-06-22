package main

import (
	"fmt"

	"github.com/ertantorizkyf/gin-gorm-sample/controllers"
	"github.com/ertantorizkyf/gin-gorm-sample/initializers"
	"github.com/ertantorizkyf/gin-gorm-sample/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		errMethod := c.Request.Method
		errPath := c.Request.URL.Path
		errMessage := fmt.Sprintf("Path [%s] %s not found!", errMethod, errPath)

		c.JSON(404, gin.H{"message": errMessage})
	})

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

	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", controllers.Register)
		userRouter.POST("/login", controllers.Login)
		userRouter.GET("/sample-middleware-impl", middlewares.AuthorizeUser, controllers.SampleMiddlewareImpl)
	}

	router.Run()
}
