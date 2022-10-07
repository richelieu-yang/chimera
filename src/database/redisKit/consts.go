package redisKit

type (
	// Mode Redis的集群模式
	Mode int8
)

const (
	// SingleNodeMode 单点集群
	SingleNodeMode Mode = 0
	// MasterSlaverMode 主从集群
	MasterSlaverMode Mode = 1
	// SentinelMode 哨兵集群
	SentinelMode Mode = 2
	// ClusterMode cluster集群
	ClusterMode Mode = 3
)
