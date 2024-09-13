package global

import (
	"github.com/spf13/viper"
	"github.com/zhime/monitor/config"
	"gorm.io/gorm"
	"sync"
)

var (
	CONFIG config.Config
	VIPER  *viper.Viper
	DB     *gorm.DB
	WG     sync.WaitGroup
)
