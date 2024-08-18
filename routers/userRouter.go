package routers

import (
	"github.com/ertantorizkyf/gin-gorm-sample/handlers"
	"github.com/ertantorizkyf/gin-gorm-sample/middlewares"
	"github.com/ertantorizkyf/gin-gorm-sample/repositories"
	"github.com/ertantorizkyf/gin-gorm-sample/usecases"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.Engine) *gin.Engine {
	userRepository := repositories.NewUserRepository()
	userUsecase := usecases.NewUserUsecase(userRepository)
	userHandler := handlers.NewUserHandler(userUsecase)

	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", userHandler.Register)
		userRouter.POST("/login", userHandler.Login)
		userRouter.GET("/sample-middleware-impl", middlewares.AuthorizeUser, userHandler.SampleMiddlewareImpl)
	}

	return router
}
