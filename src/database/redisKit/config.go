package redisKit

import "github.com/richelieu-yang/chimera/v2/src/compareKit"

type (
	Config struct {
		UserName string `json:"userName,optional"`
		Password string `json:"password,optional"`

		MinIdleConns int `json:"minIdleConns,default=64,range=[32:1000000]"`
		MaxIdleConns int `json:"maxIdleConns,default=256,range=[32:1000000]"`
		PoolSize     int `json:"poolSize,default=512,range=[32:1000000]"`

		Mode Mode `json:"mode,default=0,options=0|2|3"`

		SingleNodeConfig   SingleNodeConfig   `json:"singleNodeConfig"`
		MasterSlaverConfig MasterSlaverConfig `json:"masterSlaverConfig"`
		SentinelConfig     SentinelConfig     `json:"sentinelConfig"`
		ClusterConfig      ClusterConfig      `json:"clusterConfig"`
	}

	SingleNodeConfig struct {
		// Addr host:port address.
		Addr string `json:"addr"`
		// DB Database to be selected after connecting to the server.
		DB int `json:"db,default=0"`
	}

	MasterSlaverConfig struct {
	}

	SentinelConfig struct {
		// MasterName The master name.
		MasterName string `json:"masterName,default=mymaster"`
		// SentinelAddrs A seed list of host:port addresses of sentinel nodes.
		SentinelAddrs []string `json:"sentinelAddrs"`
		DB            int      `json:"db,default=0"`
	}

	ClusterConfig struct {
		// Addrs
		/*
			A seed list of host:port addresses of cluster nodes.
			可以是: 所有的 master 的地址，
			也可以是: 所有的 master + slave 的地址（推荐）.
		*/
		Addrs []string `json:"addrs"`
	}
)

func (config Config) Equal(config1 Config) bool {
	if config.UserName != config1.UserName {
		return false
	}
	if config.Password != config1.Password {
		return false
	}

	if config.MinIdleConns != config1.MinIdleConns {
		return false
	}
	if config.MaxIdleConns != config1.MaxIdleConns {
		return false
	}
	if config.PoolSize != config1.PoolSize {
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
