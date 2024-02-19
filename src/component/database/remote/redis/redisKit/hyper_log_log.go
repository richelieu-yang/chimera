// Package redisKit
/*
Redis 如何使用 HyperLogLog
	https://blog.csdn.net/SunnyYoona/article/details/124764009
*/
package redisKit

import "context"

func (client *Client) PFAdd(ctx context.Context, key string, els ...interface{}) (int64, error) {
	cmd := client.universalClient.PFAdd(ctx, key, els...)
	return cmd.Result()
}

func (client *Client) PFCount(ctx context.Context, keys ...string) (int64, error) {
	cmd := client.universalClient.PFCount(ctx, keys...)
	return cmd.Result()
}

func (client *Client) PFMerge(ctx context.Context, dest string, keys ...string) (string, error) {
	cmd := client.universalClient.PFMerge(ctx, dest, keys...)
	return cmd.Result()
}
