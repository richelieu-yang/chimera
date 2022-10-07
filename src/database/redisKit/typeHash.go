package redisKit

import (
	"context"
)

/*
Redis 哈希(Hash): https://www.runoob.com/redis/redis-hashes.html
*/

// HSet 将哈希表 key 中的字段 field 的值设为 value
/*
PS:
(1) 3个传参都可以为""（3个全是""也可以）；
(2) 可以一次性设置多对 field 和 value.

@param key 		如果db中不存在，会自动创建
@param values 	accepts values in following formats:
	("myhash", "key1", "value1", "key2", "value2")
	("myhash", []string{"key1", "value1", "key2", "value2"})
	("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
*/
func (client Client) HSet(key string, values ...interface{}) (int64, error) {
	return client.UC.HSet(context.Background(), key, values...).Result()
}

// HSetNX 只有在 字段field 不存在时，设置哈希表字段的值
/*
PS:
(1) 如果field未存在，将返回: (true, nil).
(2) 如果field已存在，将返回: (false, nil)，且 此次设置哈希表字段的值 将无效.

@param key 如果db中不存在，会自动创建
*/
func (client Client) HSetNX(key, field string, value interface{}) (bool, error) {
	return client.UC.HSetNX(context.Background(), key, field, value).Result()
}

// HGet 获取存储在哈希表中指定字段的值。
/*
@return key不存在 => 	("", redis.Nil)
		field不存在 => 	("", redis.Nil)
*/
func (client Client) HGet(key, field string) (string, error) {
	return client.UC.HGet(context.Background(), key, field).Result()
}

// HDel 删除一个或多个哈希表字段.
/*
@param key 如果在db中不存在的话，将返回(0, nil)
@return 第一个返回值：被成功删除字段的数量，不包括被忽略的字段
*/
func (client Client) HDel(key string, fields ...string) (int64, error) {
	return client.UC.HDel(context.Background(), key, fields...).Result()
}

// HKeys 获取哈希表中所有字段
func (client Client) HKeys(key string) ([]string, error) {
	return client.UC.HKeys(context.Background(), key).Result()
}

// HScan 迭代哈希表中的键值对
/*
@return e.g. 如果哈希表中有两对键值，那么不出错的情况下，返回的[]string实例的长度为4
*/
func (client Client) HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return client.UC.HScan(context.Background(), key, cursor, match, count).Result()
}

//func (client Client) HScanFully(key string, match string, count int64) ([]string, error) {
//	s, cursor, err := client.HScan(key, 0, match, count)
//	if err != nil {
//		return nil, err
//	}
//
//	for cursor != 0 {
//		var tmp []string
//		tmp, cursor, err = client.HScan(key, cursor, match, count)
//		if err != nil {
//			return nil, err
//		}
//		s = sliceKit.Merge(s, tmp)
//	}
//	// 返回前去重
//	//return sliceKit.RemoveDuplicate(s), nil
//	return s, nil
//}
