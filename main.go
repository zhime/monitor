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

	//r := router.InitRouter()
	//
	//if err := r.Run(fmt.Sprintf("%s:%s", global.CONFIG.Server.Host, strconv.Itoa(global.CONFIG.Server.Port))); err != nil {
	//	panic(err)
	//}

}
