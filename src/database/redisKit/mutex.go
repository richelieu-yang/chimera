package redisKit

import "github.com/go-redsync/redsync/v4"

// NewMutex 生成生成Redis分布式互斥锁
/*
PS:
(1) 不可重入锁；
(2) 更多详见"Redis分布式锁（多语言）.docx".

@param name 建议以 "mutex:" 为前缀
*/
func (client *Client) NewMutex(name string, options ...redsync.Option) *redsync.Mutex {
	return client.sync.NewMutex(name, options...)
}
