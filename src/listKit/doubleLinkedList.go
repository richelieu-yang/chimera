package listKit

import "github.com/gogf/gf/v2/container/glist"

// NewDoubleLinkedList 创建（带并发安全开关的）双向链表
/*
PS: 返回值的使用可以参考 "GoFrame.wps".
*/
func NewDoubleLinkedList(safe ...bool) *glist.List {
	return glist.New(safe...)
}

// NewDoubleLinkedListFrom 创建（带并发安全开关的）双向链表
/*
PS: 返回值的使用可以参考 "GoFrame.wps".
*/
func NewDoubleLinkedListFrom(array []interface{}, safe ...bool) *glist.List {
	return glist.NewFrom(array, safe...)
}
