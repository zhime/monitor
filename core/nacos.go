package core

import (
	"fmt"
	"github.com/zhime/monitor/global"
	"time"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func InitNacos() {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(global.CONFIG.Nacos.IpAddr, global.CONFIG.Nacos.Port, constant.WithContextPath(global.CONFIG.Nacos.ContextPath)),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
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

	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-data",
		Group:   "test-group",
		Content: "hello world!",
	})
	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-data-2",
		Group:   "test-group",
		Content: "hello world!",
	})
	if err != nil {
		fmt.Printf("PublishConfig err:%+v \n", err)
	}
	time.Sleep(1 * time.Second)
	//get config
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "test-data",
		Group:  "test-group",
	})
	fmt.Println("GetConfig,config :" + content)

	//Listen config change,key=dataId+group+namespaceId.
	err = client.ListenConfig(vo.ConfigParam{
		DataId: "test-data",
		Group:  "test-group",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
		},
	})

	err = client.ListenConfig(vo.ConfigParam{
		DataId: "test-data-2",
		Group:  "test-group",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
		},
	})

	time.Sleep(1 * time.Second)

	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-data",
		Group:   "test-group",
		Content: "test-listen",
	})

	time.Sleep(1 * time.Second)

	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-data-2",
		Group:   "test-group",
		Content: "test-listen",
	})

	time.Sleep(2 * time.Second)

	time.Sleep(1 * time.Second)
	_, err = client.DeleteConfig(vo.ConfigParam{
		DataId: "test-data",
		Group:  "test-group",
	})
	time.Sleep(1 * time.Second)

	//cancel config change
	err = client.CancelListenConfig(vo.ConfigParam{
		DataId: "test-data",
		Group:  "test-group",
	})

	searchPage, _ := client.SearchConfig(vo.SearchConfigParam{
		Search:   "blur",
		DataId:   "",
		Group:    "",
		PageNo:   1,
		PageSize: 10,
	})
	fmt.Printf("Search config:%+v \n", searchPage)
}
