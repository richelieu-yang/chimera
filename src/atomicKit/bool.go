package atomicKit

import "github.com/gogf/gf/v2/container/gtype"

// NewBool
/*
e.g. 不传参的话，初始值为false
	flag := atomicKit.NewBool()
	fmt.Println(flag.Val()) // false
*/
var NewBool func(value ...bool) *gtype.Bool = gtype.NewBool
