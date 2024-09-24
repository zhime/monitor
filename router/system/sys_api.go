package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zhime/monitor/api/v1"
)

type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemApiApi

	{
		Router.POST("createApi", apiRouterApi.CreateApi) // 创建Api
	}
}
