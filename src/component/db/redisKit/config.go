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

		Mode        Mode               `json:"mode" yaml:"mode" validate:"oneof=single sentinel cluster"`
		Single      *SingleConfig      `json:"single" yaml:"singleNode"`
		MasterSlave *MasterSlaveConfig `json:"masterSlave" yaml:"masterSlaver"`
		Sentinel    *SentinelConfig    `json:"sentinel" yaml:"sentinel"`
		Cluster     *ClusterConfig     `json:"cluster" yaml:"cluster"`
	}

	SingleConfig struct {
		// Addr address(host:port)
		Addr string `json:"addr" yaml:"addr" validate:"hostname_port"`

		// DB Database to be selected after connecting to the server.
		DB int `json:"db" yaml:"db" validate:"gte=0"`
	}

	MasterSlaveConfig struct {
	}

	SentinelConfig struct {
		// MasterName The master name.
		MasterName string `json:"masterName" yaml:"masterName"`

		// Addrs A seed list of host:port addresses of sentinel nodes.
		Addrs []string `json:"addrs" yaml:"addrs" validate:"required,gte=2,dive,hostname_port"`

		DB int `json:"db" yaml:"db" validate:"gte=0"`
	}

	ClusterConfig struct {
		// Addrs
		/*
			A seed list of host:port addresses of cluster nodes.
			可以是: 所有的 master 的地址，
			也可以是: 所有的 master + slave 的地址（推荐）.
		*/
		Addrs []string `json:"addrs" yaml:"addrs" validate:"required,gte=2,dive,hostname_port"`
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
		config.MasterSlave = nil
		config.Sentinel = nil
		config.Cluster = nil

		if err := v.Struct(config.Single); err != nil {
			return err
		}
	case SentinelMode:
		config.Single = nil
		config.MasterSlave = nil
		config.Cluster = nil

		if err := v.Struct(config.Sentinel); err != nil {
			return err
		}
	case ClusterMode:
		config.Single = nil
		config.MasterSlave = nil
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
		if !compareKit.Equal(config.Single, config1.Single) {
			return false
		}
	case MasterSlaverMode:
		if !compareKit.Equal(config.MasterSlave, config1.MasterSlave) {
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
