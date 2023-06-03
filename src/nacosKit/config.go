// Package nacosKit
/*
config（动态配置）
api说明：
	https://github.com/nacos-group/nacos-sdk-go/blob/master/README_CN.md#%E5%8A%A8%E6%80%81%E9%85%8D%E7%BD%AE
官方demo：
	https://github.com/nacos-group/nacos-sdk-go/blob/master/example/config/main.go
*/
package nacosKit

import (
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/core/strKit"
)

// PublishConfig 发布配置
/*
@param dataId		不能为""，且中间不能有空格
@param group		不能为""，且中间不能有空格
@param content		不能为blank
@param configType	配置的类型

PS: 如果配置已存在，将覆盖.
*/
func PublishConfig(client config_client.IConfigClient, dataId, group, content string, configType vo.ConfigType) (bool, error) {
	if client == nil {
		return false, errorKit.New("client is nil")
	}
	// content为blank，实际上就是删除配置
	if strKit.IsBlank(content) {
		return DeleteConfig(client, dataId, group)
	}
	return client.PublishConfig(vo.ConfigParam{
		DataId:  dataId,
		Group:   group,
		Content: content,
		Type:    configType,
	})
}

// DeleteConfig 删除配置
/*
@param dataId	不能为""
@param group	不能为""

PS: 如果配置不存在，将返回(true, nil).
*/
func DeleteConfig(client config_client.IConfigClient, dataId, group string) (bool, error) {
	if client == nil {
		return false, errorKit.New("client is nil")
	}

	return client.DeleteConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
}

// GetConfig 获取配置
/*
@param dataId	不能为""
@param group	不能为""

PS:
(1) 如果不存在对应的配置，将返回("", nil)；
(2) 发布完立即获取配置，且之前未发布过dataId、group一致的配置，此时将得到"".（需要等一会才会发布生效，但实际场景中不存在此种情况）
*/
func GetConfig(client config_client.IConfigClient, dataId, group string) (string, error) {
	if client == nil {
		return "", errorKit.New("client is nil")
	}

	return client.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
}

// ListenConfig 监听配置变化
/*
@param dataId	不能为""
@param group	不能为""

PS:
(1) 监听时，对应配置可以不存在；
(2) 如果发布内容与当前内容一致，不会触发回调；
(3) 能监听delete操作，如果原先配置存在，delete将触发回调，但二次delete将不会触发回调.
*/
func ListenConfig(client config_client.IConfigClient, dataId, group string, callback func(namespaceId, dataId, group, content string)) error {
	if client == nil {
		return errorKit.New("client is nil")
	}

	return client.ListenConfig(vo.ConfigParam{
		DataId:   dataId,
		Group:    group,
		OnChange: callback,
	})
}

// CancelListenConfig 取消配置监听
/*
PS: 如果未监听，将返回nil.
*/
func CancelListenConfig(client config_client.IConfigClient, dataId, group string) error {
	if client == nil {
		return errorKit.New("client is nil")
	}

	return client.CancelListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
}

// SearchConfig 搜索配置
/*
TODO: 待研究，目前没遇到相关使用场景
Deprecated
@param search require search=accurate--精确搜索  search=blur--模糊搜索
*/
func SearchConfig(client config_client.IConfigClient, search, dataId, group string, pageNo, pageSize int) (*model.ConfigPage, error) {
	return client.SearchConfig(vo.SearchConfigParam{
		Search:   search,
		DataId:   dataId,
		Group:    group,
		PageNo:   pageNo,
		PageSize: pageSize,
	})
}
