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
(4) 如果持有锁的时间 > Redis键的TTL，此时 "解锁" 将返回 false, error(类型: *redsync.ErrTaken; 错误内容: lock already taken, locked nodes: [0]);
(5) retry（获取锁失败后等一会，再尝试获取锁）
	相关配置: redsync.WithTries()、redsync.WithRetryDelay()、redsync.WithRetryDelayFunc()
	e.g.
		默认配置下，retry整体周期为 约4.5s（因为随机），还是获取不到则返回error(类型: *redsync.ErrTaken; 错误内容: lock already taken, locked nodes: [0]).
	e.g.1
		如果配置后，整体retry周期 > Redis中key的TTL，当上一个锁的key超时后，你就能获取到锁了（假如没有其他人也在抢锁）.

@param name 建议以 "mutex:" 为前缀

e.g. 将TTL修改为30s（原8s）
NewDistributedMutex("name", redsync.WithExpiry(time.Second * 30))
*/
func (client *Client) NewDistributedMutex(name string, options ...redsync.Option) *redsync.Mutex {
	pool := goredis.NewPool(client.core) // or, pool := redigo.NewPool(...)
	sync := redsync.New(pool)
	return sync.NewMutex(name, options...)
}
