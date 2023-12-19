package redisKit

import (
	"github.com/richelieu-yang/chimera/v2/src/compareKit"
)

type (
	Config struct {
		UserName  string `json:"userName" yaml:"userName"`
		Password  string `json:"password" yaml:"password"`
		KeyPrefix string `json:"keyPrefix" yaml:"keyPrefix"`

		Mode Mode `json:"mode" yaml:"mode" validate:"oneof=single sentinel cluster"`

		Single *SingleConfig `json:"single" yaml:"singleNode" validate:"required_if=Mode single"`
		//MasterSlave *MasterSlaveConfig `json:"masterSlave" yaml:"masterSlave" validate:"required_if=Mode "`
		Sentinel *SentinelConfig `json:"sentinel" yaml:"sentinel" validate:"required_if=Mode sentinel"`
		Cluster  *ClusterConfig  `json:"cluster" yaml:"cluster" validate:"required_if=Mode cluster"`
	}

	SingleConfig struct {
		// Addr address(host:port)
		Addr string `json:"addr" yaml:"addr" validate:"hostname_port"`

		// DB Database to be selected after connecting to the server.
		DB int `json:"db" yaml:"db" validate:"gte=0"`
	}

	//MasterSlaveConfig struct {
	//}

	SentinelConfig struct {
		// MasterName The master name.
		MasterName string `json:"masterName" yaml:"masterName"`

		// Addrs A seed list of host:port addresses of sentinel nodes.
		Addrs []string `json:"addrs" yaml:"addrs" validate:"required,gte=2,unique,dive,hostname_port"`

		DB int `json:"db" yaml:"db" validate:"gte=0"`
	}

	ClusterConfig struct {
		// Addrs
		/*
			A seed list of host:port addresses of cluster nodes.
			可以是: 所有的 master 的地址，
			也可以是: 所有的 master + slave 的地址（推荐）.
		*/
		Addrs []string `json:"addrs" yaml:"addrs" validate:"required,gte=2,unique,dive,hostname_port"`
	}
)

// Simplify 简化配置.
func (config *Config) Simplify() {
	if config == nil {
		return
	}

	switch config.Mode {
	case ModeSingle:
		config.Sentinel = nil
		config.Cluster = nil
	case ModeSentinel:
		config.Single = nil
		config.Cluster = nil
	case ModeCluster:
		config.Single = nil
		config.Sentinel = nil
	case ModeMasterSlave:
		fallthrough
	default:
		// do nothing
	}
}

func (config Config) Equal(config1 Config) bool {
	config.Simplify()
	config1.Simplify()
	return compareKit.Equal(config, config1)

	//if config.UserName != config1.UserName {
	//	return false
	//}
	//if config.Password != config1.Password {
	//	return false
	//}
	//
	//if config.Mode != config1.Mode {
	//	return false
	//}
	//switch config.Mode {
	//case ModeSingle:
	//	if !compareKit.Equal(config.Single, config1.Single) {
	//		return false
	//	}
	//case ModeMasterSlave:
	//	if !compareKit.Equal(config.MasterSlave, config1.MasterSlave) {
	//		return false
	//	}
	//case ModeSentinel:
	//	if !compareKit.Equal(config.Sentinel, config1.Sentinel) {
	//		return false
	//	}
	//case ModeCluster:
	//	if !compareKit.Equal(config.Cluster, config1.Cluster) {
	//		return false
	//	}
	//}
	//
	//return true
}
