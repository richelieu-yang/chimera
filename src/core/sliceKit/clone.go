package sliceKit

// Clone 浅克隆（浅拷贝）
/*
参考:
golang复制切片的方法（避免操作同一底层数组） https://blog.csdn.net/weixin_43970884/article/details/126051345

@param src 	可以为nil，将返回一个非nil的len==0的slice
@return		必定不为nil
*/
func Clone[T any](src []T) []T {
	dolly := make([]T, len(src))
	copy(dolly, src)
	return dolly
}

// Clone1 浅克隆（浅拷贝）
/*
Deprecated: 性能不如 Clone()

@param src 可以为nil，将返回nil
*/
func Clone1[T any](src []T) []T {
	return append(src)
}

// Clone2 浅克隆（浅拷贝）
/*
Deprecated: 返回值可能为nil

@param src 可以为nil，将返回nil
*/
func Clone2[T any](src []T) []T {
	return src[:]
}
