package redisKit

import (
	"context"
	"github.com/go-redis/redis/v8"
)

/*
Redis 有序集合(sorted set): https://www.runoob.com/redis/redis-sorted-sets.html
*/

// ZAdd [增]向有序集合添加一个或多个成员，或者更新已存在成员的分数
/*
@param key 如果在db中不存在的话，会自动创建
*/
func (client *Client) ZAdd(ctx context.Context, key string, members ...*redis.Z) (int64, error) {
	return client.UC.ZAdd(ctx, key, members...).Result()
}

// ZRem [删]移除有序集合中的一个或多个成员
/*
@return 如果传参key在db中不存在的话，将返回(0, nil)
*/
func (client *Client) ZRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return client.UC.ZRem(ctx, key, members...).Result()
}

// ZRangeByScore 通过分数返回有序集合指定区间内的成员，分数 从低到高 排序
/*
具体如何调用可以参考 ZRevRangeByScore.
*/
func (client *Client) ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) ([]string, error) {
	return client.UC.ZRangeByScore(ctx, key, opt).Result()
}

func (client *Client) ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) ([]redis.Z, error) {
	return client.UC.ZRangeByScoreWithScores(ctx, key, opt).Result()
}

// ZRevRangeByScore 返回有序集中指定分数区间内的成员，分数 从高到低 排序
/*
ZREVRANGEBYSCORE命令的语法: ZREVRANGEBYSCORE key max min [WITHSCORES] [LIMIT offset count]

@param key	如果在db中不存在的话，将返回([], nil)（[]string实例的len和cap都为0）.
@param opt	Max属性: 可以为"+inf"，Min属性: 可以为"-inf".
			默认情况下，区间的取值使用闭区间 (小于等于或大于等于)，你也可以通过给参数前增加 "("符号 来使用可选的开区间 (小于或大于).
			Offset属性: 	偏移量，从第几个开始（默认从0开始，即不偏移），但需要注意：设置此属性的值非0时，也要设置Count属性！！！
			Count属性: 	取几条数据？默认0，即返回所有符合条件的数据.（设置此属性时，可不设置Offset属性）

e.g. ZREVRANGEBYSCORE命令
ZRANGEBYSCORE zset (1 5		返回所有符合条件 1 < score <= 5 的成员
ZRANGEBYSCORE zset (5 (10	返回所有符合条件 5 < score < 10 的成员
*/
func (client *Client) ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) ([]string, error) {
	return client.UC.ZRevRangeByScore(ctx, key, opt).Result()
}

// ZRevRangeByScoreWithScores 类似 ZRevRangeByScore，返回值元素中多了score属性
func (client *Client) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) ([]redis.Z, error) {
	return client.UC.ZRevRangeByScoreWithScores(ctx, key, opt).Result()
}
