package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhime/monitor/global"
	"github.com/zhime/monitor/router"
	"strconv"
)

func RunServer() {
	Router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	systemRouter := router.RouterGroupApp.System

	systemRouter.InitApiRouter(Router.Group(""))

	if err := Router.Run(fmt.Sprintf("%s:%s", global.CONFIG.Server.Host, strconv.Itoa(global.CONFIG.Server.Port))); err != nil {
		panic(err)
	}
}
