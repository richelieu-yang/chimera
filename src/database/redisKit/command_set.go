package redisKit

import "context"

// SAdd 将一个或多个成员元素加入到集合中，已经存在于集合的成员元素将被忽略.
/*
PS:
(1) 假如集合 key 不存在，则创建一个只包含添加的元素作成员的集合；
(2) 当集合 key 不是集合类型时，返回一个错误.

@return 第1个返回值: 成功添加的成员的数量（已经在集合中的不算）
*/
func (client *Client) SAdd(ctx context.Context, key string, members ...interface{}) (int64, error) {
	cmd := client.universalClient.SAdd(ctx, key, members...)
	return cmd.Result()
}

// SRem 移除集合中的一个或多个成员元素，不存在的成员元素会被忽略.
/*
PS:
(1) 移除后，如果集合为空，将自动删除 传参key.
(2) 当 key 不是集合类型，返回一个错误.

@return 第1个返回值: 被成功移除的元素的数量，不包括被忽略的元素（集合中本来就不存在该元素）
*/
func (client *Client) SRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	cmd := client.universalClient.SRem(ctx, key, members...)
	return cmd.Result()
}

// SMembers 返回集合中的所有的成员（不存在的集合 key 被视为空集合）.
/*
Deprecated: 当 SMEMBERS命令 被用于处理一个大的数据库时，它可能会阻塞服务器达数秒之久。

e.g.
key不存在的情况 => ([]string{}, nil)
*/
func (client *Client) SMembers(ctx context.Context, key string) ([]string, error) {
	cmd := client.universalClient.SMembers(ctx, key)
	return cmd.Result()
}

// SIsMember 判断 member元素 是否是 key对应集合 的成员？
func (client *Client) SIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	cmd := client.universalClient.SIsMember(ctx, key, member)
	return cmd.Result()
}

// SMembersMap
/*
Deprecated: 当 SMEMBERS命令 被用于处理一个大的数据库时，它可能会阻塞服务器达数秒之久。

@return 第1个返回值: key为成员的值，value为空结构体（无意义）

e.g. key不存在的情况
=> (map[string]struct{}{}, nil)
*/
func (client *Client) SMembersMap(ctx context.Context, key string) (map[string]struct{}, error) {
	cmd := client.universalClient.SMembersMap(ctx, key)
	return cmd.Result()
}

// SPop 用于移除并返回集合(set)中的一个随机元素.
func (client *Client) SPop(ctx context.Context, key string) (string, error) {
	cmd := client.universalClient.SPop(ctx, key)
	return cmd.Result()
}

func (client *Client) SPopN(ctx context.Context, key string, count int64) ([]string, error) {
	cmd := client.universalClient.SPopN(ctx, key, count)
	return cmd.Result()
}

// SRandMember 返回集合中 1个 随机数.
func (client *Client) SRandMember(ctx context.Context, key string) (string, error) {
	cmd := client.universalClient.SRandMember(ctx, key)
	return cmd.Result()
}

// SRandMemberN 返回集合中 多个 随机数.
func (client *Client) SRandMemberN(ctx context.Context, key string, count int64) ([]string, error) {
	cmd := client.universalClient.SRandMemberN(ctx, key, count)
	return cmd.Result()
}
