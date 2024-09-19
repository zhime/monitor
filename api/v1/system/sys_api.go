package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhime/monitor/model/response"
)

type SystemApiApi struct{}

func (s *SystemApiApi) CreateApi(c *gin.Context) {
	//var api system.SysApi
	//err := c.ShouldBindJSON(&api)
	//if err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	response.OkWithMessage("创建成功", c)
}
