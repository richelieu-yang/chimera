package redisKit

import (
	"context"
	"github.com/go-redis/redis/v9"
)

// GeoAdd
/*
命令说明: 存储指定的地理空间位置，可以将一个或多个经度(longitude)、纬度(latitude)、位置名称(member)添加到指定的 key 中.
命令语法: GEOADD key longitude latitude member [longitude latitude member ...]
*/
func (client *Client) GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) (int64, error) {
	intCmd := client.goRedisClient.GeoAdd(ctx, key, geoLocation...)
	return intCmd.Result()
}

// GeoPos
/*
命令说明: 从给定的 key 里返回所有指定名称(member)的位置（经度和纬度），不存在的返回 nil。
命令语法: GEOPOS key member [member ...]
*/
func (client *Client) GeoPos(ctx context.Context, key string, members ...string) ([]*redis.GeoPos, error) {
	geoPosCmd := client.goRedisClient.GeoPos(ctx, key, members...)
	return geoPosCmd.Result()
}
