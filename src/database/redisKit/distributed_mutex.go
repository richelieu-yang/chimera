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
(4) Unlock:
	如果持有锁的时间 > Redis键的TTL，此时 "解锁" 将返回 false, error(类型: *redsync.ErrTaken; 错误内容: lock already taken, locked nodes: [0]);
(5) Lock:
	retry（获取锁失败后等一会，再尝试获取锁），相关配置: redsync.WithTries()、redsync.WithRetryDelay()、redsync.WithRetryDelayFunc()
	retry整体周期内，如果获取锁失败，将返回error:
		(a) 类型: *redsync.ErrTaken; 错误内容: lock already taken, locked nodes: [0]
		(b) redsync.ErrFailed
	e.g.
		PS: 尝试获取锁其实也要花一点时间.
		默认配置下，总计重试 32 次，第0次不等待，第1~31次在开始时等待随机时间（[50, 250)ms）.
		retry整体周期为 约4.8s（因为随机），还是获取不到则返回error.
	e.g.1
		如果配置后，整体retry周期 > Redis中key的TTL，当上一个锁的key超时后，你就能获取到锁了（假如没有其他人也在抢锁）.

@param name 建议以 "mutex:" 为前缀

e.g. 将TTL修改为30s（原8s）
NewDistributedMutex("name", redsync.WithExpiry(time.Second * 30))
*/
func (client *Client) NewDistributedMutex(name string, options ...redsync.Option) *redsync.Mutex {
	pool := goredis.NewPool(client.universalClient) // or, pool := redigo.NewPool(...)
	sync := redsync.New(pool)
	return sync.NewMutex(name, options...)
}
