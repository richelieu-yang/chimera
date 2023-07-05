package listKit

import "github.com/gogf/gf/v2/container/glist"

// NewDoubleLinkedList 创建（带并发安全开关的）双向链表.
/*
PS: 返回值的使用可以参考 "GoFrame.wps".
*/
var NewDoubleLinkedList func(safe ...bool) *glist.List = glist.New

// NewDoubleLinkedListFrom 创建（带并发安全开关的）双向链表.
/*
PS: 返回值的使用可以参考 "GoFrame.wps".

@param array 可以为nil，但这么干无意义
*/
var NewDoubleLinkedListFrom func(s []interface{}, safe ...bool) *glist.List = glist.NewFrom
