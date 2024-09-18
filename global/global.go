package global

import (
	"github.com/zhime/monitor/config"
	"gorm.io/gorm"
	"sync"
)

var (
	CONFIG config.Config
	DB     *gorm.DB
	WG     sync.WaitGroup
)
