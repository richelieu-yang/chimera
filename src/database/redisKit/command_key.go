package redisKit

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
	"sync"
	"time"
)

// Type
/*
命令说明: 	返回 key 所储存的值的类型.
命令语法:	TYPE KEY_NAME
命令返回值:
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

// Exists
/*
命令说明:	检查给定 key 是否存在.
命令语法:	EXISTS KEY_NAME
命令返回值:	若 key 存在返回 1 ，否则返回 0.
*/
func (client *Client) Exists(ctx context.Context, keys ...string) (bool, error) {
	intCmd := client.goRedisClient.Exists(ctx, keys...)
	i, err := intCmd.Result()
	if err != nil {
		return false, err
	}
	return i == 1, nil
}

// Del
/*
命令说明:	删除已存在的键。不存在的 key 会被忽略。
命令语法:	DEL KEY_NAME
命令返回值:	被删除 key 的数量。
*/
func (client *Client) Del(ctx context.Context, keys ...string) (int64, error) {
	intCmd := client.goRedisClient.Del(ctx, keys...)
	return intCmd.Result()
}

// TTL
/*
命令说明:	返回 key 的剩余过期时间.
命令语法：	TTL KEY_NAME
命令返回值：	当 key 不存在时，返回 -2 。 当 key 存在但没有设置剩余生存时间时，返回 -1 。 否则，以毫秒为单位，返回 key 的剩余生存时间.

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
命令说明:	设置 key 的过期时间，key 过期后将不再可用。单位以秒计。
命令语法:	Expire KEY_NAME TIME_IN_SECONDS
命令返回值:	设置成功返回 1 。 当 key 不存在或者不能为 key 设置过期时间时(比如在低于 2.1.3 版本的 Redis 中你尝试更新 key 的过期时间)返回 0 。

e.g.
key不存在	=> (false, nil)
key存在		=> (true, nil)
*/
func (client *Client) Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	boolCmd := client.goRedisClient.Expire(ctx, key, expiration)
	return boolCmd.Result()
}

// ExpireAt
/*
命令说明:	以 UNIX 时间戳(unix timestamp)格式设置 key 的过期时间。key 过期后将不再可用。
命令语法:	Expireat KEY_NAME TIME_IN_UNIX_TIMESTAMP
命令返回值:	设置成功返回 1 。 当 key 不存在或者不能为 key 设置过期时间时(比如在低于 2.1.3 版本的 Redis 中你尝试更新 key 的过期时间)返回 0 。
*/
func (client *Client) ExpireAt(ctx context.Context, key string, tm time.Time) (bool, error) {
	boolCmd := client.goRedisClient.ExpireAt(ctx, key, tm)
	return boolCmd.Result()
}

// Persist
/*
命令说明:	移除给定 key 的过期时间，使得 key 永不过期。
命令语法:	PERSIST KEY_NAME
命令返回值:	当过期时间移除成功时，返回 1 。 如果 key 不存在或 key 没有设置过期时间，返回 0 。
*/
func (client *Client) Persist(ctx context.Context, key string) (bool, error) {
	boolCmd := client.goRedisClient.Persist(ctx, key)
	return boolCmd.Result()
}

// Keys
/*
Deprecated: 当 KEYS命令 被用于处理一个大的数据库时，它可能会阻塞服务器达数秒之久。

e.g.
db为空（或者不存在与 传参match 响应的key） => ([]string{}, nil)（第一个返回值不为nil）
*/
func (client *Client) Keys(ctx context.Context, pattern string) ([]string, error) {
	stringSliceCmd := client.goRedisClient.Keys(ctx, pattern)
	return stringSliceCmd.Result()
}

// Scan 迭代当前数据库中的数据库键.
/*
Deprecated: 建议直接使用 ScanFully.

PS:
(1)	scan命令也并不是完美的，它"返回的结果有可能重复"，因此需要客户端"去重"；
(2) 用于替代keys，因为keys在大数据量有性能问题；
(3) 返回的[]string实例的长度可能会大于传参count，比如瞎传cursor的情况，编码时得注意.

@return 返回的error == nil的情况下，第1个返回值([]string)必定不为nil

e.g. db为空（|| db中不存在符合条件的key）
(context.TODO(), 0, "*", 10) => ([]string{}, 0, nil)
*/
func (client *Client) Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error) {
	scanCmd := client.goRedisClient.Scan(ctx, cursor, match, count)
	return scanCmd.Result()
}

// ScanFully 对 Scan 进行了封装，用于替代 Keys 命令.
/*
增量迭代命令的缺点:
	因为在对键进行增量式迭代的过程中， 键可能会被修改， 所以增量式迭代命令只能对被返回的元素"提供有限的保证" （offer limited guarantees about the returned elements）。

PS:
(1) 如果db为空，将返回: [] <nil>
(2) redis cluster模式下，需要特殊处理（详见代码），否则：明明有数据的情况下，可能取不到数据，或者取到的数据不全（因为只找1个节点要）.

@return 返回的error == nil的情况下，第1个返回值([]string)必定不为nil

e.g. db为空（|| db中不存在符合条件的key）
(context.TODO(), "*", 10) => ([]string{}, nil)
*/
func (client *Client) ScanFully(ctx context.Context, match string, count int64) ([]string, error) {
	// 抽出来的通用代码，作用: 让1个节点（node）执行命令
	// 方法1: 自行通过重复调用scan实现
	f := func(ctx context.Context, client redis.UniversalClient) ([]string, error) {
		var cursor uint64 = 0
		var keys []string

		for {
			var s []string
			var err error
			s, cursor, err = client.Scan(ctx, cursor, match, count).Result()
			if err != nil {
				return nil, err
			}

			keys = sliceKit.Merge(keys, s)

			if cursor == 0 {
				// 已经完整的过一遍了，中断循环
				break
			}
		}
		return sliceKit.RemoveDuplicate(keys), nil
	}

	// 方法2:
	//f := func(ctx context.Context, client redis.UniversalClient) ([]string, error) {
	//	var cursor uint64 = 0
	//	var keys []string
	//
	//	scanCmd := client.Scan(ctx, cursor, match, count)
	//	iter := scanCmd.Iterator()
	//	for iter.Next(ctx) {
	//		keys = append(keys, iter.Val())
	//	}
	//	if err := iter.Err(); err != nil {
	//		return nil, err
	//	}
	//
	//	if keys == nil {
	//		keys = []string{}
	//	} else {
	//		keys = sliceKit.RemoveDuplicate(keys)
	//	}
	//	return keys, nil
	//}

	clusterClient, ok := client.goRedisClient.(*redis.ClusterClient)
	if ok {
		// (1) cluster集群，特殊处理（还是有小概率漏数据）
		var keys []string
		var lock = new(sync.Mutex)

		err := clusterClient.ForEachMaster(ctx, func(ctx context.Context, client *redis.Client) error {
			s, err := f(ctx, client)
			if err != nil {
				return err
			}

			// !!!: 对于ForEachMaster()，会起多个goroutine"同时"执行传入的函数，此处加锁是为了处理并发问题，以防: 同时修改变量keys，导致最终的keys漏东西
			lock.Lock()
			defer lock.Unlock()

			keys = sliceKit.Merge(keys, s)
			return nil
		})
		if err != nil {
			return nil, err
		}
		return sliceKit.RemoveDuplicate(keys), nil
	}
	// (2) 非cluster集群，常规处理
	return f(ctx, client.goRedisClient)
}

// RandomKey
/*
命令说明:	从当前数据库中随机返回一个 key.
命令语法:	RANDOMKEY
命令返回值:	当数据库不为空时，返回一个 key；当数据库为空时，返回 nil（windows 系统返回 null）.
*/
func (client *Client) RandomKey(ctx context.Context) (string, error) {
	stringCmd := client.goRedisClient.RandomKey(ctx)
	return stringCmd.Result()
}

// Rename
/*
命令说明:	修改 key 的名称 。
命令语法:	RENAME OLD_KEY_NAME NEW_KEY_NAME
命令返回值:
	改名成功时提示 OK ，失败时候返回一个错误。
	当 OLD_KEY_NAME 和 NEW_KEY_NAME 相同，或者 OLD_KEY_NAME 不存在时，返回一个错误。 当 NEW_KEY_NAME 已经存在时， RENAME 命令将覆盖旧值。
*/
func (client *Client) Rename(ctx context.Context, key, newKey string) (string, error) {
	statusCmd := client.goRedisClient.Rename(ctx, key, newKey)
	return statusCmd.Result()
}

// RenameNX
/*
命令说明:	在新的 key 不存在时修改 key 的名称 。
命令语法:	RENAMENX OLD_KEY_NAME NEW_KEY_NAME
命令返回值:	修改成功时，返回 1 。 如果 NEW_KEY_NAME 已经存在，返回 0 。
*/
func (client *Client) RenameNX(ctx context.Context, key, newKey string) (bool, error) {
	boolCmd := client.goRedisClient.RenameNX(ctx, key, newKey)
	return boolCmd.Result()
}
