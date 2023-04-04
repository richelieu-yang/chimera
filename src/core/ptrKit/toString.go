package ptrKit

import "fmt"

// ToString 指针 => 指针的地址字符串（十六进制表示，前缀 0x）
/*
e.g.
	tmp := 1
	(&tmp) => "0xc00000a228"
e.g.1
	(nil) => "%!p(<nil>)"
*/
func ToString(ptr interface{}) string {
	return fmt.Sprintf("%p", ptr)
}
