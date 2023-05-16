package redisKit

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
)

// NewDistributedMutex 生成Redis分布式互斥锁.
/*
PS:
(1) 不可重入锁;
(2) 更多详见"Redis分布式锁（多语言）.docx";
(3) 写入Redis中键的默认TTL为: 8s.
(4) 如果持有锁的时间 > Redis键的TTL，此时解锁将返回 false, error(类型: *redsync.ErrTaken; 错误内容: lock already taken, locked nodes: [0]);
(5) 全部采用默认配置的情况下，锁被别人占着，你在尝试获取锁 约4.5s 后，返回error(类型: *redsync.ErrTaken; 错误内容: lock already taken, locked nodes: [0]).
	想要修改获取锁失败后的整体等待时间，可以尝试配置 redsync.WithTries()、redsync.WithRetryDelay()、redsync.WithRetryDelayFunc().

@param name 建议以 "mutex:" 为前缀

e.g. 将TTL修改为30s（原8s）
NewDistributedMutex("name", redsync.WithExpiry(time.Second * 30))
*/
func (client *Client) NewDistributedMutex(name string, options ...redsync.Option) *redsync.Mutex {
	pool := goredis.NewPool(client.core) // or, pool := redigo.NewPool(...)
	sync := redsync.New(pool)
	return sync.NewMutex(name, options...)
}
