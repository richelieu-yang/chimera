package redisKit

import "github.com/richelieu-yang/chimera/v2/src/compareKit"

type (
	Config struct {
		UserName string `json:"userName,optional" yaml:"userName"`
		Password string `json:"password,optional" yaml:"password"`

		Mode               Mode                `json:"mode,default=0,options=0|2|3" yaml:"mode"`
		SingleNodeConfig   *SingleNodeConfig   `json:"singleNodeConfig" yaml:"singleNodeConfig"`
		MasterSlaverConfig *MasterSlaverConfig `json:"masterSlaverConfig" yaml:"masterSlaverConfig"`
		SentinelConfig     *SentinelConfig     `json:"sentinelConfig" yaml:"sentinelConfig"`
		ClusterConfig      *ClusterConfig      `json:"clusterConfig" yaml:"clusterConfig"`
	}

	SingleNodeConfig struct {
		// Addr address(host:port)
		Addr string `json:"addr" yaml:"addr"`
		// DB Database to be selected after connecting to the server.
		DB int `json:"db,default=0" yaml:"db"`
	}

	MasterSlaverConfig struct {
	}

	SentinelConfig struct {
		// MasterName The master name.
		MasterName string `json:"masterName,default=mymaster" yaml:"masterName"`
		// SentinelAddrs A seed list of host:port addresses of sentinel nodes.
		SentinelAddrs []string `json:"sentinelAddrs" yaml:"sentinelAddrs"`
		DB            int      `json:"db,default=0" yaml:"db"`
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
	case ModeSingleNode:
		if !compareKit.Equal(config.SingleNodeConfig, config1.SingleNodeConfig) {
			return false
		}
	case ModeMasterSlaver:
		if !compareKit.Equal(config.MasterSlaverConfig, config1.MasterSlaverConfig) {
			return false
		}
	case ModeSentinel:
		if !compareKit.Equal(config.SentinelConfig, config1.SentinelConfig) {
			return false
		}
	case ModeCluster:
		if !compareKit.Equal(config.ClusterConfig, config1.ClusterConfig) {
			return false
		}
	}

	return true
}
