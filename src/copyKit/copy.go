package copyKit

import (
	"github.com/jinzhu/copier"
	"github.com/mohae/deepcopy"
	"github.com/richelieu42/go-scales/src/core/errorKit"
)

// DeepCopyStruct 深拷贝结构体
/*
PS: 不支持map!!!

@param toValue 		必须是指针类型
@param fromValue 	指针类型 || 结构体实例
*/
func DeepCopyStruct(toValue interface{}, fromValue interface{}) error {
	return copier.Copy(toValue, fromValue)
}

// DeepCopy 深拷贝，支持：map、结构体...
/*
@param src 指针 || 结构体实例
*/
func DeepCopy[T any](src T) (dest T, err error) {
	obj := deepcopy.Copy(src)

	if t, ok := obj.(T); ok {
		dest = t
		return
	}
	err = errorKit.Simple("different types: src(%T) vs obj(%T)", src, obj)
	return
}
