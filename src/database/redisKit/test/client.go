package test

import "gitee.com/richelieu042/go-scales/src/database/redisKit"

// NewSingleNodeClient 单节点集群的客户端
func NewSingleNodeClient() (*redisKit.Client, error) {
	config := &redisKit.RedisConfig{
		Mode: redisKit.SingleNodeMode,
		SingleNodeConfig: &redisKit.SingleNodeConfig{
			Addr: "192.168.63.42:6379",
			DB:   0,
		},
	}
	return redisKit.NewClient(config)
}

// NewSentinelClient 哨兵集群的客户端
func NewSentinelClient() (*redisKit.Client, error) {
	config := &redisKit.RedisConfig{
		Mode: redisKit.SentinelMode,
		SentinelConfig: &redisKit.SentinelConfig{
			MasterName:    "",
			SentinelAddrs: []string{"192.168.63.42:27000", "192.168.63.42:27001", "192.168.63.42:27002"},
			DB:            0,
		},
	}
	return redisKit.NewClient(config)
}
