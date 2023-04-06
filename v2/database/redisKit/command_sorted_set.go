package redisKit

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// ZAdd
/*
命令说明:	将一个或多个成员元素及其分数值加入到有序集当中。
			(1) 如果某个成员已经是有序集的成员，那么更新这个成员的分数值，并通过重新插入这个成员元素，来保证该成员在正确的位置上。
			(2) 分数值可以是 整数值 或 双精度浮点数。
			(3) 如果有序集合 key 不存在，则创建一个空的有序集并执行 ZADD 操作。
			(4) 当 key 存在但不是有序集类型时，返回一个错误。
命令语法:	ZADD KEY_NAME SCORE1 VALUE1.. SCOREN VALUEN
命令返回值:	被成功添加的新成员的数量，不包括那些被更新的、已经存在的成员。
*/
func (client *Client) ZAdd(ctx context.Context, key string, members ...redis.Z) (int64, error) {
	intCmd := client.core.ZAdd(ctx, key, members...)
	return intCmd.Result()
}

// ZRem
/*
命令说明:	移除有序集中的一个或多个成员，不存在的成员将被忽略。(当 key 存在但不是有序集类型时，返回一个错误。)
命令语法:	ZREM key member [member ...]
命令返回值:	被成功移除的成员的数量，不包括被忽略的成员。
*/
func (client *Client) ZRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	intCmd := client.core.ZRem(ctx, key, members...)
	return intCmd.Result()
}

// ZCard
/*
命令说明:	计算集合中元素的数量。
命令语法:	ZCARD KEY_NAME
命令返回值:	当 key 存在且是有序集类型时，返回有序集的基数。 当 key 不存在时，返回 0。
*/
func (client *Client) ZCard(ctx context.Context, key string) (int64, error) {
	intCmd := client.core.ZCard(ctx, key)
	return intCmd.Result()
}

// ZCount
/*
命令说明:	计算有序集合中指定分数区间的成员数量。
命令语法:	ZCOUNT key min max
命令返回值:	分数值在 min 和 max 之间的成员的数量。
*/
func (client *Client) ZCount(ctx context.Context, key, min, max string) (int64, error) {
	intCmd := client.core.ZCount(ctx, key, min, max)
	return intCmd.Result()
}

// ZRangeByScore
/*
命令说明:
	返回有序集合中指定分数区间的成员列表。有序集成员按分数值递增(从小到大)次序排列。
	具有相同分数值的成员按字典序来排列(该属性是有序集提供的，不需要额外的计算)。
	默认情况下，区间的取值使用闭区间 (小于等于或大于等于)，你也可以通过给参数前增加 ( 符号来使用可选的开区间 (小于或大于)。
命令语法:	ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]
命令返回值:	指定区间内，带有分数值(可选)的有序集成员的列表。
*/
func (client *Client) ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) ([]string, error) {
	stringSliceCmd := client.core.ZRangeByScore(ctx, key, opt)
	return stringSliceCmd.Result()
}

// ZRangeByScoreWithScores
/*
详见: ZRangeByScore，命令中有 WITHSCORES.
*/
func (client *Client) ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) ([]redis.Z, error) {
	zSliceCmd := client.core.ZRangeByScoreWithScores(ctx, key, opt)
	return zSliceCmd.Result()
}

// ZRevRangeByScore
/*
命令说明:
	返回有序集中指定分数区间内的所有的成员。有序集成员按分数值递减(从大到小)的次序排列。
	具有相同分数值的成员按字典序的逆序(reverse lexicographical order )排列。
	除了成员按分数值递减的次序排列这一点外， ZREVRANGEBYSCORE 命令的其他方面和 ZRANGEBYSCORE 命令一样。
命令语法:	ZREVRANGEBYSCORE key max min [WITHSCORES] [LIMIT offset count]
命令返回值:	指定区间内，带有分数值(可选)的有序集成员的列表。
*/
func (client *Client) ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) ([]string, error) {
	stringSliceCmd := client.core.ZRevRangeByScore(ctx, key, opt)
	return stringSliceCmd.Result()
}

// ZRevRangeByScoreWithScores
/*
详见: ZRevRangeByScore，命令中有 WITHSCORES.
*/
func (client *Client) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) ([]redis.Z, error) {
	zSliceCmd := client.core.ZRevRangeByScoreWithScores(ctx, key, opt)
	return zSliceCmd.Result()
}
