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
	r := gin.Default()

	r.GET("/articles", controllers.ArticleIndex)
	r.GET("/articles/:id", controllers.ArticleDetail)
	r.POST("/articles", controllers.ArticleCreate)
	r.PUT("/articles/:id", controllers.ArticleUpdate)
	r.DELETE("/articles/:id", controllers.ArticleDelete)

	r.Run()
}
