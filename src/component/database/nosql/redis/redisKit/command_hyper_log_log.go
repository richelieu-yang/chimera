// Package redisKit
/*
Redis HyperLogLog
	https://www.runoob.com/redis/redis-hyperloglog.html
Redis 如何使用 HyperLogLog
	https://blog.csdn.net/SunnyYoona/article/details/124764009
*/
package redisKit

import "context"

// PFAdd 添加指定元素到 HyperLogLog 中。
/*
语法: PFADD key element [element ...]
*/
func (client *Client) PFAdd(ctx context.Context, key string, els ...interface{}) (int64, error) {
	cmd := client.universalClient.PFAdd(ctx, key, els...)
	return cmd.Result()
}

// PFCount 返回给定 HyperLogLog 的基数估算值。
/*
语法: PFCOUNT key [key ...]
*/
func (client *Client) PFCount(ctx context.Context, keys ...string) (int64, error) {
	cmd := client.universalClient.PFCount(ctx, keys...)
	return cmd.Result()
}

// PFMerge 将多个 HyperLogLog 合并为一个 HyperLogLog。
/*
语法: PFMERGE destkey sourcekey [sourcekey ...]
*/
func (client *Client) PFMerge(ctx context.Context, dest string, keys ...string) (string, error) {
	cmd := client.universalClient.PFMerge(ctx, dest, keys...)
	return cmd.Result()
}
