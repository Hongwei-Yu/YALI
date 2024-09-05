package initialize

import (
	"YALI/server/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {

	Routers := gin.Default()
	// 配置跨域

	groups := Routers.Group("/")
	router.InitRouter(groups)

	return Routers
}
