package config

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v2"
	"log"
)

// nacos服务配置
type NacosCnfig struct {
	DataId string
	Group  string
}

func NewConfigConn(nacos *NacosCnfig) (*NacosData, error) {
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "127.0.0.1",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}
	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return nil, err
	}
	config, err := configClient.GetConfig(vo.ConfigParam{
		DataId: nacos.DataId,
		Group:  nacos.Group,
	})
	if err != nil {
		return nil, err
	}
	var data = &NacosData{}
	err = yaml.Unmarshal([]byte(config), data)
	if err != nil {
		return nil, err
	}
	log.Println(data)
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: nacos.DataId,
		Group:  nacos.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}

type NacosData struct {
	Mysql Mysql `yaml:"Mysql"`
	Redis Redis `yaml:"Redis"`
}

type Mysql struct {
	Username  string `yaml:"Username"`
	Password  string `yaml:"Password"`
	Port      string `yaml:"Port"`
	Databases string `yaml:"Databases"`
}

type Redis struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
}
