package sliceKit

// Intercept 截取 [from, to)
/*
参考:
golang2021数据格式（23）切片截取 https://baijiahao.baidu.com/s?id=1711198159326157378

PS:
(1) 不存在越界 && s != nil（包括len(s) == 0） 的情况下，如果from == to，将返回空的slice实例（非nil）.
(2) from > to || 下标越界，会导致 panic: runtime error: slice bounds out of range [1:0]

@param s	可以为nil（将返回nil）
@param from	取值范围: [0, cap(s))]
@param to	取值范围: [0, cap(s))]
@return	(1) 可能为nil（当传参s为nil时）
		(2) 返回值不涉及深浅拷贝，修改会导致"同步修改"

e.g.	返回值不涉及深浅拷贝，修改会导致"同步修改"
   s := []int{0, 1, 2, 3}
   s1 := s[1:]

   fmt.Println(s, unsafe.Pointer(&s))   // [0 1 2 3] 0x140000a0018
   fmt.Println(s1, unsafe.Pointer(&s1)) // [1 2 3] 0x140000a0030

   s1[2] = 9
   fmt.Println(s, unsafe.Pointer(&s))   // [0 1 2 9] 0x140000a0018
   fmt.Println(s1, unsafe.Pointer(&s1)) // [1 2 9] 0x140000a0030

e.g.1	不存在越界的情况下，如果from == to，将返回空的slice实例（非nil）.
	s := []int{0, 1, 2}
	s1 := sliceKit.Intercept(s, len(s), len(s))
	fmt.Println(s1)        // []
	fmt.Println(s1 != nil) // true

e.g.2
	var s []int = []int{}
	s1 := sliceKit.Intercept(s, 0, 0)

	fmt.Println(s1)
	fmt.Println(s1 != nil)
*/
func Intercept[T any](s []T, from, to int, maxArgs ...int) []T {
	if len(s) == 0 {
		return s
	}
	if to < 0 {
		to = len(s)
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
@return 		(1) [0, index)
				(2) s为空，返回: s
				(3) s非空 && index == 0，返回: []（空）
				(4) s非空 && index == length，返回: s

e.g.
	s := []int{1}
	fmt.Println(sliceKit.InterceptBefore(s, 0)) // []
	fmt.Println(sliceKit.InterceptBefore(s, 1)) // [1]
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
@return 		(1) [index, length)
				(2) s为空，返回: s
				(3) s非空 && index == 0，返回: s
				(4) s非空 && index == length，返回: []（空）

e.g.
	s := []int{1}
	fmt.Println(sliceKit.InterceptAfter(s, 0)) // [1]
	fmt.Println(sliceKit.InterceptAfter(s, 1)) // []
*/
func InterceptAfter[T any](s []T, index int) []T {
	if len(s) == 0 {
		return s
	}

	// 此时返回值必定非nil，且len >= 0
	return s[index:]
}
