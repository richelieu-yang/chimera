// Package nacosKit
package nacosKit

import (
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

// RegisterInstance 注册实例
/**
 * ClusterName	默认：DEFAULT
 * GroupName	默认：DEFAULT_GROUP
 */
func RegisterInstance(namingClient naming_client.INamingClient, param vo.RegisterInstanceParam) (bool, error) {
	return namingClient.RegisterInstance(param)
}

func UpdateInstance(namingClient naming_client.INamingClient, param vo.UpdateInstanceParam) (bool, error) {
	return namingClient.UpdateInstance(param)
}

// DeregisterInstance 注销实例
/**
 * ClusterName	默认：DEFAULT
 * GroupName	默认：DEFAULT_GROUP
 */
func DeregisterInstance(namingClient naming_client.INamingClient, instance vo.DeregisterInstanceParam) (bool, error) {
	return namingClient.DeregisterInstance(instance)
}

// SelectAllInstances 获取所有的实例列表
/*
返回全部实例列表，包括：healthy=false、enable=false、weight<=0
TODO: ??? nacos1.4.2，通过此方法无法获取到“下线”的服务，难道"下线"不等于"enable == false"？后续到 nacos2.x 最新版上再试试
*/
func SelectAllInstances(namingClient naming_client.INamingClient, param vo.SelectAllInstancesParam) ([]model.Instance, error) {
	return namingClient.SelectAllInstances(param)
}

// SelectInstances 获取实例列表
/*
只返回满足这些条件的实例列表：healthy=${HealthyOnly}、enable=true、weight>0.
！！！：返回的错误不为nil的话，需要用进行 "nacosError" 判断.（1种特殊情况）
*/
func SelectInstances(namingClient naming_client.INamingClient, param vo.SelectInstancesParam) ([]model.Instance, error) {
	return namingClient.SelectInstances(param)
}

// SelectOneHealthyInstance 获取一个健康的实例（加权随机轮询）
/*
按加权随机轮询的负载均衡策略返回一个健康的实例
实例必须满足的条件：health=true、enable=true、weight>0
！！！：返回的错误不为nil的话，需要用进行 "nacosError" 判断.（2种特殊情况，通过"nacosError.IsInstanceListEmpty()"和"nacosError.IsHealthyInstanceListEmpty()"）

@return 2个返回值，必定有1个不为nil
*/
func SelectOneHealthyInstance(namingClient naming_client.INamingClient, param vo.SelectOneHealthInstanceParam) (*model.Instance, error) {
	// 防止blank的情况
	param.GroupName = strKit.Trim(param.GroupName)

	return namingClient.SelectOneHealthyInstance(param)
}
