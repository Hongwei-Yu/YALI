package api

import (
	"YALI/engine/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RunApi(c *gin.Context) {

}

func RunApiGroup(c *gin.Context) {

}

func RunScena(c *gin.Context) {

}
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": global.GlobalEngine.Status,
		"engineId":     global.GlobalEngine.Node,
		"engineStatus": global.GlobalEngine.HostStatus,
		"EngineHost":   global.GlobalEngine.Host,
	})
}
