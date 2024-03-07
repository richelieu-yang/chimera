package reflectKit

import (
	"unsafe"
)

// GetUnexportedFieldAddrOfBasicType 获取未导出字段的地址.
func GetUnexportedFieldAddrOfBasicType(ptr any, fieldName string) unsafe.Pointer {
	field := GetField(ptr, fieldName)

	// 获取未导出字段地址并转换为 uintptr，然后转回 unsafe.Pointer 类型（这里假设字段是可寻址的并且类型是可以转换的）
	return unsafe.Pointer(field.UnsafeAddr())
}

// GetUnexportedFieldAddr 获取未导出字段的地址.
/*
参考: gorm.io/gorm v1.25.7 中的 DB.DB()
*/
func GetUnexportedFieldAddr(ptr any, fieldName string) unsafe.Pointer {
	field := GetField(ptr, fieldName)

	// 部分字段类型不支持，比如: uint...
	return field.UnsafePointer()
}
