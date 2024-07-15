package main

import (
	"github.com/ertantorizkyf/gin-gorm-sample/initializers"
	"github.com/ertantorizkyf/gin-gorm-sample/routers"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	router := routers.InitRouter()
	router = routers.InitArticleRouter(router)
	router = routers.InitRedisRouter(router)
	router = routers.InitUserRouter(router)
	router = routers.InitExcelRouter(router)

	router.Run()
}
