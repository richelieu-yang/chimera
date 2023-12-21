package typeKit

import "fmt"

// GetTypeString 获取对象的类型字符串.
/*
@param 要注意nil的情况!!!
*/
func GetTypeString(obj interface{}) string {
	return fmt.Sprintf("%T", obj)
}
