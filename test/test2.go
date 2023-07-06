package main

import (
	"container/list"
	"fmt"
	"github.com/gogf/gf/v2/container/glist"
)

func main() {
	l := glist.New(true)

	l.PushBack(0)
	l.PushBack(1)
	l.PushBack(2)

	l.LockFunc(func(list *list.List) {
		if e := list.Back(); e != nil {
			value = list.Remove(e)
		}
	})

	fmt.Println(l.Back().Value)
	fmt.Println(l.Back().Value)
	fmt.Println(l.Back().Value)
}
