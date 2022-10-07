package redisKit

type (
	RedisConfig struct {
		UserName string
		Password string

		Mode Mode

		SingleNodeConfig   *SingleNodeConfig
		MasterSlaverConfig *MasterSlaverConfig
		SentinelConfig     *SentinelConfig
		ClusterConfig      *ClusterConfig
	}

	SingleNodeConfig struct {
		// host:port address.
		Addr string
		// Database to be selected after connecting to the server.
		DB int
	}

	MasterSlaverConfig struct {
	}

	SentinelConfig struct {
		// The master name.
		MasterName string
		// A seed list of host:port addresses of sentinel nodes.
		SentinelAddrs []string
		DB            int
	}

	ClusterConfig struct {
		// A seed list of host:port addresses of cluster nodes.
		Addrs []string
	}
)

// Simplify 简化，去掉与 RedisConfig.Mode 无关的配置
func (rc *RedisConfig) Simplify() {
	if rc == nil {
		return
	}

	switch rc.Mode {
	case SingleNodeMode:
		rc.MasterSlaverConfig = nil
		rc.SentinelConfig = nil
		rc.ClusterConfig = nil
	case MasterSlaverMode:
		rc.SingleNodeConfig = nil
		rc.SentinelConfig = nil
		rc.ClusterConfig = nil
	case SentinelMode:
		rc.SingleNodeConfig = nil
		rc.MasterSlaverConfig = nil
		rc.ClusterConfig = nil
	case ClusterMode:
		rc.SingleNodeConfig = nil
		rc.MasterSlaverConfig = nil
		rc.SentinelConfig = nil
	default:
		rc.SingleNodeConfig = nil
		rc.MasterSlaverConfig = nil
		rc.SentinelConfig = nil
		rc.ClusterConfig = nil
	}
}
