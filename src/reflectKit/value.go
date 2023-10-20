package reflectKit

import "reflect"

// ValueOf
/*
e.g.
	v := reflectKit.ValueOf(111)
	fmt.Println(v.Kind() == reflect.Int)	// true
*/
var ValueOf func(i any) reflect.Value = reflect.ValueOf

func KindOf(i any) reflect.Kind {
	v := ValueOf(i)
	return v.Kind()
}
