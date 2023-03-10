package redisKit

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
)

// NewDistributedMutex 生成Redis分布式互斥锁.
/*
PS:
(1) 不可重入锁；
(2) 更多详见"Redis分布式锁（多语言）.docx".
(3) 写入Redis中的键，默认TTL为8s，可以通过 redsync.WithExpiry() 返回的 redsync.Option实例 来修改.

@param name 建议以 "mutex:" 为前缀
*/
func (client *Client) NewDistributedMutex(name string, options ...redsync.Option) *redsync.Mutex {
	pool := goredis.NewPool(client.goRedisClient) // or, pool := redigo.NewPool(...)
	sync := redsync.New(pool)
	return sync.NewMutex(name, options...)
}
