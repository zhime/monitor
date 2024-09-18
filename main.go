package main

import (
	"github.com/zhime/monitor/core"
	"github.com/zhime/monitor/global"
)

func main() {
	core.Viper()
	global.DB = core.InitMysql()
	//_ = global.DB.AutoMigrate(&model.User{})

	//global.WG.Add(1)
	//go database.CheckSlaveStatus()
	//global.WG.Wait()

	//r := gin.Default()
	//
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"data": "ok",
	//	})
	//})
	//if err := r.Run(); err != nil {
	//	panic(err)
	//}

}
