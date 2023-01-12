package redisKit

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

// NewDistributedMutex 生成生成Redis分布式互斥锁
/*
PS:
(1) 不可重入锁；
(2) 更多详见"Redis分布式锁（多语言）.docx".

@param name 建议以 "mutex:" 为前缀
*/
func (client *Client) NewDistributedMutex(name string, options ...redsync.Option) *redsync.Mutex {
	pool := goredis.NewPool(client.goRedisClient) // or, pool := redigo.NewPool(...)
	sync := redsync.New(pool)
	return sync.NewMutex(name, options...)
}
