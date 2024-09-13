package core

import (
	"github.com/zhime/monitor/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       global.CONFIG.Mysql.Dsn(),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
