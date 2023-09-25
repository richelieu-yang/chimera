package nacosKit

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

// NewConfigClient 创建 动态配置(config) 客户端.
/*
@param options 可以用于修改: NamespaceId...
*/
func NewConfigClient(options ...constant.ClientOption) (config_client.IConfigClient, error) {
	if clientConfig == nil || serverConfigs == nil {
		return nil, NotSetUpError
	}

	clientConfig1, err := GetClientConfigCopy(options...)
	if err != nil {
		return nil, err
	}
	return clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  clientConfig1,
		ServerConfigs: serverConfigs,
	})
}

// NewNamingClient 创建 服务发现(naming) 客户端.
/*
@param options 可以用于修改: NamespaceId...
*/
func NewNamingClient(options ...constant.ClientOption) (naming_client.INamingClient, error) {
	if clientConfig == nil || serverConfigs == nil {
		return nil, NotSetUpError
	}

	clientConfig1, err := GetClientConfigCopy(options...)
	if err != nil {
		return nil, err
	}
	return clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig:  clientConfig1,
		ServerConfigs: serverConfigs,
	})
}
