package redisKit

import (
	"context"
	"time"
)

// TTL 返回给定 key 的剩余生存时间
/*
@return
(1) key不存在，返回值: 	-2ns（即-2）, nil（能直接通过 == 或 switch 进行比较）
(2) key为持久化键，返回值: -1ns（即-1）, nil（能直接通过 == 或 switch 进行比较）
*/
func (client Client) TTL(ctx context.Context, key string) (time.Duration, error) {
	return client.UC.TTL(ctx, key).Result()
}

// Expire
/*
语法：EXPIRE key seconds
说明：为给定 key 设置过期时间，以秒计。
*/
func (client Client) Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return client.UC.Expire(ctx, key, expiration).Result()
}

// ExpireAt
/*
语法：EXPIREAT key timestamp
说明：EXPIREAT 的作用和 EXPIRE 类似，都用于为 key 设置过期时间。 不同在于 EXPIREAT 命令接受的时间参数是 UNIX 时间戳(unix timestamp)。
*/
func (client Client) ExpireAt(ctx context.Context, key string, tm time.Time) (bool, error) {
	return client.UC.ExpireAt(ctx, key, tm).Result()
}
