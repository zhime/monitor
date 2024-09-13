package main

import (
	"github.com/zhime/monitor/core"
	"github.com/zhime/monitor/database"
	"github.com/zhime/monitor/global"
)

func main() {
	global.VIPER = core.Viper()
	global.DB = core.InitMysql()

	global.WG.Add(1)
	go database.CheckSlaveStatus()
	global.WG.Wait()

}
