package ringKit

import "github.com/gogf/gf/v2/container/gring"

var (
	// NewRing
	/*
		并发安全环-gring
			https://goframe.org/pages/viewpage.action?pageId=1114360

		gring.Ring: 支持并发安全开关的环结构，循环双向链表。
	*/
	NewRing func(cap int, safe ...bool) *gring.Ring = gring.New
)
