package router

import "github.com/zhime/monitor/router/system"

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
