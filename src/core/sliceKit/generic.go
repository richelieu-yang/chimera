package sliceKit

import (
	"math/rand"
	"time"
)

// Get 根据下标获取slice中的元素
/*
Deprecated: 考虑性能的场景下，不建议直接调用此方法（此方法仅供展示传参规范）.

PS:
(1) 如果s == nil，会导致panic；（不管index为何值，即使为0）
(2) 如果s != nil && len(s) == 0，会导致panic；（不管index为何值，即使为0）
(3) 如果s != nil && len(s) > 0，index的取值范围: [0, length).
*/
func Get[T any](s []T, index int) T {
	return s[index]
}

// Append 向slice实例的"最后面"添加元素
/*
Deprecated: 考虑性能的场景下，不建议直接调用此方法（此方法仅供展示传参规范）.

PS:
(1) 传参s == nil的情况下，此时如果eles数量>=1，将返回1个非nil的slice实例，否则将返回nil.
(2) append()返回的是1个新的slice实例.

@param s 可能为nil
*/
func Append[T any](s []T, eles ...T) []T {
	return append(s, eles...)
}

// Merge 合并多个切片（不会去重）
/*
PS:
(1) 先要传参nil的话，必须要造型. e.g. []string(nil)
(2) 第1个传参可以为nil

e.g.
() 				=> 编译报错：cannot infer T
(nil) 			=> 编译报错：cannot infer T

([]string(nil), []string{"1", "2"}) => [1 2]
([]string{"1", "2"}, []string(nil)) => [1 2]

([]string(nil))	=> nil
([]string{}) 	=> []
([]string(nil), []string{"a", "b"}, []string(nil), []string{"b", "c"}) => [a b b c]
*/
func Merge[T comparable](slices ...[]T) []T {
	var rst []T

	for _, s := range slices {
		if rst == nil {
			rst = s
		} else {
			rst = append(rst, s...)
		}
	}
	return rst
}

// RemoveDuplicate 切片实例去重.
/*
PS: 并不会改变传参s的内容（因为只是获取切片s的内容，并未修改）.

Golang数组去重&切片去重: https://www.cnblogs.com/enumx/p/12323081.html
通过map实现去重: https://mp.weixin.qq.com/s/tvy9L-pb_8WFWAmA9u-bMg

e.g.
[false true hello true false world] => [hello true false world]
*/
func RemoveDuplicate[T comparable](s []T) []T {
	length := len(s)
	if length <= 1 {
		return s
	}

	// 一个新的切片实例
	rst := make([]T, 0, length)
	for i := 0; i < length; i++ {
		repeat := false
		ele := s[i]

		for j := i + 1; j < length; j++ {
			if ele == s[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			rst = append(rst, ele)
		}
	}
	return rst
}

// Shuffle 随机打乱切片
/*
PS:
(1) 可能会修改传参s！！！
(2) 想要打乱元素，我们必须引入随机数，然后再交换元素；
(3) 小概率没打乱切片实例.

e.g.
[1 2 3 45 7 0 ccc true false] => [true 2 ccc 1 false 0 3 7 45]
*/
func Shuffle[T any](s []T) {
	length := len(s)
	if length > 1 {
		// seed random for changing order of elements
		random := rand.New(rand.NewSource(time.Now().UnixNano()))

		for i := length - 1; i > 0; i-- {
			// [0, i + 1)范围内的随机整数
			j := random.Intn(i + 1)
			// 交换值
			s[i], s[j] = s[j], s[i]
		}
	}
}

// Reverse 反转切片
/*
PS: 可能会修改传参s！！！
参考: https://mp.weixin.qq.com/s/tvy9L-pb_8WFWAmA9u-bMg

e.g.
[] 	=> []
nil => nil
*/
func Reverse[T any](s []T) []T {
	for i := len(s)/2 - 1; i >= 0; i-- {
		pos := len(s) - 1 - i
		s[i], s[pos] = s[pos], s[i]
	}
	return s
}

// Swap 交换切片实例中两个元素的值
/*
PS:
(1) 下标越界会导致panic!!!
(2) 此方法会修改传参s（虽然golang是值传递）.

@param i 第1个元素的下标（从0开始）
@param j 第2个元素的下标（从0开始）

e.g.
s := []int{0, 1, 2, 3}
fmt.Println(s)	// [0 1 2 3]
sliceKit.Swap(s, 1, 2)
fmt.Println(s)	// [0 2 1 3]
*/
func Swap[T any](s []T, i, j int) {
	s[i], s[j] = s[j], s[i]
}

// GetFirstItemWithDefault 主要用于: 从不定参数(...)中取第一个值（不存在则取默认值）
/*
PS:
(1) 因为Golang不支持方法重载；
(2) T类型值可能为nil的情况，要注意防坑.

@param args 要么是: nil；要么是: 长度>=1的切片实例
*/
func GetFirstItemWithDefault[T any](def T, args ...T) T {
	if args != nil {
		return args[0]
	}
	return def
}

func EmptyToNil[T any](s []T) []T {
	if IsEmpty(s) {
		return nil
	}
	return s
}

// IsEmpty
/*
@param s 可以为nil
*/
func IsEmpty[T any](s []T) bool {
	return len(s) == 0
}

// IsNotEmpty
/*
@param s 可以为nil
*/
func IsNotEmpty[T any](s []T) bool {
	return len(s) > 0
}

// Contains 切片s是否包含元素t？（区分大小写，因为使用"=="比较）
/*
@param s 如果为nil，返回值必定为 false
*/
func Contains[T comparable](s []T, t T) bool {
	for _, ele := range s {
		if t == ele {
			return true
		}
	}
	return false
}
