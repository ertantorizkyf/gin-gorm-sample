package routers

import (
	"github.com/ertantorizkyf/gin-gorm-sample/handlers"
	"github.com/ertantorizkyf/gin-gorm-sample/helpers"
	"github.com/ertantorizkyf/gin-gorm-sample/usecases"
	"github.com/gin-gonic/gin"
)

func InitRedisRouter(router *gin.Engine) *gin.Engine {
	redisUsecase := usecases.NewRedisUsecase(helpers.RedisClient())
	redisHandler := handlers.NewRedisHandler(redisUsecase)

	redisRouter := router.Group("/redis")
	{
		redisRouter.GET("", redisHandler.RedisGet)
		redisRouter.POST("", redisHandler.RedisSet)
	}

	return router
}
