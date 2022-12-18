package redisKit

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/richelieu42/go-scales/src/core/timeKit"
	"github.com/richelieu42/go-scales/src/idKit"

	"time"
)

type (
	Client struct {
		Mode Mode
		/* 真正的go-redis客户端 */
		UC redis.UniversalClient
		/* 用于生成Redis分布式互斥锁 */
		rs *redsync.Redsync
	}
)

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
	uc := redis.NewUniversalClient(opts)
	pool := goredis.NewPool(uc) // or, pool := redigo.NewPool(...)
	rs := redsync.New(pool)
	client := &Client{
		Mode: config.Mode,
		UC:   uc,
		rs:   rs,
	}

	if err := testConnection(client); err != nil {
		return nil, err
	}
	return client, nil
}

// testConnection 进行简单的测试，以防Redis服务部署的有问题（方便甩锅）
func testConnection(client *Client) error {
	if _, err := client.Ping(); err != nil {
		return err
	}

	key := fmt.Sprintf("go-scales:test:redis:%s", idKit.NewSimpleUUID())
	value := timeKit.FormatCurrentTime()
	expiration := time.Second * 8 // 8s后过期（PS: 如果要debug此函数的话，建议将时间延长）

	// set
	ok, err := client.Set(key, value, expiration)
	if err != nil {
		return err
	}
	if !ok {
		return errorKit.Simple("fail to set item(key: %s, value: %s)", key, value)
	}

	// get
	value1, err := client.Get(key)
	if err != nil {
		if err == redis.Nil {
			return errorKit.Simple("key(%s) doesn't exist", key)
		}
		return err
	}
	if value != value1 {
		return errorKit.Simple("value(%s) != value1(%s)", value, value1)
	}

	// del && get
	ok, err = client.Del(key)
	if err != nil {
		return err
	}
	if !ok {
		return errorKit.Simple("fail to delete item(key: %s)", key)
	}

	return nil
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
