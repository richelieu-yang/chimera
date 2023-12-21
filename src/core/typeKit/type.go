package typeKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/reflectKit"
)

// GetTypeString 获取对象的类型字符串.
/*
@param 要注意nil的情况!!!
*/
func GetTypeString[T any](obj T) string {
	reflectKit.IsNil(obj)

	if obj == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%T", obj)
}
