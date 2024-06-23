package tests

import (
	"os"

	"github.com/ertantorizkyf/gin-gorm-sample/routers"
	"github.com/gin-gonic/gin"
)

func InitTestEnv() {
	os.Setenv("JWT_SECRET", "secret-key")
}

func InitTestRouter() *gin.Engine {
	router := routers.InitRouter()
	router = routers.InitArticleRouter(router)
	router = routers.InitRedisRouter(router)
	router = routers.InitUserRouter(router)

	return router
}
