package ptrKit

import (
	"github.com/duke-git/lancet/v2/pointer"
)

// Of
func Of[T any](v T) *T {
	return pointer.Of(v)
}

// Unwrap
func Unwrap[T any](p *T) T {
	return pointer.Unwrap(p)
}

// UnwarpOr
func UnwarpOr[T any](p *T, fallback T) T {
	return pointer.UnwarpOr(p, fallback)
}

// UnwarpOrDefault
func UnwarpOrDefault[T any](p *T) T {
	return pointer.UnwarpOrDefault(p)
}

// ExtractPointer
func ExtractPointer(value any) any {
	return pointer.ExtractPointer(value)
}

//// ToPtr Returns a pointer copy of value.
///*
//@param x 可以为nil（但需要指定类型T，且这么干貌似没意义）
//
//e.g.
//	ptr := ptrKit.ToPtr("hello world")
//	fmt.Println(ptr) 	// 0x140000105e0
//e.g.1 x == nil的情况
//	ptr := ptrKit.ToPtr[interface{}](nil)
//	fmt.Println(ptr)                     // 0x1400010c5c0
//	fmt.Println(ptr == nil)              // false
//	fmt.Println(interfaceKit.IsNil(ptr)) // false
//*/
//func ToPtr[T any](x T) *T {
//	lo.IsNil()
//
//	return lo.ToPtr(x)
//}
//
//// FromPtr 返回指针的值（如果指针为nil则返回T类型的零值）.
///*
//@param x 可以为nil（此时将返回类型T的zero value）
//@return 可能为nil
//
//e.g.
//	str := "测试test"
//	value := ptrKit.FromPtr(&str)
//	fmt.Println(value) // "测试test"
//	value = ptrKit.FromPtr[string](nil)
//	fmt.Println(value) // ""
//*/
//func FromPtr[T any](x *T) T {
//	return lo.FromPtr(x)
//}
//
//// FromPtrOr 返回指针的值（如果指针为nil则返回传参fallback）.
///*
//@param x 可以为nil（此时将返回传参fallback）
//*/
//func FromPtrOr[T any](x *T, fallback T) T {
//	return lo.FromPtrOr(x, fallback)
//}
