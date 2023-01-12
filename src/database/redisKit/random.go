package redisKit

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/richelieu42/go-scales/src/randomKit"
)

// RandomKey 从当前数据库中随机返回一个key
/*
PS: 如果当前db是空的，将返回error（redis.Nil）.
*/
func (client *Client) RandomKey(ctx context.Context) (string, error) {
	return client.goRedisClient.RandomKey(ctx).Result()
}

// RandomKeyWithMatch 从当前数据库中随机返回一个key(指定match).
/*
Deprecated: 传参match 对应的key有很多的情况下，会有性能问题.
PS:
(1) 官方没有实现此功能的方法，此方法是通过Scan命令来实现的；
(2) 如果当前db不存在符合条件的key，将返回error(redis.Nil).
*/
func (client *Client) RandomKeyWithMatch(ctx context.Context, match string) (string, error) {
	s, err := client.ScanFully(ctx, match, -1)
	if err != nil {
		return "", err
	}

	length := len(s)
	if length > 0 {
		return s[randomKit.Int(length)], nil
	}
	return "", redis.Nil

	//// 可复用函数
	//common := func(cursor uint64, match string) (string, error) {
	//	s, _, err := client.Scan(cursor, match, 6)
	//	if err != nil {
	//		return "", err
	//	}
	//	length := len(s)
	//	if length > 0 {
	//		return s[randomKit.Int(length)], nil
	//	}
	//	return "", redis.Nil
	//}
	//
	//// 第1次，cursor为0~100的随机数，但由于乱传cursor，可能导致返回的切片实例为空，因此有了第2次
	//cursor := randomKit.Int(101)
	//str, err := common(uint64(cursor), match)
	//if err == nil {
	//	return str, nil
	//}
	//if err != redis.Nil {
	//	return "", err
	//}
	//// 第2次，cursor为0，避免情况：命名db中有符合条件的key，却告知调用方不存在符合条件的key
	//return common(0, match)
}
