package conditionKit

import "github.com/duke-git/lancet/v2/condition"

// Bool 返回传入参数的bool值.
/*
	如果出入类型参数含有Bool方法, 会调用该方法并返回
	如果传入类型参数有IsZero方法, 返回IsZero方法返回值的取反
	slices和map的length大于0时，返回true，否则返回false
	其他类型会判断是否是零值
*/
func Bool[T any](value T) bool {
	return condition.Bool(value)
}

// And 逻辑且操作，当切仅当a和b都为true时返回true
func And[T, U any](a T, b U) bool {
	return condition.And(a, b)
}

// Or 逻辑或操作，当切仅当a和b都为false时返回false
func Or[T, U any](a T, b U) bool {
	return condition.Or(a, b)
}

// Xor 逻辑异或操作，a和b相同返回false，a和b不相同返回true
func Xor[T, U any](a T, b U) bool {
	return condition.Xor(a, b)
}

// Nor 异或的取反操作
func Nor[T, U any](a T, b U) bool {
	return condition.Nor(a, b)
}

// Xnor 如果a和b都是真的或a和b均是假的，则返回true
func Xnor[T, U any](a T, b U) bool {
	return condition.Xnor(a, b)
}

// Nand 如果a和b都为真，返回false，否则返回true
func Nand[T, U any](a T, b U) bool {
	return condition.Nand(a, b)
}

// TernaryOperator 三目运算符（三元运算符）
/*
!!!: 传参 ifValue 和 elseValue，如果涉及复杂运算（调用方法或函数...），还是老老实实使用 if else 吧.
*/
func TernaryOperator[T, U any](isTrue T, ifValue U, elseValue U) U {
	return condition.TernaryOperator(isTrue, ifValue, elseValue)
}
