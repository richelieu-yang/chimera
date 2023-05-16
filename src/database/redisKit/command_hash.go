package redisKit

import (
	"context"
)

// HExists
/*
命令说明:	用于查看哈希表的指定字段是否存在
命令语法:	HEXISTS KEY_NAME FIELD_NAME
命令返回值:	如果哈希表含有给定字段，返回 1 。 如果哈希表不含有给定字段，或 key 不存在，返回 0 。
*/
func (client *Client) HExists(ctx context.Context, key, field string) (bool, error) {
	cmd := client.universalClient.HExists(ctx, key, field)
	return cmd.Result()
}

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
	cmd := client.universalClient.HSet(ctx, key, values...)
	return cmd.Result()
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
	cmd := client.universalClient.HSetNX(ctx, key, field, value)
	return cmd.Result()
}

// HGet
/*
命令说明:	返回哈希表中指定字段的值。
命令语法:	HGET KEY_NAME FIELD_NAME
命令返回值:	返回给定字段的值。如果给定的字段或 key 不存在时，返回 nil 。
*/
func (client *Client) HGet(ctx context.Context, key, field string) (string, error) {
	cmd := client.universalClient.HGet(ctx, key, field)
	return cmd.Result()
}

// HGetAll
/*
命令说明:
	返回哈希表中，所有的字段和值。
	在返回值里，紧跟每个字段名(field name)之后是字段的值(value)，所以返回值的长度是哈希表大小的两倍。
命令语法:	HGETALL KEY_NAME
命令返回值:	以列表形式返回哈希表的字段及字段值。 若 key 不存在，返回空列表。
*/
func (client *Client) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	cmd := client.universalClient.HGetAll(ctx, key)
	return cmd.Result()
}

// HDel
/*
命令说明:	删除哈希表 key 中的一个或多个指定字段，不存在的字段将被忽略。
命令语法:	HDEL KEY_NAME FIELD1.. FIELDN
命令返回值:	被成功删除字段的数量，不包括被忽略的字段。
*/
func (client *Client) HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	cmd := client.universalClient.HDel(ctx, key, fields...)
	return cmd.Result()
}

func (client *Client) HKeys(ctx context.Context, key string) ([]string, error) {
	cmd := client.universalClient.HKeys(ctx, key)
	return cmd.Result()
}

func (client *Client) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	cmd := client.universalClient.HScan(ctx, key, cursor, match, count)
	return cmd.Result()
}
