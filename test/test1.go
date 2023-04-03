package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/sliceKit"
)

func main() {
	s := []int{0, 1, 2, 3}
	s1, item, ok := sliceKit.RemoveByIndex(s, 2)

	fmt.Println(s)    // [0 1 2 3]
	fmt.Println(s1)   // [0 1 3]
	fmt.Println(item) // 2
	fmt.Println(ok)   // true
}

//func RemoveByIndex[T any](s []T, index int) (s1 []T, item T, ok bool) {
//	if len(s) == 0 {
//		s1 = s
//		return
//	}
//
//	item = s[index]
//
//	tmp := s[:index]
//	tmp1 := s[index+1:]
//
//	// !!!: 下面一行代码执行后，会修改外部的slice
//	s1 = append(tmp, tmp1...)
//
//	//s1 = append(s1, tmp...)
//	//s1 = append(s1, tmp1...)
//
//	ok = true
//	return
//}
