package core

import (
	"bytes"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"github.com/zhime/monitor/config"
	"github.com/zhime/monitor/global"
)

func InitNacos() {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(global.CONFIG.Nacos.IpAddr, global.CONFIG.Nacos.Port, constant.WithContextPath(global.CONFIG.Nacos.ContextPath)),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId("test"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("./logs/nacos/log"),
		constant.WithCacheDir("./logs/nacos/cache"),
		constant.WithLogLevel("debug"),
		constant.WithUsername(global.CONFIG.Nacos.UserName),
		constant.WithPassword(global.CONFIG.Nacos.Password),
	)

	// create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	//get config
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "netopstec-monitor.yml",
		Group:  "DEFAULT_GROUP",
	})

	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewBuffer([]byte(content))); err != nil {
		panic(fmt.Errorf("unable to read config: %v", err))
	}

	var conf config.Config

	if err := viper.Unmarshal(&conf); err != nil {
		panic(fmt.Errorf("unable to decode into struct: %v", err))
	}
	fmt.Println(conf.Nacos)
}
