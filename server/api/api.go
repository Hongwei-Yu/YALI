package api

import (
	"YALI/engine/http"
	"YALI/global"
	"YALI/server/service"
	"github.com/gin-gonic/gin"
	HTTP "net/http"
)

func RunApi(c *gin.Context) {
	var api http.ApiHttp

	err := c.ShouldBindJSON(&api)
	api.RequestHttp.Debug = api.Api.Debug
	//fmt.Println(api)
	if err != nil {
		global.ReturnMsg(c, 500, "API参数错误", err.Error())
	}

	//var resp string
	result, resp := service.DisposeApi(&api)
	if result != true {
		global.ReturnMsg(c, 500, "API执行错误", resp)
		return
	}
	global.ReturnMsg(c, 200, "API执行成功", resp)

}

func RunApiGroup(c *gin.Context) {

}

func RunScena(c *gin.Context) {

}
func Health(c *gin.Context) {
	c.JSON(HTTP.StatusOK, gin.H{"status": global.GlobalEngine.Status,
		"engineId":     global.GlobalEngine.Node,
		"engineStatus": global.GlobalEngine.HostStatus,
		"EngineHost":   global.GlobalEngine.Host,
	})
}
