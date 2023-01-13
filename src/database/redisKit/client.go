package redisKit

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"time"
)

type (
	Client struct {
		mode Mode

		// goRedisClient 真正的go-redis客户端
		goRedisClient redis.UniversalClient
	}
)

func (client *Client) GetMode() Mode {
	return client.mode
}

func (client *Client) GetGoRedisClient() redis.UniversalClient {
	return client.goRedisClient
}

// NewClient 新建一个go-redis客户端（内置连接池，调用方无需额外考虑并发问题）
/*
！！！：每一个命令都会重新取得一个连接，执行后立即回收，而且回收到资源池的顺序类似于堆. https://www.cnblogs.com/yangqi7/p/13289232.html

连接哨兵集群的demo: https://blog.csdn.net/supery071/article/details/109491404
*/
func NewClient(config *RedisConfig) (*Client, error) {
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
	goRedisClient := redis.NewUniversalClient(opts)

	client := &Client{
		mode:          config.Mode,
		goRedisClient: goRedisClient,
	}

	// 简单测试是否可用
	str, err := client.Ping(context.TODO())
	if err != nil {
		return nil, errorKit.Wrap(err, "fail to ping")
	}
	if str != "PONG" {
		return nil, errorKit.Simple("result(%s) of ping in invalid", str)
	}

	return client, nil
}

func newBaseOptions(userName, password string) *redis.UniversalOptions {
	return &redis.UniversalOptions{
		IdleTimeout: time.Minute * 5, // 默认5min

		MinIdleConns: 32,
		PoolSize:     128,

		Username: strKit.Trim(userName),
		Password: strKit.Trim(password),
	}
}

// 单点模式
func newSingleNodeOptions(config *RedisConfig) (*redis.UniversalOptions, error) {
	c := config.SingleNodeConfig
	if c == nil {
		return nil, errorKit.Simple("SingleNodeConfig is nil")
	}

	opts := newBaseOptions(config.UserName, config.Password)

	opts.Addrs = []string{c.Addr}
	opts.DB = c.DB

	return opts, nil
}

// 主从模式
func newMasterSlaverOptions(config *RedisConfig) (*redis.UniversalOptions, error) {
	return nil, errorKit.Simple("mode(%d) is unsupported now", config.Mode)
}

// 哨兵模式
func newSentinelOptions(config *RedisConfig) (*redis.UniversalOptions, error) {
	c := config.SentinelConfig
	if c == nil {
		return nil, errorKit.Simple("SentinelConfig is nil")
	}
	if len(c.SentinelAddrs) == 0 {
		return nil, errorKit.Simple("length of SentinelAddrs is 0")
	}

	opts := newBaseOptions(config.UserName, config.Password)

	// MasterName默认为"mymaster"
	opts.MasterName = strKit.EmptyToDefault(c.MasterName, "mymaster", true)
	opts.Addrs = c.SentinelAddrs
	opts.DB = c.DB

	return opts, nil
}

// cluster模式
func newClusterOptions(config *RedisConfig) (*redis.UniversalOptions, error) {
	c := config.ClusterConfig
	if c == nil {
		return nil, errorKit.Simple("ClusterConfig is nil")
	}

	opts := newBaseOptions(config.UserName, config.Password)

	opts.Addrs = c.Addrs

	return opts, nil
}
