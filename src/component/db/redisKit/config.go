package redisKit

import (
	"github.com/richelieu-yang/chimera/v2/src/compareKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

type (
	Config struct {
		UserName string `json:"userName" yaml:"userName"`
		Password string `json:"password" yaml:"password"`
		Prefix   string `json:"prefix" yaml:"prefix"`

		Mode         Mode                `json:"mode" yaml:"mode" validate:"oneof=singleNode sentinel cluster"`
		SingleNode   *SingleNodeConfig   `json:"singleNode" yaml:"singleNode"`
		MasterSlaver *MasterSlaverConfig `json:"masterSlaver" yaml:"masterSlaver"`
		Sentinel     *SentinelConfig     `json:"sentinel" yaml:"sentinel"`
		Cluster      *ClusterConfig      `json:"cluster" yaml:"cluster"`
	}

	SingleNodeConfig struct {
		// Addr address(host:port)
		Addr string `json:"addr" yaml:"addr" validate:"hostname_port"`
		// DB Database to be selected after connecting to the server.
		DB int `json:"db" yaml:"db" validate:"gte=0"`
	}

	MasterSlaverConfig struct {
	}

	SentinelConfig struct {
		// MasterName The master name.
		MasterName string `json:"masterName" yaml:"masterName"`
		// Addrs A seed list of host:port addresses of sentinel nodes.
		Addrs []string `json:"addrs" yaml:"addrs" validate:"required"`
		DB    int      `json:"db" yaml:"db" validate:"gte=0"`
	}

	ClusterConfig struct {
		// Addrs
		/*
			A seed list of host:port addresses of cluster nodes.
			可以是: 所有的 master 的地址，
			也可以是: 所有的 master + slave 的地址（推荐）.
		*/
		Addrs []string `json:"addrs" yaml:"addrs" validate:"required"`
	}
)

// Validate
/*
PS: config可能为nil.
*/
func (config *Config) Validate() error {
	v := validateKit.New()
	if err := v.Struct(config); err != nil {
		return errorKit.Wrap(err, "Fail to validate")
	}

	switch config.Mode {
	case SingleNodeMode:
		config.MasterSlaver = nil
		config.Sentinel = nil
		config.Cluster = nil

		if err := v.Struct(config.SingleNode); err != nil {
			return err
		}
	case SentinelMode:
		config.SingleNode = nil
		config.MasterSlaver = nil
		config.Cluster = nil

		if err := v.Struct(config.Sentinel); err != nil {
			return err
		}
	case ClusterMode:
		config.SingleNode = nil
		config.MasterSlaver = nil
		config.Sentinel = nil

		if err := v.Struct(config.Cluster); err != nil {
			return err
		}
	case MasterSlaverMode:
		fallthrough
	default:
		return errorKit.New("invalid mode(%s)", config.Mode)
	}

	return nil
}

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
