package reflectKit

import (
	"unsafe"
)

// GetUnexportedFieldAddr 获取未导出字段的地址.
func GetUnexportedFieldAddr(ptr any, fieldName string) unsafe.Pointer {
	field := GetField(ptr, fieldName)

	// (1)
	//return field.UnsafePointer()

	// (2) 获取未导出字段地址并转换为 uintptr，然后转回 unsafe.Pointer 类型（这里假设字段是可寻址的并且类型是可以转换的）
	return unsafe.Pointer(field.UnsafeAddr())
}
