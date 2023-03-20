package operationKit

// Ternary 三目运算符（ternary operator）
/*
PS:
(1) Golang原生不支持三目运算符，因此只能自己实现.
(2) ternary adj.三元的，三重的

参考：
golang三目运算的实现 https://www.cnblogs.com/GetcharZp/p/15172602.html

@param rst0 true 条件下的返回值
@param rst1 false条件下的返回值
*/
func Ternary[T any](flag bool, trueRst, falseRst T) T {
	if flag {
		return trueRst
	}
	return falseRst
}
