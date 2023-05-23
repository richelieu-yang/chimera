package mapKit

// Copy 浅拷贝
/*
@param src 可以为nil
@return 保底为空的map实例（不为nil）
*/
func Copy[K comparable, V any](src map[K]V) map[K]V {
	dest := map[K]V{}
	for k, v := range src {
		dest[k] = v
	}
	return dest
}
