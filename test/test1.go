package main

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/richelieu42/go-scales/src/core/mapKit"
	"unsafe"
)

func main() {
	tmp := &map[string]interface{}{
		"a": true,
	}
	src := map[string]interface{}{
		"b":   false,
		"tmp": tmp,
	}

	dest := mapKit.Clone(src)

	//dest, err := mapKit.DeepClone(src)
	//if err != nil {
	//	panic(err)
	//}

	//(*tmp)["a"] = "aaa"
	//fmt.Printf("%p\n", tmp)
	//fmt.Printf("%+v\n", src)
	//fmt.Printf("%+v\n", dest)

	{
		obj := src["tmp"]
		fmt.Println(unsafe.Pointer(obj.(*map[string]interface{})))
	}
	{
		obj := dest["tmp"]
		fmt.Println(unsafe.Pointer(obj.(*map[string]interface{})))
	}

	//fmt.Println(unsafe.Pointer(dest["tmp"]))
	//
	//fmt.Println(src)
	//fmt.Println(dest)
}

// Clone 浅拷贝
func Clone[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}

	dolly := make(map[K]V)
	for k, v := range m {
		dolly[k] = v
	}
	return dolly
}

func DeepClone[K comparable, V any](m map[K]V) (map[K]V, error) {
	dolly := make(map[K]V)
	if err := copier.CopyWithOption(&dolly, &m, copier.Option{
		DeepCopy: true,
	}); err != nil {
		return nil, err
	}
	return dolly, nil
}
