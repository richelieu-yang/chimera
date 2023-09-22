package mathKit

// Ternary 三目运算符（ternary operator）
/*
PS:
(1) Golang原生不支持三目运算符，因此只能自己实现
(2) ternary adj.三元的，三重的
(3) 简单的情况可以使用此函数，但复杂情况（比如第2、3个传参需要动态计算）请自行处理

参考：
golang三目运算的实现 https://www.cnblogs.com/GetcharZp/p/15172602.html

@param rst0 true条件下的返回值
@param rst1 false条件下的返回值
*/
func Ternary[T any](flag bool, trueRst, falseRst T) T {
	if flag {
		return trueRst
	}
	return falseRst
}
