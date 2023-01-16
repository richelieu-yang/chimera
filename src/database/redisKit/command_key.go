package redisKit

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
	"sync"
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

// ScanFully 对 Scan 进行了封装，用于替代 Keys 命令.
/*
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
