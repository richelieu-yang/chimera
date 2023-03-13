package rocketmq5Kit

import (
	rmq_client "github.com/apache/rocketmq-clients/golang"
	v2 "github.com/apache/rocketmq-clients/golang/protocol/v2"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/richelieu42/go-scales/src/netKit"
	"github.com/richelieu42/go-scales/src/reflectKit"
)

// polyfillProducer
/*
Deprecated: 临时方法，因为官方依赖不支持连接RocketMQ5集群（pkg/utils/utils.go）
*/
func polyfillProducer(producer rmq_client.Producer, endpointsStr string) (rmq_client.Producer, error) {
	s := strKit.Split(endpointsStr, ";")
	if len(s) < 2 {
		return producer, nil
	}

	// pSetting/endpoints
	v, err := reflectKit.GetNestedField(producer, "pSetting", "endpoints")
	if err != nil {
		return nil, err
	}
	endpoints := (*v2.Endpoints)(v.UnsafePointer())
	if err := reviseEndpoints(endpoints, s); err != nil {
		return nil, err
	}

	// cli/accessPoint
	v, err = reflectKit.GetNestedField(producer, "cli", "accessPoint")
	if err != nil {
		return nil, err
	}
	endpoints = (*v2.Endpoints)(v.UnsafePointer())
	if err := reviseEndpoints(endpoints, s); err != nil {
		return nil, err
	}

	return producer, nil
}

// polyfillConsumer
/*
Deprecated: 临时方法，因为官方依赖不支持连接RocketMQ5集群（pkg/utils/utils.go）
*/
func polyfillConsumer(consumer rmq_client.SimpleConsumer, endpointsStr string) (rmq_client.SimpleConsumer, error) {
	s := strKit.Split(endpointsStr, ";")
	if len(s) < 2 {
		return consumer, nil
	}

	// scSettings/endpoints
	v, err := reflectKit.GetNestedField(consumer, "scSettings", "endpoints")
	if err != nil {
		return nil, err
	}
	endpoints := (*v2.Endpoints)(v.UnsafePointer())
	if err := reviseEndpoints(endpoints, s); err != nil {
		return nil, err
	}

	// cli/accessPoint
	v, err = reflectKit.GetNestedField(consumer, "cli", "accessPoint")
	if err != nil {
		return nil, err
	}
	endpoints = (*v2.Endpoints)(v.UnsafePointer())
	if err := reviseEndpoints(endpoints, s); err != nil {
		return nil, err
	}

	return consumer, nil
}

func reviseEndpoints(endpoints *v2.Endpoints, s []string) error {
	addresses := make([]*v2.Address, 0, len(s))

	for _, str := range s {
		addr, err := netKit.ParseToAddress(str)
		if err != nil {
			return err
		}
		addresses = append(addresses, &v2.Address{
			Host: addr.Hostname,
			Port: int32(addr.Port),
		})
	}
	endpoints.Addresses = addresses
	return nil
}
