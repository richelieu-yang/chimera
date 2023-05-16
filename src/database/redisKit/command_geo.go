package redisKit

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// GeoAdd
/*
命令说明: 存储指定的地理空间位置，可以将一个或多个经度(longitude)、纬度(latitude)、位置名称(member)添加到指定的 key 中.
命令语法: GEOADD key longitude latitude member [longitude latitude member ...]
*/
func (client *Client) GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) (int64, error) {
	cmd := client.universalClient.GeoAdd(ctx, key, geoLocation...)
	return cmd.Result()
}

// GeoPos
/*
命令说明: 从给定的 key 里返回所有指定名称(member)的位置（经度和纬度），不存在的返回 nil。
命令语法: GEOPOS key member [member ...]
*/
func (client *Client) GeoPos(ctx context.Context, key string, members ...string) ([]*redis.GeoPos, error) {
	cmd := client.universalClient.GeoPos(ctx, key, members...)
	return cmd.Result()
}

// GeoDist
/*
命令说明: 用于返回两个给定位置之间的距离.
命令语法: GEODIST key member1 member2 [m|km|ft|mi]
*/
func (client *Client) GeoDist(ctx context.Context, key string, member1, member2, unit string) (float64, error) {
	cmd := client.universalClient.GeoDist(ctx, key, member1, member2, unit)
	return cmd.Result()
}
