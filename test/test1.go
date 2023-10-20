package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/reflectKit"
	"reflect"
)

func main() {
	{
		var obj interface{} = nil
		v := reflectKit.ValueOf(obj)
		k := v.Kind()
		fmt.Println(k)                    // invalid
		fmt.Println(k == reflect.Invalid) // true
	}

	{
		type bean struct {
		}

		var obj *bean = nil
		v := reflectKit.ValueOf(obj)
		k := v.Kind()
		fmt.Println(k)                    // ptr
		fmt.Println(k == reflect.Invalid) // false
	}
}
