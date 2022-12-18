package pointerKit

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"reflect"
)

// ToString 指针 => 指针的地址字符串（十六进制表示，前缀 0x）
/*
e.g.
tmp := 1
(&tmp) => "0xc00000a228"
e.g.1
(nil) => "%!p(<nil>)"
*/
func ToString(ptr interface{}) string {
	return fmt.Sprintf("%p", ptr)
}

// IsPointer
/*
e.g.
(nil) => false
*/
func IsPointer(obj interface{}) bool {
	// 参数obj为nil的话，变量v不为nil
	v := reflect.ValueOf(obj)
	return v.Kind() == reflect.Ptr
}

// IsPointer1
/**
 * 判断参数obj是否为指针（pointer）.
 * 参考：github.com/json-iterator/go => reflect.go => func (iter *Iterator) ReadVal(obj interface{})
 *
 * @param obj nil => false
 */
func IsPointer1(obj interface{}) bool {
	typ := reflect2.TypeOf(obj)
	return typ != nil && typ.Kind() == reflect.Ptr
}

// IsPointer2
/*
e.g.
(nil) => false
*/
func IsPointer2(obj interface{}) bool {
	str := fmt.Sprintf("%T", obj)
	return strKit.StartWith(str, "*")
}
