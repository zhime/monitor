package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//global.VIPER = core.Viper()
	//global.DB = core.InitMysql()

	//global.WG.Add(1)
	//go database.CheckSlaveStatus()
	//global.WG.Wait()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "ok",
		})
	})
	if err := r.Run(); err != nil {
		panic(err)
	}

}
