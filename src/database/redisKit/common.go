package redisKit

import (
	"context"
)

// Del （删）key 存在时，删除 key
/*
@return 第一个返回值代表：是否删除成功

e.g.
如果key不存在，将返回: (false, nil)
*/
func (client *Client) Del(ctx context.Context, key string) (bool, error) {
	reply, err := client.goRedisClient.Del(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return reply == 1, nil
}

// Exists 检查给定 key 是否存在
/*
PS: 如果传参的key有多个，只要其中有一个key存在，就返回true（不报错的情况下）.
*/
func (client *Client) Exists(ctx context.Context, keys ...string) (bool, error) {
	reply, err := client.goRedisClient.Exists(ctx, keys...).Result()
	if err != nil {
		return false, err
	}
	return reply == 1, nil
}

// Publish 发布
/*
e.g.
("", "") => nil
*/
func (client *Client) Publish(ctx context.Context, channel string, message interface{}) error {
	_, err := client.goRedisClient.Publish(ctx, channel, message).Result()
	return err
}

// FlushDB 清空当前数据库中的所有 key
/*
慎用！！！
*/
func (client *Client) FlushDB(ctx context.Context) (string, error) {
	return client.goRedisClient.FlushDB(ctx).Result()
}

// FlushAll 清空整个 Redis 服务器的数据(删除所有数据库的所有 key )
/*
慎用！！！
*/
func (client *Client) FlushAll(ctx context.Context) (string, error) {
	return client.goRedisClient.FlushAll(ctx).Result()
}
