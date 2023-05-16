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
@param expiration 	e.g.	120*time.Second			120s后过期
					 		0 						持久化的键（即TTL为-1），无论：键是否存在、存在的键是否有超时时间
					 		redis.KeepTTL(即-1) 	保持已经存在的TTL（需要确保Redis版本 >= 6.0，否则会返回error: ERR syntax error）
*/
func (client *Client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	statusCmd := client.universalClient.Set(ctx, key, value, expiration)
	str, err := statusCmd.Result()
	if err != nil {
		return false, err
	}
	return str == "OK", nil
}

// SetNX
/*
命令说明: 	在指定的 key 不存在时，为 key 设置指定的值.
命令语法: 	SETNX KEY_NAME VALUE
命令返回值: 	设置成功，返回 1 ；设置失败，返回 0 .

@return 第一个返回值代表: 是否设置成功
*/
func (client *Client) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	boolCmd := client.universalClient.SetNX(ctx, key, value, expiration)
	return boolCmd.Result()
}

// SetEx
/*
命令说明:	为指定的 key 设置值及其过期时间。如果 key 已经存在， SETEX 命令将会替换旧的值。
命令语法:	SETEX KEY_NAME TIMEOUT VALUE
命令返回值:	设置成功时返回 OK 。
*/
func (client *Client) SetEx(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	statusCmd := client.universalClient.SetEx(ctx, key, value, expiration)
	str, err := statusCmd.Result()
	if err != nil {
		return false, err
	}
	return str == "OK", nil
}

// MSet
/*
命令说明:	同时设置一个或多个 key-value 对。
命令语法:	MSET key1 value1 key2 value2 .. keyN valueN
命令返回值:	总是返回 OK 。
*/
func (client *Client) MSet(ctx context.Context, values ...interface{}) (bool, error) {
	statusCmd := client.universalClient.MSet(ctx, values...)
	str, err := statusCmd.Result()
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
	boolCmd := client.universalClient.MSetNX(ctx, values...)
	return boolCmd.Result()
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
	return client.universalClient.Get(ctx, key).Result()
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
	sliceCmd := client.universalClient.MGet(ctx, keys...)
	return sliceCmd.Result()
}
