package redisKit

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
)

func (client *Client) Ping() (string, error) {
	return client.UC.Ping(context.TODO()).Result()
}

// Del （删）key 存在时，删除 key
/*
@return 第一个返回值代表：是否删除成功

e.g.
如果key不存在，将返回: (false, nil)
*/
func (client *Client) Del(ctx context.Context, key string) (bool, error) {
	reply, err := client.UC.Del(ctx, key).Result()
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
	reply, err := client.UC.Exists(ctx, keys...).Result()
	if err != nil {
		return false, err
	}
	return reply == 1, nil
}

func scan(client redis.UniversalClient, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return client.Scan(context.TODO(), cursor, match, count).Result()
}

func scanFully(client redis.UniversalClient, match string, count int64) ([]string, error) {
	// 使用 0 作为游标，开始新的迭代
	keys, cursor, err := scan(client, 0, match, count)
	if err != nil {
		return nil, err
	}

	for cursor != 0 {
		var tmp []string
		tmp, cursor, err = scan(client, cursor, match, count)
		if err != nil {
			return nil, err
		}
		keys = sliceKit.Merge(keys, tmp)
	}
	// 返回前去重
	return sliceKit.RemoveDuplicate(keys), nil
}

// Scan 迭代当前数据库中的数据库键
/*
！！！：
(1)	scan命令也并不是完美的，它"返回的结果有可能重复"，因此需要客户端"去重"；
(2) 用于替代keys，因为keys在大数据量有性能问题；
(3) 如果db为空，将返回: [] 0 <nil>
(4) 返回的[]string实例的长度可能会大于传参count，比如瞎传cursor的情况，编码时得注意.

@return 3个值分别为：keys、新的cursor、err
*/
func (client *Client) Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	return scan(client.UC, cursor, match, count)
}

// ScanFully 对 Scan 进行了封装，用于替代 Keys 命令
/*
PS:
(1) 如果db为空，将返回: [] <nil>
(2) redis cluster模式下，需要特殊处理（详见代码），否则：明明有数据的情况下，可能取不到数据，或者取到的数据不全（因为只找1个节点要）.
*/
func (client *Client) ScanFully(match string, count int64) ([]string, error) {
	if count <= 0 {
		count = 10
	}

	if clusterClient, ok := client.UC.(*redis.ClusterClient); ok {
		// cluster集群的情况，遍历每个master节点（由于主从复制，slave节点没必要去执行）
		var keys []string

		err := clusterClient.ForEachMaster(context.TODO(), func(ctx context.Context, client *redis.Client) error {
			tmp, err := scanFully(client, match, count)
			keys = sliceKit.Merge(keys, tmp)
			return err
		})
		if err != nil {
			return nil, err
		}
		return sliceKit.RemoveDuplicate(keys), nil
	}
	return scanFully(client.UC, match, count)
}

// Keys
// Deprecated: 禁止在生产环境使用Keys正则匹配操作（实际即便是开发、测试环境也要慎重使用）！！！
func (client *Client) Keys(match string) ([]string, error) {
	return client.UC.Keys(context.TODO(), match).Result()
}

// Publish 发布
/*
e.g.
("", "") => nil
*/
func (client *Client) Publish(channel string, message interface{}) error {
	_, err := client.UC.Publish(context.TODO(), channel, message).Result()
	return err
}

// FlushDB 清空当前数据库中的所有 key
/*
慎用！！！
*/
func (client *Client) FlushDB() error {
	_, err := client.UC.FlushDB(context.TODO()).Result()
	return err
}

// FlushAll 清空整个 Redis 服务器的数据(删除所有数据库的所有 key )
/*
慎用！！！
*/
func (client *Client) FlushAll() error {
	_, err := client.UC.FlushAll(context.TODO()).Result()
	return err
}
