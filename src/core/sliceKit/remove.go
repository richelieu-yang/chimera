package sliceKit

// RemoveByIndex 根据下标删除指定元素.
/*
Deprecated: 考虑性能的场景下，不建议直接调用此方法（此方法仅供展示传参规范）.

PS:
(0) 也可以用于移除"第一个"或"最后一个"个元素；
(1) 为什么要有两个返回值且第一个是slice？因为 append() 并不会修改传参slice，返回的是一个新的slice.
(2) 调用此方法后，必须第一时间将 第一个返回值 赋值给 传进来的那个slice变量！！！

@param s		可以为nil
@param index 	下标（索引），取值范围: [0, length)
@return 移除后的slice + 被移除的元素 + 是否成功移除？（主要是针对nil的情况）
*/
func RemoveByIndex[T any](s []T, index int) (s1 []T, item T, ok bool) {
	if len(s) == 0 {
		s1 = s
		return
	}

	item = s[index]
	// 下面一行代码执行后，会修改外部的slice！！！
	s1 = append(s[:index], s[index+1:]...)
	ok = true
	return
}

// RemoveFirst 移除第一个元素（如果有的话）
/*
@param s 可以为nil

e.g.
[int](nil) 		=> [] 0 false
([]int{0}) 		=> [] 0 true
([]int{0, 1}) 	=> [1] 0 true
*/
func RemoveFirst[T any](s []T) (s1 []T, item T, ok bool) {
	if len(s) == 0 {
		s1 = s
		return
	}

	// 此时第一个返回值 必定不为nil && len >= 0
	item = s[0]
	s1 = s[1:]
	ok = true
	return
}

// RemoveLast 移除最后一个元素（如果有的话）
/*
@param s 可以为nil

e.g.
[int](nil) 		=> [] 0 false
([]int{0}) 		=> [] 0 true
([]int{0, 1}) 	=> [0] 1 true
*/
func RemoveLast[T any](s []T) (s1 []T, item T, ok bool) {
	if len(s) == 0 {
		s1 = s
		return
	}

	// 此时第一个返回值 必定不为nil && len >= 0
	index := len(s) - 1
	item = s[index]
	s1 = s[0:index]
	ok = true
	return
}

// Remove 移除元素
/*
PS:
(1) 不会修改传参s.
(2) 切片实例s中，存在多个item的话，仅会移除第一个.
*/
func Remove[T comparable](s []T, item T) ([]T, bool) {
	index := Index(s, item)
	if index == -1 {
		// 如果元素不存在于切片中，返回原始切片
		return s, false
	}
	return append(s[:index], s[index+1:]...), true
}

//// Remove 移除元素
///*
//@param s			可以为nil
//@param item			可以为nil
//@param entireArgs 	slice中存在多个传参item的情况下，是否全部移除？默认: false
//@return 第1个返回值: 移除后的slice（也有可能没变）; 第2个返回值: 是否移除了元素？
//
//e.g.
//	texts := []string{"0", "1", "2"}
//	fmt.Println(texts) // [0 1 2]
//
//	texts1, _ := sliceKit.Remove(texts, "1")
//	fmt.Println(texts)  // [0 2 2]
//	fmt.Println(texts1) // [0 2]
//*/
//func Remove[T comparable](s []T, item T, entireArgs ...bool) ([]T, bool) {
//	DeepCopy()
//
//	var judge RemoveJudge[T]
//
//	// 默认: false
//	entire := GetFirstItemWithDefault(false, entireArgs...)
//	if entire {
//		// 移除所有item（slice允许内部元素重复）
//		judge = func(ele T) (removeFlag bool, interrupt bool) {
//			removeFlag = ele == item
//			interrupt = false
//			return
//		}
//	} else {
//		// 只移除第1个item
//		judge = func(ele T) (removeFlag bool, interrupt bool) {
//			removeFlag = ele == item
//			interrupt = removeFlag
//			return
//		}
//	}
//	return RemoveByCondition(s, judge)
//}
//
//// RemoveByCondition
///*
//@param s 		可以为nil
//@param judge 	可以为nil，此时将不删除任何元素
//*/
//func RemoveByCondition[T comparable](s []T, judge RemoveJudge[T]) ([]T, bool) {
//	removeFlag := false
//
//	if judge != nil {
//		for i := 0; i < len(s); {
//			ele := s[i]
//
//			remove, interrupt := judge(ele)
//			if remove {
//				// 删除s中下标为i的元素
//				s = append(s[:i], s[i+1:]...)
//				removeFlag = true
//				if interrupt {
//					break
//				}
//			} else {
//				if interrupt {
//					break
//				}
//				i++
//			}
//		}
//	}
//	return s, removeFlag
//}
