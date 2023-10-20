package reflectKit

import "reflect"

// ValueOf
/*
@param i 可以为nil

e.g.
	v := reflectKit.ValueOf(111)
	fmt.Println(v.Kind() == reflect.Int)	// true

e.g.1
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
*/
var ValueOf func(i any) reflect.Value = reflect.ValueOf

func KindOf(i any) reflect.Kind {
	v := ValueOf(i)
	return v.Kind()
}
