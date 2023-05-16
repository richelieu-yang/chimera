package redisKit

import "context"

// FlushDB
/*
命令说明: 	清空当前数据库中的所有 key.
命令语法: 	FLUSHDB
命令返回值: 	总是返回 OK .
*/
func (client *Client) FlushDB(ctx context.Context) (string, error) {
	cmd := client.universalClient.FlushDB(ctx)
	return cmd.Result()
}

// FlushDBAsync 异步地.
func (client *Client) FlushDBAsync(ctx context.Context) (string, error) {
	cmd := client.universalClient.FlushDBAsync(ctx)
	return cmd.Result()
}

// FlushAll
/*
命令说明:	清空整个 Redis 服务器的数据(删除所有数据库的所有 key ).
命令语法:	FLUSHALL
命令返回值:	总是返回 OK .
*/
func (client *Client) FlushAll(ctx context.Context) (string, error) {
	cmd := client.universalClient.FlushAll(ctx)
	return cmd.Result()
}

// FlushAllAsync 异步地.
func (client *Client) FlushAllAsync(ctx context.Context) (string, error) {
	cmd := client.universalClient.FlushAllAsync(ctx)
	return cmd.Result()
}

// DBSize
/*
命令说明: 	返回当前数据库的 key 的数量.
命令语法:	DBSIZE
命令返回值:	当前数据库的 key 的数量.
*/
func (client *Client) DBSize(ctx context.Context) (int64, error) {
	cmd := client.universalClient.DBSize(ctx)
	return cmd.Result()
}
