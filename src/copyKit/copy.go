package copyKit

import (
	"github.com/gogf/gf/v2/util/gutil"
	"github.com/jinzhu/copier"
)

// Copy 浅拷贝
func Copy(dest, src interface{}) error {
	return copier.CopyWithOption(dest, src, copier.Option{
		IgnoreEmpty: false,
		DeepCopy:    false,
		Converters:  nil,
	})
}

// DeepCopy 深拷贝
/*
PS:
(1) github.com/mohae/deepcopy 虽然效果一样，但不推荐使用（star少; 最后更新时间2017）
(2) unable to copy unexported fields in a struct (lowercase field names)

@param src 可以为nil（此时将返回nil）
@return The returned value will need to be asserted to the correct type.（需要手动断言）

e.g.

*/
func DeepCopy(src interface{}) interface{} {
	return gutil.Copy(src)
}
