package redisKit

import "github.com/richelieu-yang/chimera/v2/src/compareKit"

type (
	Config struct {
		UserName string `json:"userName" yaml:"userName"`
		Password string `json:"password" yaml:"password"`
		Prefix   string `json:"prefix" yaml:"prefix"`

		Mode         Mode                `json:"mode" yaml:"mode"`
		SingleNode   *SingleNodeConfig   `json:"singleNode" yaml:"singleNode"`
		MasterSlaver *MasterSlaverConfig `json:"masterSlaver" yaml:"masterSlaver"`
		Sentinel     *SentinelConfig     `json:"sentinel" yaml:"sentinel"`
		Cluster      *ClusterConfig      `json:"cluster" yaml:"cluster"`
	}

	SingleNodeConfig struct {
		// Addr address(host:port)
		Addr string `json:"addr" yaml:"addr"`
		// DB Database to be selected after connecting to the server.
		DB int `json:"db" yaml:"db"`
	}

	MasterSlaverConfig struct {
	}

	SentinelConfig struct {
		// MasterName The master name.
		MasterName string `json:"masterName" yaml:"masterName"`
		// SentinelAddrs A seed list of host:port addresses of sentinel nodes.
		SentinelAddrs []string `json:"sentinelAddrs" yaml:"sentinelAddrs"`
		DB            int      `json:"db" yaml:"db"`
	}

	ClusterConfig struct {
		// Addrs
		/*
			A seed list of host:port addresses of cluster nodes.
			可以是: 所有的 master 的地址，
			也可以是: 所有的 master + slave 的地址（推荐）.
		*/
		Addrs []string `json:"addrs" yaml:"addrs"`
	}
)

func (config Config) Equal(config1 Config) bool {
	if config.UserName != config1.UserName {
		return false
	}
	if config.Password != config1.Password {
		return false
	}

	if config.Mode != config1.Mode {
		return false
	}
	switch config.Mode {
	case SingleNodeMode:
		if !compareKit.Equal(config.SingleNode, config1.SingleNode) {
			return false
		}
	case MasterSlaverMode:
		if !compareKit.Equal(config.MasterSlaver, config1.MasterSlaver) {
			return false
		}
	case SentinelMode:
		if !compareKit.Equal(config.Sentinel, config1.Sentinel) {
			return false
		}
	case ClusterMode:
		if !compareKit.Equal(config.Cluster, config1.Cluster) {
			return false
		}
	}

	return true
}
