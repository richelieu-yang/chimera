package randomKit

import "github.com/gogf/gf/v2/util/grand"

// Bytes 返回指定长度的二进制[]byte数据
/*
e.g.
	data := randomKit.Bytes(20)
	fmt.Println(data) // [193 72 130 210 86 72 180 144 91 246 162 215 176 4 169 153 130 140 89 15]
*/
var Bytes func(n int) []byte = grand.B
