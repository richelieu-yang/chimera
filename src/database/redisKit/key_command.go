package redisKit

import (
	"context"
)

// Type 返回 key 所储存的值的类型.
/*
语法:
	TYPE KEY_NAME
返回值:
	none (key不存在)
	string (字符串)
	list (列表)
	set (集合)
	zset (有序集)
	hash (哈希表)

e.g.
传参key不存在的情况 => ("none", nil)
*/
func (client *Client) Type(key string) (string, error) {
	return client.UC.Type(context.TODO(), key).Result()
}
