package main

import (
	"fmt"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"reflect"
	"unsafe"
)

func main() {
	type Bean struct {
		id uint
	}

	b := &Bean{
		id: 123,
	}
	addr := GetUnexportedFieldAddr(b, "id")

	// 转换回原始类型，这里假设我们确切知道原始类型是什么
	i := *(*uint)(addr)
	fmt.Println(i)
}

// GetUnexportedFieldAddr 获取未导出字段的地址
func GetUnexportedFieldAddr(ptr any, fieldName string) unsafe.Pointer {
	field := reflect.ValueOf(ptr).Elem().FieldByName(fieldName)

	//return field.UnsafePointer()
	/*
		获取未导出字段地址并转换为 uintptr，然后转回 unsafe.Pointer 类型
		这里假设字段是可寻址的并且类型是可以转换的
	*/
	return unsafe.Pointer(field.UnsafeAddr())
}

//func GetUnexportedFieldValue(i any, fieldName string) interface{} {
//	rv := reflect.ValueOf(i)
//	if rv.Kind() != reflect.Ptr {
//		panic("需要传入指针类型的值")
//	}
//	rv = rv.Elem()
//
//	field := rv.FieldByName(fieldName)
//	if !field.IsValid() {
//		panic("无效的字段名")
//	}
//
//	// 对于未导出字段，即使找到了，也无法直接访问其值，除非它实现了某些接口
//	// 在Go 1.17及以后版本，可以通过 unsafe 来读取未导出字段的值，但这违反了封装性且不安全
//	// 下面的代码仅用于演示目的，实际开发中应避免这样的操作：
//	if field.CanSet() { // CanSet 可能为 false，对于未导出字段通常是这样
//		fmt.Println("警告：违反了封装原则，不建议直接访问未导出字段！")
//		// 获取未导出字段地址并转换为 uintptr，然后转回 unsafe.Pointer 类型
//		// 这里假设字段是可寻址的并且类型是可以转换的
//		addr := unsafe.Pointer(field.UnsafeAddr())
//		// 转换回原始类型，这里假设我们确切知道原始类型是什么
//		fieldValue := *(*string)(addr)
//		return fieldValue
//	} else {
//		fmt.Println("警告：违反了封装原则，不建议直接访问未导出字段！")
//		// 获取未导出字段地址并转换为 uintptr，然后转回 unsafe.Pointer 类型
//		// 这里假设字段是可寻址的并且类型是可以转换的
//		addr := unsafe.Pointer(field.UnsafeAddr())
//		// 转换回原始类型，这里假设我们确切知道原始类型是什么
//		fieldValue := *(*string)(addr)
//		return fieldValue
//	}
//}
