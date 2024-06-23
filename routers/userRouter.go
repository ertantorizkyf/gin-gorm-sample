package routers

import (
	"github.com/ertantorizkyf/gin-gorm-sample/controllers"
	"github.com/ertantorizkyf/gin-gorm-sample/middlewares"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.Engine) *gin.Engine {
	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", controllers.Register)
		userRouter.POST("/login", controllers.Login)
		userRouter.GET("/sample-middleware-impl", middlewares.AuthorizeUser, controllers.SampleMiddlewareImpl)
	}

	return router
}
