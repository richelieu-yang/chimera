package redisKit

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"time"
)

type (
	Client struct {
		mode Mode

		// universalClient go-redis的客户端
		universalClient redis.UniversalClient
	}
)

func (client *Client) Close() error {
	if client != nil && client.universalClient != nil {
		return client.universalClient.Close()
	}
	return nil
}

func (client *Client) GetMode() Mode {
	return client.mode
}

// GetUniversalClient 返回go-redis客户端
func (client *Client) GetUniversalClient() redis.UniversalClient {
	return client.universalClient
}

// NewClient 新建一个go-redis客户端（内置连接池，调用方无需额外考虑并发问题）
/*
!!!: 每一个命令都会重新取得一个连接，执行后立即回收，而且回收到资源池的顺序类似于堆. https://www.cnblogs.com/yangqi7/p/13289232.html

连接哨兵集群的demo: https://blog.csdn.net/supery071/article/details/109491404

@return	(1) 两个返回值，必定有一个为nil，另一个非nil；
		(2) Cluster模式下，第1个返回值的类型: *redis.ClusterClient.
*/
func NewClient(config Config) (client *Client, err error) {
	var opts *redis.UniversalOptions
	switch config.Mode {
	case ModeSingleNode:
		opts, err = newSingleNodeOptions(config)
	case ModeMasterSlaver:
		opts, err = newMasterSlaverOptions(config)
	case ModeSentinel:
		opts, err = newSentinelOptions(config)
	case ModeCluster:
		opts, err = newClusterOptions(config)
	default:
		err = errorKit.New("mode(%d) is invalid", config.Mode)
	}
	if err != nil {
		return
	}

	//opts.OnConnect = func(ctx context.Context, conn *redis.Conn) error {
	//	logrus.Infof("conn: %v", conn)
	//	return nil
	//}

	goRedisClient := redis.NewUniversalClient(opts)
	client = &Client{
		mode:            config.Mode,
		universalClient: goRedisClient,
	}
	defer func() {
		if err != nil {
			_ = client.Close()
			client = nil
		}
	}()

	// 简单测试是否Redis服务可用
	pingTimeout := time.Second * 3
	ctx, cancel := context.WithTimeout(context.TODO(), pingTimeout)
	defer cancel()
	str, err := client.Ping(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			err = errorKit.New("initial ping timeout(%s)", pingTimeout)
		} else {
			err = errorKit.Wrap(err, "initial ping fails")
		}
		return
	}
	if !strKit.EqualsIgnoreCase(str, "PONG") {
		err = errorKit.New("result(%s) of initial ping is invalid", str)
		return
	}
	return
}

func newBaseOptions(config Config) *redis.UniversalOptions {
	return &redis.UniversalOptions{
		Username: config.UserName,
		Password: config.Password,
	}
}

// newSingleNodeOptions 单点模式
func newSingleNodeOptions(config Config) (*redis.UniversalOptions, error) {
	c := config.SingleNodeConfig
	if c == nil {
		return nil, errorKit.New("SingleNodeConfig is nil")
	}

	opts := newBaseOptions(config)
	opts.Addrs = []string{c.Addr}
	opts.DB = c.DB
	return opts, nil
}

// newMasterSlaverOptions 主从模式
func newMasterSlaverOptions(config Config) (*redis.UniversalOptions, error) {
	return nil, errorKit.New("mode(%d) is unsupported now", config.Mode)
}

// newSentinelOptions 哨兵模式
func newSentinelOptions(config Config) (*redis.UniversalOptions, error) {
	c := config.SentinelConfig
	if c == nil {
		return nil, errorKit.New("SentinelConfig is nil")
	}
	if len(c.SentinelAddrs) == 0 {
		return nil, errorKit.New("length of SentinelAddrs is 0")
	}

	opts := newBaseOptions(config)
	// MasterName默认为"mymaster"
	opts.MasterName = strKit.EmptyToDefault(c.MasterName, "mymaster", true)
	opts.Addrs = c.SentinelAddrs
	opts.DB = c.DB
	return opts, nil
}

// newClusterOptions cluster模式
func newClusterOptions(config Config) (*redis.UniversalOptions, error) {
	c := config.ClusterConfig
	if c == nil {
		return nil, errorKit.New("ClusterConfig is nil")
	}

	opts := newBaseOptions(config)
	opts.Addrs = c.Addrs
	return opts, nil
}
