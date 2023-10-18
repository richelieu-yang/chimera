package redisKit

type (
	// Mode Redis的集群模式
	Mode string
)

const (
	// SingleMode 单点
	SingleMode Mode = "single"
	// MasterSlaveMode 主从集群
	MasterSlaveMode Mode = "masterSlave"
	// SentinelMode 哨兵集群
	SentinelMode Mode = "sentinel"
	// ClusterMode cluster集群
	ClusterMode Mode = "cluster"

	// DefaultMasterName 哨兵模式下，默认的MasterName
	DefaultMasterName = "mymaster"
)
