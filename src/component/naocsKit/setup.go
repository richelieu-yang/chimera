package nacosKit

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

var (
	clientConfig  *constant.ClientConfig
	serverConfigs []constant.ServerConfig
)

func MustSetUp() {
	err := SetUp()
	if err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func SetUp() (err error) {
	defer func() {
		if err != nil {
			clientConfig = nil
			serverConfigs = nil
		}
	}()

	return
}

// NewNamingClient 创建 服务发现(naming) 客户端.
func NewNamingClient() (naming_client.INamingClient, error) {
	if clientConfig == nil || serverConfigs == nil {
		return nil, UninitializedError
	}

	return clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
}

// NewConfigClient 创建 动态配置(config) 客户端.
func NewConfigClient() (config_client.IConfigClient, error) {
	if clientConfig == nil || serverConfigs == nil {
		return nil, UninitializedError
	}

	return clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
}
