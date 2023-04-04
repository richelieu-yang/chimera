package ptrKit

import "github.com/samber/lo"

// ToPtr Returns a pointer copy of value.
/*
@param x 可以为nil（但需要指定类型T）

e.g.
	ptr := ptrKit.ToPtr("hello world")
	fmt.Println(ptr) // 0x140000105e0
*/
func ToPtr[T any](x T) *T {
	return lo.ToPtr(x)
}

// FromPtr 返回指针的值（如果指针为nil则返回T类型的零值）.
/*
@param x 可以为nil（此时将返回类型T的zero value）
@return 可能为nil

e.g.
	str := "测试test"
	value := ptrKit.FromPtr(&str)
	fmt.Println(value) // "测试test"
	value = ptrKit.FromPtr[string](nil)
	fmt.Println(value) // ""
*/
func FromPtr[T any](x *T) T {
	return lo.FromPtr(x)
}

// FromPtrOr 返回指针的值（如果指针为nil则返回传参fallback）.
/*
@param x 可以为nil（此时将返回传参fallback）
*/
func FromPtrOr[T any](x *T, fallback T) T {
	return lo.FromPtrOr(x, fallback)
}
