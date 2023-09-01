package redisKit

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

// Set 设置指定key的值（string类型）.
/*
命令说明:	设置给定 key 的值。如果 key 已经存储其他值， SET 就覆写旧值，且无视类型。
命令语法:	SET KEY_NAME VALUE
命令返回值:
	在 Redis 2.6.12 以前版本， SET 命令总是返回 OK 。
	从 Redis 2.6.12 版本开始， SET 在设置操作成功完成时，才返回 OK 。

@param key 			可以为""
@param value 		支持的类型: string、[]byte、int、float64、bool(true: "1"; false: "0")...
					不支持的类型（会返回error）: map、自定义结构体...
@param expiration 	e.g. 120*time.Second		120s后过期
					 	 0 						持久化的键（即TTL为-1），无论：键是否存在、存在的键是否有超时时间
					 	 -1(即redis.KeepTTL)	（需要确保Redis版本 >= 6.0，否则会返回error: ERR syntax error）
												(a) 键存在:		保持已经存在的TTL
												(b) 键不存在:	持久化的键
*/
func (client *Client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	cmd := client.universalClient.Set(ctx, key, value, expiration)
	str, err := cmd.Result()
	if err != nil {
		return false, err
	}
	return str == "OK", nil
}

// SetNX key不存在才会: 设置值 && 更新TTL
/*
PS: 如果传参key已经存在，不会修改该key的TTL.

命令说明: 	在指定的 key 不存在时，为 key 设置指定的值.
命令语法: 	SETNX KEY_NAME VALUE
命令返回值: 	设置成功，返回 1 ；设置失败，返回 0 .

@return 第一个返回值代表: 是否设置成功
*/
func (client *Client) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	cmd := client.universalClient.SetNX(ctx, key, value, expiration)
	return cmd.Result()
}

// SetEx
/*
Deprecated: 个人感觉在 go-redis 中，使用 Set 就够了.

命令说明:	为指定的 key 设置值及其过期时间。。(如果 key 已经存在，SETEX 命令将会 替换旧的值 并 更新TTL)
命令语法:	SETEX KEY_NAME TIMEOUT VALUE
命令返回值:	设置成功时返回 OK 。
*/
func (client *Client) SetEx(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	cmd := client.universalClient.SetEx(ctx, key, value, expiration)
	str, err := cmd.Result()
	if err != nil {
		return false, err
	}
	return str == "OK", nil
}

// SetXX （与 SetNX 相反）key存在才会: 设置值 && 更新TTL
/*
Redis的数据类型详解 https://blog.csdn.net/m0_53474063/article/details/112647028
*/
func (client *Client) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	cmd := client.universalClient.SetXX(ctx, key, value, expiration)
	return cmd.Result()
}

// MSet
/*
命令说明:	同时设置一个或多个 key-value 对。
命令语法:	MSET key1 value1 key2 value2 .. keyN valueN
命令返回值:	总是返回 OK 。
*/
func (client *Client) MSet(ctx context.Context, values ...interface{}) (bool, error) {
	cmd := client.universalClient.MSet(ctx, values...)
	str, err := cmd.Result()
	if err != nil {
		return false, err
	}
	return str == "OK", nil
}

// MSetNX
/*
命令说明:	所有给定 key 都不存在时，同时设置一个或多个 key-value 对。
命令语法:	MSETNX key1 value1 key2 value2 .. keyN valueN
命令返回值:	当所有 key 都成功设置，返回 1 。 如果所有给定 key 都设置失败(至少有一个 key 已经存在)，那么返回 0 。
*/
func (client *Client) MSetNX(ctx context.Context, values ...interface{}) (bool, error) {
	cmd := client.universalClient.MSetNX(ctx, values...)
	return cmd.Result()
}

// Get
/*
命令说明:	获取指定 key 的值。（如果 key 不存在，返回 nil；如果key 储存的值不是字符串类型，返回一个错误）
命令语法:	GET KEY_NAME
命令返回值:	返回 key 的值，如果 key 不存在时，返回 nil。 如果 key 不是字符串类型，那么返回一个错误。

e.g.
当前db中不存在 传参key => ("", redis.Nil)
*/
func (client *Client) Get(ctx context.Context, key string) (string, error) {
	cmd := client.universalClient.Get(ctx, key)
	return cmd.Result()
}

// GetWithoutRedisNil 对 Get 进行封装（特殊处理）: 当前db中不存在传参key时，返回 ("", nil).
/*
PS：
(1) 如果当前db中不存在传参key，将返回 ("", nil)
(2) 如果不关心key是否存在，只关心值，可以调用此方法
(3) 如果对应value的类型不为string，会返回error:	WRONGTYPE Operation against a key holding the wrong kind of value
*/
func (client *Client) GetWithoutRedisNil(ctx context.Context, key string) (string, error) {
	str, err := client.Get(ctx, key)
	if err != nil {
		if err != redis.Nil {
			return "", err
		}
		// err == redis.Nil: 当前db中不存在 传参key
		return "", nil
	}
	return str, nil
}

// MGet
/*
命令说明:	返回所有(一个或多个)给定 key 的值。 如果给定的 key 里面，有某个 key 不存在，那么这个 key 返回特殊值 nil 。
命令语法:	MGET KEY1 KEY2 .. KEYN
命令返回值:	一个包含所有给定 key 的值的列表。
*/
func (client *Client) MGet(ctx context.Context, keys ...string) ([]interface{}, error) {
	cmd := client.universalClient.MGet(ctx, keys...)
	return cmd.Result()
}

// Incr
/*
命令说明:
Redis Incr 命令将 key 中储存的数字值增一。
(1) 如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 INCR 操作。
(2) 如果值包含错误的类型，或字符串类型的值不能表示为数字，那么返回一个错误。
(3) 本操作的值限制在 64 位(bit)有符号数字表示之内。

e.g. 	key不存在的情况
	(context.Background(), "a") => 1 <nil>

e.g.1	key存在且key为"1"的情况
	(context.Background(), "a") => 2 <nil>

e.g.2	key存在且key为"-1000"的情况
	(context.Background(), "a") => -999 <nil>
*/
func (client *Client) Incr(ctx context.Context, key string) (int64, error) {
	cmd := client.universalClient.Incr(ctx, key)
	return cmd.Result()
}

// IncrBy
/*
命令说明:
Redis Incrby 命令将 key 中储存的数字加上指定的增量值。
(1) 如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 INCRBY 命令。
(2) 如果值包含错误的类型，或字符串类型的值不能表示为数字，那么返回一个错误。
(3) 本操作的值限制在 64 位(bit)有符号数字表示之内。
*/
func (client *Client) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	cmd := client.universalClient.IncrBy(ctx, key, value)
	return cmd.Result()
}

// IncrByFloat
/*
命令说明:
Redis Incrbyfloat 命令为 key 中所储存的值加上指定的浮点数增量值。
(1) 如果 key 不存在，那么 INCRBYFLOAT 会先将 key 的值设为 0 ，再执行加法操作。
*/
func (client *Client) IncrByFloat(ctx context.Context, key string, value float64) (float64, error) {
	cmd := client.universalClient.IncrByFloat(ctx, key, value)
	return cmd.Result()
}

// Decr
/*
命令说明:
Redis Decr 命令将 key 中储存的数字值减一。
(1) 如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 DECR 操作。
(2) 如果值包含错误的类型，或字符串类型的值不能表示为数字，那么返回一个错误。
(3) 本操作的值限制在 64 位(bit)有符号数字表示之内。
*/
func (client *Client) Decr(ctx context.Context, key string) (int64, error) {
	cmd := client.universalClient.Decr(ctx, key)
	return cmd.Result()
}

// DecrBy
/*
命令说明:
Redis Decrby 命令将 key 所储存的值减去指定的减量值。
(1) 如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 DECRBY 操作。
(2) 如果值包含错误的类型，或字符串类型的值不能表示为数字，那么返回一个错误。
(3) 本操作的值限制在 64 位(bit)有符号数字表示之内。
*/
func (client *Client) DecrBy(ctx context.Context, key string, decrement int64) (int64, error) {
	cmd := client.universalClient.DecrBy(ctx, key, decrement)
	return cmd.Result()
}

// Append
/*
命令说明:
Redis Append 命令用于为指定的 key 追加值。
(1) 如果 key 已经存在并且是一个字符串， APPEND 命令将 value 追加到 key 原来的值的末尾。
(2) 如果 key 不存在， APPEND 就简单地将给定 key 设为 value ，就像执行 SET key value 一样。

@return 追加后值的长度
*/
func (client *Client) Append(ctx context.Context, key, value string) (int64, error) {
	cmd := client.universalClient.Append(ctx, key, value)
	return cmd.Result()
}
