package redisKit

import "context"

// Publish
/*
参考: https://www.runoob.com/redis/pub-sub-publish.html

命令说明:	命令用于将信息发送到指定的频道.
命令语法:	PUBLISH channel message
命令返回值:	接收到信息的订阅者数量.
*/
func (client *Client) Publish(ctx context.Context, channel string, message interface{}) (int64, error) {
	intCmd := client.goRedisClient.Publish(ctx, channel, message)
	return intCmd.Result()
}
