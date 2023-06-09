package redisKit

type (
	Config struct {
		UserName string `json:"userName,optional"`
		Password string `json:"password,optional"`

		MinIdleConns int `json:"minIdleConns,default=64,range=[32:100000]"`
		MaxIdleConns int `json:"maxIdleConns,default=128,range=[32:100000]"`
		PoolSize     int `json:"poolSize,default=512,range=[32:100000]"`

		Mode Mode `json:"mode,default=0,options=0|2|3"`

		SingleNodeConfig   *SingleNodeConfig   `json:"singleNodeConfig"`
		MasterSlaverConfig *MasterSlaverConfig `json:"masterSlaverConfig"`
		SentinelConfig     *SentinelConfig     `json:"sentinelConfig"`
		ClusterConfig      *ClusterConfig      `json:"clusterConfig"`
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

// Simplify 简化，去掉与 Config.Mode 无关的配置
func (rc *Config) Simplify() {
	if rc == nil {
		return
	}

	switch rc.Mode {
	case ModeSingleNode:
		rc.MasterSlaverConfig = nil
		rc.SentinelConfig = nil
		rc.ClusterConfig = nil
	case ModeMasterSlaver:
		rc.SingleNodeConfig = nil
		rc.SentinelConfig = nil
		rc.ClusterConfig = nil
	case ModeSentinel:
		rc.SingleNodeConfig = nil
		rc.MasterSlaverConfig = nil
		rc.ClusterConfig = nil
	case ModeCluster:
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
