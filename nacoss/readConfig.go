package nacoss

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"
)

type ServerConfig struct {
	Name       string       `mapstructure:"name" json:"name"`
	Host       string       `mapstructure:"host" json:"host"`
	Tags       []string     `mapstructure:"tags" json:"tags"`
	MysqlInfo  MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"`
	EsInfo     EsConfig     `mapstructure:"es" json:"es"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}
type EsConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

var (
	Server ServerConfig
)

// ReadConfig 读取配置 goods_srv.json dev
func ReadConfig(dataId string, groUp string) *ServerConfig {
	//从哪个nacos中进行获取配置信息
	sc := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1", // 主机公网地址
			Port:   8848,        // nacos 地址
		},
	}
	// 获取的是哪个命名空间下的配置信息
	cc := constant.ClientConfig{
		NamespaceId:         "9bfd42ff-31b4-4f0a-990b-ca737996b3b8", // 需要映射的命名空间ID
		TimeoutMs:           5000,                                   //推出时间П
		NotLoadCacheAtStart: true,                                   //不启动加载缓存
		LogLevel:            "debug",                                // 日志等级
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc, // 服务器
		"clientConfig":  cc, // 客户端
	})
	if err != nil {
		panic(err)
	}
	// 从刚刚创建的nacos节点中中获取配置信息
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId, // 数据ID
		Group:  groUp,  // 分组
	})

	if err != nil {
		panic(err)
	}
	// 因为获取到的是json 格式需要转格式
	err = json.Unmarshal([]byte(content), &Server)
	if err != nil {
		zap.S().Fatalf("读取nacos配置失败： %s", err.Error())
	}
	return &Server
}
