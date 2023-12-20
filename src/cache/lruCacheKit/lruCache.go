package cache

import "github.com/duke-git/lancet/v2/algorithm"

// NewLRUCache lru算法实现缓存。
/*
https://www.golancet.cn/api/packages/algorithm.html#LRUCache
*/
func NewLRUCache[K comparable, V any](capacity int) *algorithm.LRUCache[K, V] {
	return algorithm.NewLRUCache[K, V](capacity)
}
