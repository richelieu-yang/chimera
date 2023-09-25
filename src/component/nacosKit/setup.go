package nacosKit

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/richelieu-yang/chimera/v2/src/copyKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/intKit"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/netKit"
	"github.com/richelieu-yang/chimera/v2/src/urlKit"
	"github.com/sirupsen/logrus"
	"net/url"
)

var (
	clientConfig  *constant.ClientConfig
	serverConfigs []constant.ServerConfig
)

func MustSetUp(config Config, options ...constant.ClientOption) {
	err := SetUp(config, options...)
	if err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

// SetUp
/*
@param options		!!!:
					(1) 建议配置 客户端的缓存目录（default value is current path）	constant.WithCacheDir
					(2) 建议配置 客户端的日志目录（default is current path）			constant.WithLogDir
					(3) 建议配置 客户端的日志级别（default value is info）			constant.WithLogLevel
*/
func SetUp(config Config, options ...constant.ClientOption) (err error) {
	defer func() {
		if err != nil {
			clientConfig = nil
			serverConfigs = nil
		}
	}()

	/* (1) clientConfig */
	options1 := []constant.ClientOption{
		constant.WithNamespaceId(config.NamespaceId),
	}
	options1 = append(options1, options...)
	clientConfig = constant.NewClientConfig(options1...)

	/* (2) serverConfigs */
	if err = sliceKit.AssertNotEmpty(config.Addresses, "config.Addresses"); err != nil {
		return
	}
	for _, addr := range config.Addresses {
		if strKit.IsBlank(addr) {
			continue
		}

		var u *url.URL
		u, err = urlKit.Parse(addr)
		if err != nil {
			err = errorKit.Wrap(err, "fail to parse address(%s)", addr)
			return
		}
		var port uint64
		port, err = intKit.ToUint64E(u.Port())
		if err != nil {
			err = errorKit.Wrap(err, "invalid address(%s) with port string(%s)", addr, u.Port())
			return
		}
		if err = netKit.AssertValidPort(int(port)); err != nil {
			err = errorKit.Wrap(err, "invalid address(%s)  with port(%d)", addr, port)
			return
		}

		serverConfigs = append(serverConfigs, constant.ServerConfig{
			Scheme:      u.Scheme,
			IpAddr:      u.Hostname(),
			Port:        port,
			ContextPath: u.Path,
		})
	}
	if sliceKit.IsEmpty(serverConfigs) {
		err = errorKit.New("no valid address")
		return
	}

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
