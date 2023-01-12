package redisKit

import "context"

// SAdd 向集合添加一个或多个成员
/*
语法: SADD key member1 [member2]

@param key 不存在的话，会自动创建此key
*/
func (client Client) SAdd(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return client.UC.SAdd(ctx, key, members...).Result()
}

// SRem 移除集合中一个或多个成员
/*
语法: SREM key member1 [member2]

@param key 如果移除成员后集合为空，将删除此key
*/
func (client Client) SRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return client.UC.SRem(ctx, key, members...).Result()
}

// SMembers 返回集合中的所有成员
/*
语法: SMEMBERS key

@return e.g. 	([a c b], nil)
		e.g.1	如果key不存在，将返回 ([], nil)
*/
func (client Client) SMembers(ctx context.Context, key string) ([]string, error) {
	return client.UC.SMembers(ctx, key).Result()
}

// SMembersMap
/*
类似于 SMembers，返回值中map实例的所有值都是空结构体(struct{}).

@return 如果key不存在，将返回 (map[], nil)
*/
func (client Client) SMembersMap(ctx context.Context, key string) (map[string]struct{}, error) {
	return client.UC.SMembersMap(ctx, key).Result()
}
