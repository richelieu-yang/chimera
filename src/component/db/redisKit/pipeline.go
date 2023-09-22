package redisKit

import (
	"github.com/redis/go-redis/v9"
)

// Pipeline 管道
/*
TODO: cluster集群模式下，使用pipeline有问题.

redis.Pipeliner实例调用Exec函数执行时，第二个返回值为error类型：
(1) 假如管道中每条命令都正常执行，将返回nil；
(2) 假如管道中某些命令报错了，将返回第一个报错命令的error实例.
*/
func (client *Client) Pipeline() redis.Pipeliner {
	return client.universalClient.Pipeline()
}

// TxPipeline 事务管道
/*
TODO: cluster集群模式下，使用pipeline有问题.

详见"Redis.docx".
*/
func (client *Client) TxPipeline() redis.Pipeliner {
	return client.universalClient.TxPipeline()
}
