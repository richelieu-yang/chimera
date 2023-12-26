package ptrKit

import (
	"github.com/duke-git/lancet/v2/pointer"
)

// Of 返回传入参数的指针值.
/*
e.g.
	var ptr *string = ptrKit.Of("123")
	fmt.Println(ptr) // 0x14000110360
*/
func Of[T any](v T) *T {
	return pointer.Of(v)
}

// Unwrap 返回指针的值.
/*
e.g.
	var ptr *string = ptrKit.Of("123")
	fmt.Println(ptr) // 0x140000103b0

	fmt.Println(ptrKit.Unwrap(ptr)) // 123
*/
func Unwrap[T any](p *T) T {
	return pointer.Unwrap(p)
}

// UnwarpOr 返回指针的值（如果指针为零值，则返回fallback）.
/*
e.g.
	var ptr *int = nil
	fmt.Println(ptrKit.UnwarpOr(ptr, 123)) // 123
*/
func UnwarpOr[T any](p *T, fallback T) T {
	return pointer.UnwarpOr(p, fallback)
}

// UnwarpOrDefault 返回指针的值（如果指针为零值，则返回相应零值）.
/*
e.g.
	var ptr *int = nil
	fmt.Println(ptrKit.UnwarpOrDefault(ptr)) // 0
*/
func UnwarpOrDefault[T any](p *T) T {
	return pointer.UnwarpOrDefault(p)
}

// ExtractPointer 返回传入interface的底层值.
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
