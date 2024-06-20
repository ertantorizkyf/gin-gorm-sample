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

	articles := router.Group("/articles")
	{
		articles.GET("", controllers.ArticleIndex)
		articles.GET(":id", controllers.ArticleDetail)
		articles.POST("", controllers.ArticleCreate)
		articles.PUT(":id", controllers.ArticleUpdate)
		articles.DELETE(":id", controllers.ArticleDelete)
	}

	router.Run()
}
