package redisKit

import (
	"context"
	"time"
)

// Expire
/*
语法：EXPIRE key seconds
说明：为给定 key 设置过期时间，以秒计。
*/
func (client *Client) Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return client.goRedisClient.Expire(ctx, key, expiration).Result()
}

// ExpireAt
/*
语法：EXPIREAT key timestamp
说明：EXPIREAT 的作用和 EXPIRE 类似，都用于为 key 设置过期时间。 不同在于 EXPIREAT 命令接受的时间参数是 UNIX 时间戳(unix timestamp)。
*/
func (client *Client) ExpireAt(ctx context.Context, key string, tm time.Time) (bool, error) {
	return client.goRedisClient.ExpireAt(ctx, key, tm).Result()
}
