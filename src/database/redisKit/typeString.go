package redisKit

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

// Set 设置指定key的值.
/*
@param key 			可以为""
@param expiration 	e.g.	120*time.Second
					 		0: 						持久化的键（即TTL为-1），无论：键是否存在、存在的键是否有超时时间
					 		-1(即 redis.KeepTTL): 	保持已经存在的TTL.(1)如果key不存在，则TTL为-1；(2)如果"Redis服务器版本"<6.0，会报错：ERR syntax error.
@return 第一个返回值代表: 是否设置成功
*/
func (client Client) Set(key string, value interface{}, expiration time.Duration) (bool, error) {
	reply, err := client.UC.Set(context.TODO(), key, value, expiration).Result()
	if err != nil {
		return false, err
	}
	return reply == "OK", nil
}

// SetNX 只有在 key 不存在时设置 key 的值
/*
@return 第一个返回值代表: 是否设置成功
*/
func (client Client) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return client.UC.SetNX(context.TODO(), key, value, expiration).Result()
}

// Get
/*
PS:
(1) 如果当前db中不存在传参key，将返回 ("", redis.Nil)
(2) 如果对应value的类型不为string，会返回error:	WRONGTYPE Operation against a key holding the wrong kind of value
*/
func (client Client) Get(key string) (string, error) {
	return client.UC.Get(context.TODO(), key).Result()
}

// GetWithoutRedisNil
/*
PS：
(1) 如果当前db中不存在传参key，将返回 ("", nil)
(2) 如果不关心key是否存在，只关心值，可以调用此方法
(3) 如果对应value的类型不为string，会返回error:	WRONGTYPE Operation against a key holding the wrong kind of value
*/
func (client Client) GetWithoutRedisNil(key string) (string, error) {
	str, err := client.Get(key)
	if err != nil {
		if err != redis.Nil {
			return "", err
		}
		// 当前db中不存在传参key的情况
		return "", nil
	}
	return str, nil
}
