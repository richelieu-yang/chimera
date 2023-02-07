package redisKit

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// Subscribe
/*
命令说明:	订阅给定的一个或多个频道的信息.
命令语法:	SUBSCRIBE channel [channel ...]
命令返回值:	接收到的信息.
*/
func (client *Client) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
	return client.goRedisClient.Subscribe(ctx, channels...)
}

// PSubscribe
/*
PS: 每个模式以 * 作为匹配符，e.g.
	it* 匹配所有以 it 开头的频道( it.news 、 it.blog 、 it.tweets 等等)；
	news.* 匹配所有以 news. 开头的频道( news.it 、 news.global.today 等等)，诸如此类.

命令说明:	命令订阅一个或多个符合给定模式的频道.
命令语法:	PSUBSCRIBE pattern [pattern ...]
命令返回值:	接收到的信息.
*/
func (client *Client) PSubscribe(ctx context.Context, patterns ...string) *redis.PubSub {
	return client.goRedisClient.PSubscribe(ctx, patterns...)
}

// SSubscribe
/*
TODO: 目前还未在网上找到"ssubscribe"命令的相关说明,
*/
func (client *Client) SSubscribe(ctx context.Context, channels ...string) *redis.PubSub {
	return client.goRedisClient.SSubscribe(ctx, channels...)
}
