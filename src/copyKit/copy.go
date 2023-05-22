package copyKit

import (
	"github.com/gogf/gf/v2/util/gutil"
	"github.com/jinzhu/copier"
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/core/interfaceKit"
)

// Copy 浅拷贝
/*
@param dest	必须是指针类型
@param src 	指针类型 || 结构体实例

e.g.

*/
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
(1) unable to copy unexported fields in a struct (lowercase field names)
(2) 不使用 github.com/mohae/deepcopy: 	虽然效果一样，但不推荐使用（star少; 最后更新时间2017）
(3) 不使用 github.com/jinzhu/copier: 	深拷贝有bug，详见"Golang.wps"

@param src 可以为nil（此时将返回: nil, nil）

e.g. 传参为nil的情况
	a, err := copyKit.DeepCopy[interface{}](nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
	fmt.Println(a == nil) // true

	b, err := copyKit.DeepCopy[[]int](nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	fmt.Println(b == nil) // true

	c, err := copyKit.DeepCopy[map[string]interface{}](nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
	fmt.Println(c == nil) // true
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
	err = errorKit.Simple("types of src(%T) and dest(%T) are different")
	return
}
