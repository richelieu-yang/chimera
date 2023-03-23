package redisKit

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/richelieu42/chimera/src/core/errorKit"
	"github.com/richelieu42/chimera/src/core/strKit"
)

type (
	Client struct {
		mode Mode

		// core go-redis客户端
		core redis.UniversalClient
	}
)

func (client *Client) GetMode() Mode {
	return client.mode
}

// GetGoRedisClient 返回go-redis客户端
func (client *Client) GetGoRedisClient() redis.UniversalClient {
	return client.core
}

// NewClient 新建一个go-redis客户端（内置连接池，调用方无需额外考虑并发问题）
/*
!!!: 每一个命令都会重新取得一个连接，执行后立即回收，而且回收到资源池的顺序类似于堆. https://www.cnblogs.com/yangqi7/p/13289232.html

连接哨兵集群的demo: https://blog.csdn.net/supery071/article/details/109491404

@return cluster模式下，第1个返回值的类型: *redis.ClusterClient
*/
func NewClient(config *Config) (*Client, error) {
	if config == nil {
		return nil, errorKit.Simple("config is nil")
	}

	var opts *redis.UniversalOptions
	var err error
	switch config.Mode {
	case SingleNodeMode:
		opts, err = newSingleNodeOptions(config)
	case MasterSlaverMode:
		opts, err = newMasterSlaverOptions(config)
	case SentinelMode:
		opts, err = newSentinelOptions(config)
	case ClusterMode:
		opts, err = newClusterOptions(config)
	default:
		err = errorKit.Simple("mode(%d) is invalid", config.Mode)
	}
	if err != nil {
		return nil, err
	}

	// test
	//opts.OnConnect = func(ctx context.Context, conn *redis.Conn) error {
	//	logrus.Infof("conn: %v", conn)
	//	return nil
	//}

	goRedisClient := redis.NewUniversalClient(opts)
	client := &Client{
		mode: config.Mode,
		core: goRedisClient,
	}

	// 简单测试是否Redis服务可用
	str, err := client.Ping(context.TODO())
	if err != nil {
		return nil, errorKit.Wrap(err, "fail to ping")
	}
	if str != "PONG" {
		return nil, errorKit.Simple("result(%s) of ping in invalid", str)
	}

	return client, nil
}

func newBaseOptions(config *Config) *redis.UniversalOptions {
	return &redis.UniversalOptions{
		MinIdleConns: 32,
		PoolSize:     128,

		Username: config.UserName,
		Password: config.Password,
	}
}

// newSingleNodeOptions 单点模式
func newSingleNodeOptions(config *Config) (*redis.UniversalOptions, error) {
	c := config.SingleNodeConfig
	if c == nil {
		return nil, errorKit.Simple("SingleNodeConfig is nil")
	}

	opts := newBaseOptions(config)

	opts.Addrs = []string{c.Addr}
	opts.DB = c.DB

	return opts, nil
}

// newMasterSlaverOptions 主从模式
func newMasterSlaverOptions(config *Config) (*redis.UniversalOptions, error) {
	return nil, errorKit.Simple("mode(%d) is unsupported now", config.Mode)
}

// newSentinelOptions 哨兵模式
func newSentinelOptions(config *Config) (*redis.UniversalOptions, error) {
	c := config.SentinelConfig
	if c == nil {
		return nil, errorKit.Simple("SentinelConfig is nil")
	}
	if len(c.SentinelAddrs) == 0 {
		return nil, errorKit.Simple("length of SentinelAddrs is 0")
	}

	opts := newBaseOptions(config)

	// MasterName默认为"mymaster"
	opts.MasterName = strKit.EmptyToDefault(c.MasterName, "mymaster", true)
	opts.Addrs = c.SentinelAddrs
	opts.DB = c.DB

	return opts, nil
}

// newClusterOptions cluster模式
func newClusterOptions(config *Config) (*redis.UniversalOptions, error) {
	c := config.ClusterConfig
	if c == nil {
		return nil, errorKit.Simple("ClusterConfig is nil")
	}

	opts := newBaseOptions(config)

	opts.Addrs = c.Addrs

	return opts, nil
}
