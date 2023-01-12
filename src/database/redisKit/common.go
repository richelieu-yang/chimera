package redisKit

import (
	"context"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
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

// ScanFully 对 Scan 进行了封装，用于替代 Keys 命令.
/*
PS:
(1) 如果db为空，将返回: [] <nil>
(2) redis cluster模式下，需要特殊处理（详见代码），否则：明明有数据的情况下，可能取不到数据，或者取到的数据不全（因为只找1个节点要）.
*/
func (client *Client) ScanFully(ctx context.Context, match string, count int64) ([]string, error) {
	var cursor uint64 = 0
	var s []string

	for {
		var tmp []string
		var err error
		tmp, cursor, err = client.Scan(ctx, cursor, match, count)
		if err != nil {
			return nil, err
		}

		s = sliceKit.Merge(s, tmp)
		if cursor == 0 {
			break
		}
	}
	return sliceKit.RemoveDuplicate(s), nil

	//if count <= 0 {
	//	count = 10
	//}
	//if clusterClient, ok := client.goRedisClient.(*redis.ClusterClient); ok {
	//	// cluster集群的情况，遍历每个master节点（由于主从复制，slave节点没必要去执行）
	//	var keys []string
	//
	//	err := clusterClient.ForEachMaster(ctx, func(ctx context.Context, client *redis.Client) error {
	//		tmp, err := scanFully(client, match, count)
	//		keys = sliceKit.Merge(keys, tmp)
	//		return err
	//	})
	//	if err != nil {
	//		return nil, err
	//	}
	//	return sliceKit.RemoveDuplicate(keys), nil
	//}
	//return scanFully(client.goRedisClient, match, count)
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
