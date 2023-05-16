package redisKit

import "context"

/*
Redis 列表(List): https://www.runoob.com/redis/redis-lists.html
*/

// LPush 将一个或多个值插入到列表头部
/*
语法: LPUSH key value1 [value2]

@param key 不存在的话，会自动创建此key
*/
func (client *Client) LPush(ctx context.Context, key string, values ...interface{}) (int64, error) {
	cmd := client.universalClient.LPush(ctx, key, values...)
	return cmd.Result()
}

// LPushX 将一个或多个值插入到列表头部
/*
PS:
(1) 如果 key 不存在，一个空列表会被创建并执行 LPUSH 操作；
(2) 当 key 存在但不是列表类型时，返回一个错误。

@return key不存在的话，将返回 (0, nil)
*/
func (client *Client) LPushX(ctx context.Context, key string, values ...interface{}) (int64, error) {
	cmd := client.universalClient.LPushX(ctx, key, values...)
	return cmd.Result()
}

// LPop 移除并返回列表的第一个元素
/*
PS: 如果移除后列表为空。将删除该key.

@return key不存在的情况，将返回 ("", redis.Nil)
*/
func (client *Client) LPop(ctx context.Context, key string) (string, error) {
	cmd := client.universalClient.LPop(ctx, key)
	return cmd.Result()
}
