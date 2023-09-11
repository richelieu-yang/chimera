package redisKit

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// Subscribe 频道的订阅（对频道的完全匹配）.
/*
PS:
(1) 建议 起一个goroutine 来进行订阅操作.
(2) *redis.PubSub实例的2种使用方法:	(a) PubSub.Channel()（推荐）
									(b) PubSub.ReceiveMessage()
(3) 取消订阅详见"_info.md".

命令说明:	订阅给定的一个或多个频道的信息.
命令语法:	SUBSCRIBE channel [channel ...]
命令返回值:	接收到的信息.

e.g. 监听 db0 中key的超时
	!!!: 需要先配置好Redis并重启，详见"Redis.wps".
	传参channels: "__keyevent@0__:expired"
*/
func (client *Client) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
	return client.universalClient.Subscribe(ctx, channels...)
}

// PSubscribe 模式的订阅（对频道的模糊匹配，通过*）.
/*
PS: 每个模式以 * 作为匹配符，e.g.
	it* 匹配所有以 it 开头的频道( it.news 、 it.blog 、 it.tweets 等等)；
	news.* 匹配所有以 news. 开头的频道( news.it 、 news.global.today 等等)，诸如此类.

命令说明:	命令订阅一个或多个符合给定模式的频道.
命令语法:	PSUBSCRIBE pattern [pattern ...]
命令返回值:	接收到的信息.

e.g. 监听 所有db 中key的超时
	!!!: 需要先配置好Redis并重启，详见"Redis.wps".
	传参channels: "__keyevent@*__:expired"
*/
func (client *Client) PSubscribe(ctx context.Context, patterns ...string) *redis.PubSub {
	return client.universalClient.PSubscribe(ctx, patterns...)
}

// SSubscribe
/*
TODO: 目前还未在网上找到"SSUBSCRIBE"命令的相关说明,
*/
func (client *Client) SSubscribe(ctx context.Context, channels ...string) *redis.PubSub {
	return client.universalClient.SSubscribe(ctx, channels...)
}

// Publish 发布.
/*
参考: https://www.runoob.com/redis/pub-sub-publish.html

命令说明:	命令用于将信息发送到指定的频道.
命令语法:	PUBLISH channel message
命令返回值:	接收到信息的订阅者数量.
*/
func (client *Client) Publish(ctx context.Context, channel string, message interface{}) (int64, error) {
	return client.universalClient.Publish(ctx, channel, message).Result()
}
