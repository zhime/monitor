package database

import (
	"fmt"
	"github.com/zhime/monitor/global"
	"github.com/zhime/monitor/utils"
	"time"
)

type SlaveStatus struct {
	IOThreadRunning     string `gorm:"column:Slave_IO_Running"`
	SQLThreadRunning    string `gorm:"column:Slave_SQL_Running"`
	SecondsBehindMaster int    `gorm:"column:Seconds_Behind_Master"`
	MasterHost          string `gorm:"column:Master_Host"`
}

func CheckSlaveStatus() {
	query := "SHOW SLAVE STATUS"

	defer global.WG.Done()
	for {
		var status SlaveStatus
		result := global.DB.Raw(query).Scan(&status)
		if result.Error != nil {
			fmt.Printf("查询失败: %v\n", result.Error)
			return
		}

		if status.IOThreadRunning != "Yes" || status.SQLThreadRunning != "Yes" || status.SecondsBehindMaster >= 2 {
			msg := fmt.Sprintf("从库(%s)状态异常：\nMasterHost:  %s\nIOThreadRunning:  %s\nSQLThreadRunning:  %s\nSecondsBehindMaster:  %d\n",
				global.CONFIG.Mysql.Host, status.MasterHost, status.IOThreadRunning, status.SQLThreadRunning, status.SecondsBehindMaster)
			_ = utils.SendDingTalkNotification(msg)
		}

		fmt.Printf("%s  slave(%s) status ok.", time.Now().Format("2006-01-02 15:04:05"), global.CONFIG.Mysql.Host)
		time.Sleep(600 * time.Second)
	}

}
