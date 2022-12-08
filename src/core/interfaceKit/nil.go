package interfaceKit

import (
	"reflect"
)

// IsNil
/*
Deprecated: 反射有性能问题，尽可能避免使用此方法.

PS:
(1) 如果某一变量的类型"不为指针"，那么可以直接用"== nil"来判空.
(2) 如果某一变量的类型"可能为指针"，那么需要用此方法来判空.
*/
func IsNil(obj interface{}) bool {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		// obj的类型：指针
		return v.IsNil()
	}
	// obj的类型：非指针
	return obj == nil
}
