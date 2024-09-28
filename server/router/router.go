package router

import (
	"YALI/server/api"
	"github.com/gin-gonic/gin"
)

func InitRouter(Router *gin.RouterGroup) {
	{
		Router.GET("/engine/health", api.Health)
		Router.POST("/engine/RunApi", api.RunApi)
	}
}
