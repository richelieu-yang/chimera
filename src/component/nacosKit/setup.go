package nacosKit

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/richelieu-yang/chimera/v2/src/copyKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
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

func SetUp(config Config, clientLogDir, clientCacheDir string, options ...constant.ClientOption) (err error) {
	clientConfig = *constant.NewClientConfig(
		constant.WithNamespaceId("e525eafa-f7d7-4029-83d9-008937f9d468"), //当namespace是public时，此处填空字符串。
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	defer func() {
		if err != nil {
			clientConfig = nil
			serverConfigs = nil
		}
	}()

	return
}

// NewNamingClient
/*
PS:
*/
func NewNamingClient() (naming_client.INamingClient, error) {
	return NewNamingClientWithNamespaceId(nil)
}

// NewNamingClientWithNamespaceId 创建 服务发现(naming) 客户端.
/*
@param namespaceId 可以为""（此时namespace是public）
*/
func NewNamingClientWithNamespaceId(namespaceId *string) (naming_client.INamingClient, error) {
	if clientConfig == nil || serverConfigs == nil {
		return nil, NotSetUpError
	}

	clientConfig1, err := copyKit.DeepCopy(clientConfig)
	if err != nil {
		return nil, errorKit.Wrap(err, "fail to deep copy")
	}
	if namespaceId != nil {
		clientConfig1.NamespaceId = *namespaceId
	}
	return clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  clientConfig1,
			ServerConfigs: serverConfigs,
		},
	)
}

func NewConfigClient() (config_client.IConfigClient, error) {
	return NewConfigClientWithNamespaceId(nil)
}

// NewConfigClientWithNamespaceId 创建 动态配置(config) 客户端.
/*
@param namespaceId 可以为""（此时namespace是public）
*/
func NewConfigClientWithNamespaceId(namespaceId *string) (config_client.IConfigClient, error) {
	if clientConfig == nil || serverConfigs == nil {
		return nil, NotSetUpError
	}

	clientConfig1, err := copyKit.DeepCopy(clientConfig)
	if err != nil {
		return nil, errorKit.Wrap(err, "fail to deep copy")
	}
	if namespaceId != nil {
		clientConfig1.NamespaceId = *namespaceId
	}
	return clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  clientConfig1,
			ServerConfigs: serverConfigs,
		},
	)
}
