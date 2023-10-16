package redisKit

type (
	// Mode Redis的集群模式
	Mode int8
)

const (
	// ModeSingleNode 单点集群
	ModeSingleNode Mode = 0
	// ModeMasterSlaver 主从集群
	ModeMasterSlaver Mode = 1
	// ModeSentinel 哨兵集群
	ModeSentinel Mode = 2
	// ModeCluster cluster集群
	ModeCluster Mode = 3

	// DefaultMasterName 哨兵模式下
	DefaultMasterName = "mymaster"
)
