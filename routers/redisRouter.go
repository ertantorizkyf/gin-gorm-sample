package routers

import (
	"github.com/ertantorizkyf/gin-gorm-sample/controllers"
	"github.com/gin-gonic/gin"
)

func InitRedisRouter(router *gin.Engine) *gin.Engine {
	redisRouter := router.Group("/redis")
	{
		redisRouter.GET("", controllers.RedisGet)
		redisRouter.POST("", controllers.RedisSet)
	}

	return router
}
