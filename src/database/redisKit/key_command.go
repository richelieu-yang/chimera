package redisKit

import (
	"context"
	"time"
)

// Type 返回 key 所储存的值的类型.
/*
语法:
	TYPE KEY_NAME
返回值:
	none (key不存在)
	string (字符串)
	list (列表)
	set (集合)
	zset (有序集)
	hash (哈希表)

e.g.
传参key不存在的情况 => ("none", nil)
*/
func (client *Client) Type(ctx context.Context, key string) (string, error) {
	return client.goRedisClient.Type(ctx, key).Result()
}

// Keys
/*
Deprecated: 禁止在生产环境使用Keys正则匹配操作（实际即便是开发、测试环境也要慎重使用）！！！

e.g.
db为空（或者不存在与 传参match 响应的key） => ([]string{}, nil)（第一个返回值不为nil）
*/
func (client *Client) Keys(ctx context.Context, match string) ([]string, error) {
	return client.goRedisClient.Keys(ctx, match).Result()
}

// TTL 返回 key 的剩余过期时间.
/*
语法：
	TTL KEY_NAME
返回值：
	当 key 不存在时，返回 -2 。 当 key 存在但没有设置剩余生存时间时，返回 -1 。 否则，以毫秒为单位，返回 key 的剩余生存时间。

e.g. key不存在
	duration, err := client.TTL(context.TODO(), "a")
	if err != nil {
		panic(err)
	}
	fmt.Println(duration)       // -2ns
	fmt.Println(duration == -2) // true

e.g.1 key为持久化键
	duration, err := client.TTL(context.TODO(), "a")
	if err != nil {
		panic(err)
	}
	fmt.Println(duration)       // -1ns
	fmt.Println(duration == -1) // true
*/
func (client *Client) TTL(ctx context.Context, key string) (time.Duration, error) {
	return client.goRedisClient.TTL(ctx, key).Result()
}

// Expire
/*
语法：EXPIRE key seconds
说明：为给定 key 设置过期时间，以秒计。

e.g.
key不存在	=> (false, nil)
key存在		=> (true, nil)
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

// Scan 迭代当前数据库中的数据库键.
/*
PS:
(1)	scan命令也并不是完美的，它"返回的结果有可能重复"，因此需要客户端"去重"；
(2) 用于替代keys，因为keys在大数据量有性能问题；
(3) 返回的[]string实例的长度可能会大于传参count，比如瞎传cursor的情况，编码时得注意.

@return 返回的error == nil的情况下，第1个返回值([]string)必定不为nil

e.g. db为空（|| db中不存在符合条件的key）
(context.TODO(), 0, "*", 10) => ([]string{}, 0, nil)
*/
func (client *Client) Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return client.goRedisClient.Scan(ctx, cursor, match, count).Result()
}
