package redisKit

type (
	// Mode Redis的集群模式
	Mode string
)

const (
	// ModeSingle 单点
	ModeSingle Mode = "single"
	// ModeMasterSlave 主从集群
	ModeMasterSlave Mode = "masterSlave"
	// ModeSentinel 哨兵集群
	ModeSentinel Mode = "sentinel"
	// ModeCluster cluster集群
	ModeCluster Mode = "cluster"

	// DefaultMasterName 哨兵模式下，默认的MasterName
	DefaultMasterName = "mymaster"
)
