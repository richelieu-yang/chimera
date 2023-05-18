package main

import (
	"container/list"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/glist"
)

func main() {
	l := glist.NewFrom(garray.NewArrayRange(1, 9, 1).Slice(), true)
	fmt.Println(l) // [1,2,3,4,5,6,7,8,9]

	// iterate reading from head.
	l.RLockFunc(func(list *list.List) {
		length := list.Len()
		if length > 0 {
			for i, e := 0, list.Front(); i < length; i, e = i+1, e.Next() {
				fmt.Print(e.Value)
			}
		}
	})
	fmt.Println()

	// iterate reading from tail.
	l.RLockFunc(func(list *list.List) {
		length := list.Len()
		if length > 0 {
			for i, e := 0, list.Back(); i < length; i, e = i+1, e.Prev() {
				fmt.Print(e.Value)
			}
		}
	})

	fmt.Println()

	//// iterate reading from head using IteratorAsc.
	//l.IteratorAsc(func(e *glist.Element) bool {
	//	fmt.Print(e.Value)
	//	return true
	//})
	//fmt.Println()
	//
	//// iterate reading from tail using IteratorDesc.
	//l.IteratorDesc(func(e *glist.Element) bool {
	//	fmt.Print(e.Value)
	//	return true
	//})
	//fmt.Println()

	// iterate writing from head.
	l.LockFunc(func(list *list.List) {
		length := list.Len()
		if length > 0 {
			for i, e := 0, list.Front(); i < length; i, e = i+1, e.Next() {
				if e.Value == 6 {
					e.Value = "M"
					break
				}
			}
		}
	})
	fmt.Println(l)
}
