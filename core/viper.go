package core

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zhime/monitor/global"
)

func Viper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.WatchConfig()

	if err := viper.Unmarshal(&global.CONFIG); err != nil {
		panic(err)
	}
}
