package redisKit

import (
	"context"
)

// HSet 将哈希表 key 中的字段 field 的值设为 value
/*
命令说明:
	为哈希表中的字段赋值 。
	PS:
	(1) 如果哈希表不存在，一个新的哈希表被创建并进行 HSET 操作。
	(2) 如果字段已经存在于哈希表中，旧值将被覆盖。
命令语法:	HSET KEY_NAME FIELD VALUE
命令返回值:	如果字段是哈希表中的一个新建字段，并且值设置成功，返回 1 。 如果哈希表中域字段已经存在且旧值已被新值覆盖，返回 0 。

PS:
(1) 3个传参都可以为""（3个全是""也可以）；
(2) 可以一次性设置多对 field 和 value.

@param key 		如果db中不存在，会自动创建
@param values 	accepts values in following formats:
	("myhash", "key1", "value1", "key2", "value2")
	("myhash", []string{"key1", "value1", "key2", "value2"})
	("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
*/
func (client *Client) HSet(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return client.goRedisClient.HSet(ctx, key, values...).Result()
}

// HSetNX 只有在 字段field 不存在时，设置哈希表字段的值
/*
命令说明:
	为哈希表中不存在的的字段赋值。
	PS:
	(1) 如果哈希表不存在，一个新的哈希表被创建并进行 HSET 操作。
	(2) 如果字段已经存在于哈希表中，操作无效。
	(3) 如果 key 不存在，一个新哈希表被创建并执行 HSETNX 命令。
命令语法:	HSETNX KEY_NAME FIELD VALUE
命令返回值:	设置成功，返回 1 。 如果给定字段已经存在且没有操作被执行，返回 0 。

PS:
(1) 如果field未存在，将返回: (true, nil).
(2) 如果field已存在，将返回: (false, nil)，且 此次设置哈希表字段的值 将无效.

@param key 如果db中不存在，会自动创建
*/
func (client *Client) HSetNX(ctx context.Context, key, field string, value interface{}) (bool, error) {
	return client.goRedisClient.HSetNX(ctx, key, field, value).Result()
}

// HGet 获取存储在哈希表中指定字段的值。
/*
@return key不存在 => 	("", redis.Nil)
		field不存在 => 	("", redis.Nil)
*/
func (client *Client) HGet(ctx context.Context, key, field string) (string, error) {
	return client.goRedisClient.HGet(ctx, key, field).Result()
}

// HDel 删除一个或多个哈希表字段.
/*
@param key 如果在db中不存在的话，将返回(0, nil)
@return 第一个返回值：被成功删除字段的数量，不包括被忽略的字段
*/
func (client *Client) HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	return client.goRedisClient.HDel(ctx, key, fields...).Result()
}

// HKeys 获取哈希表中所有字段
func (client *Client) HKeys(ctx context.Context, key string) ([]string, error) {
	return client.goRedisClient.HKeys(ctx, key).Result()
}

// HScan 迭代哈希表中的键值对
/*
@return e.g. 如果哈希表中有两对键值，那么不出错的情况下，返回的[]string实例的长度为4
*/
func (client *Client) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return client.goRedisClient.HScan(ctx, key, cursor, match, count).Result()
}
