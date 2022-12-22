package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
)

type (
	Bean struct {
		Id int
	}
)

func main() {
	s, err := sliceKit.DeepCopy([]string(nil))
	fmt.Println(s, err)

	//var c interface{} = nil

	//c1, err := copyKit.DeepCopy(interface{}(nil))
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(c1)

	//b := &Bean{
	//	Id: 666,
	//}
	//
	//s := []*Bean{b}
	//s1, err := copyKit.DeepCopy(s)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%+v\n", s)
	//fmt.Printf("%+v\n", s1)
	//
	//fmt.Println(jsonKit.MarshalToString(s))
	//fmt.Println(jsonKit.MarshalToString(s1))
	//s[0].Id = 999
	//fmt.Println(jsonKit.MarshalToString(s))
	//fmt.Println(jsonKit.MarshalToString(s1))

	//b1, err := copyKit.DeepCopy(b)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("%+v\n", b)
	//fmt.Printf("%+v\n", b1)
	//
	//u.Name = "李四"
	//fmt.Println(jsonKit.MarshalToString(b))
	//fmt.Println(jsonKit.MarshalToString(b1))
}
