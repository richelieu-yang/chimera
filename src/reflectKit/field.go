package reflectKit

import "reflect"

// GetField
/*
PS:
这样获取到的值是不可以修改的，如果要修改name字段，需要用到reflect.NewAt函数，这个函数通过一个类型值的底层地址（指针 p）和类型，
返回指向该值的一个指针，这个返回值是可寻址的，即可通过它直接访问该值。具体的修改是使用reflect.Value.Set函数。

@param fieldName 字段名（属性名），private 或 public的都可以
*/
func GetField(ptr interface{}, fieldName string) reflect.Value {
	return reflect.ValueOf(ptr).Elem().FieldByName(fieldName)
}
