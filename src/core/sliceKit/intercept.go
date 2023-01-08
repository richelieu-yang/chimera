package sliceKit

// Intercept 截取切片实例
/*
参考:
golang2021数据格式（23）切片截取 https://baijiahao.baidu.com/s?id=1711198159326157378

PS:
(1) 如果from == to且不存在越界的情况，将返回len为0的slice实例（非nil）.

@param s	可以为nil
@param from	取值范围: [0, len(s))]
@param to	取值范围: [0, len(s))]
@return 	[from, to)

e.g.
s := []int{0, 1, 2, 3, 4, 5}
fmt.Println(sliceKit.Intercept(s, len(s), len(s))) // []（非nil）
*/
func Intercept[T any](s []T, from, to int, maxArgs ...int) []T {
	if len(s) == 0 {
		return s
	}

	// 此时返回值必定非nil，且len >= 0
	if maxArgs == nil {
		/* 情况1: 返回slice的cap采用默认值（cap = len(s) - from） */
		return s[from:to]
	}
	/* 情况2: 人为干预返回slice的cap（cap = max - from），适用场景: 想要减少内存消耗. */
	// max的理论取值范围: [to, len(s)]
	max := maxArgs[0]
	if max < to {
		max = to
	} else if max > len(s) {
		max = len(s)
	}
	return s[from:to:max]
}

// InterceptBefore
/*
@param s		可以为nil
@param index	取值范围: [0, length]
@return 		[0, index)
*/
func InterceptBefore[T any](s []T, index int) []T {
	if len(s) == 0 {
		return s
	}
	// 此时返回值必定非nil，且len >= 0
	return s[:index]
}

// InterceptAfter
/*
@param s		可以为nil
@param index	取值范围: [0, length]
@return 		[index, length)
*/
func InterceptAfter[T any](s []T, index int) []T {
	if len(s) == 0 {
		return s
	}
	// 此时返回值必定非nil，且len >= 0
	return s[index:]
}
