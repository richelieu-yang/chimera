package operationKit

// Ternary 三目运算符（ternary operator）
/*
PS:
(1) Golang原生不支持三目运算符，因此只能自己实现.
(2) ternary adj.三元的，三重的
*/
func Ternary[T any](flag bool, rst, rst1 T) T {
	if flag {
		return rst
	}
	return rst1
}
