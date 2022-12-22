package copyKit

import (
	"github.com/jinzhu/copier"
	"github.com/mohae/deepcopy"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/interfaceKit"
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

// DeepCopy 深拷贝，支持: slice、map、结构体...
/*
PS:
参考: 「Go工具箱」推荐一个非常简单的深拷贝工具：deepcopy https://mp.weixin.qq.com/s/e3bL1i6WT-4MwK-SEpUa6Q；

@param src 指针 || 结构体实例
@return e.g. (interface{}(nil)) => (nil, nil)
*/
func DeepCopy[T any](src T) (dest T, err error) {
	// 处理nil的情况
	if interfaceKit.IsNil(src) {
		return
	}

	obj := deepcopy.Copy(src)
	if t, ok := obj.(T); ok {
		dest = t
		return
	}
	err = errorKit.Simple("different types: src(%T) vs obj(%T)", src, obj)
	return
}
