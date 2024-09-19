package main

import (
	"github.com/zhime/monitor/core"
)

func main() {
	core.Viper()
	core.InitNacos()
	//global.DB = core.InitMysql()
	//_ = global.DB.AutoMigrate(&model.User{})

	//global.WG.Add(1)
	//go database.CheckSlaveStatus()
	//global.WG.Wait()

	core.RunServer()

}
