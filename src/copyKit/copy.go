package copyKit

import (
	"github.com/gogf/gf/v2/util/gutil"
	"github.com/jinzhu/copier"
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/core/interfaceKit"
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
(1) github.com/mohae/deepcopy 虽然效果一样，但不推荐使用（star少; 最后更新时间2017）.

e.g.

*/
func DeepCopy[T any](src T) (dest T, err error) {
	if interfaceKit.IsNil(src) {
		return
	}

	obj := gutil.Copy(src)
	var ok bool
	if dest, ok = obj.(T); ok {
		return
	}
	err = errorKit.Simple("different types of src(%T) and dest(%T) are different", src, obj)
	return
}
