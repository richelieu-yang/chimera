package redisKit

type (
	// Mode Redis的集群模式
	Mode string
)

const (
	// SingleNodeMode 单点
	SingleNodeMode Mode = "singleNode"
	// MasterSlaverMode 主从集群
	MasterSlaverMode Mode = "masterSlaver"
	// SentinelMode 哨兵集群
	SentinelMode Mode = "sentinel"
	// ClusterMode cluster集群
	ClusterMode Mode = "cluster"

	// DefaultMasterName 哨兵模式，默认的MasterName
	DefaultMasterName = "mymaster"
)
