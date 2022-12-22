package main

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/richelieu42/go-scales/src/jsonKit"
)

type (
	Bean struct {
		Id int
		U  *User
	}

	User struct {
		Name string
	}
)

func main() {
	u := &User{
		Name: "张三",
	}
	b := Bean{
		Id: 666,
		U:  u,
	}
	b1 := Bean{}

	if err := copier.CopyWithOption(&b1, b, copier.Option{
		DeepCopy:    false,
	}); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", b)
	fmt.Printf("%+v\n", b1)

	u.Name = "李四"
	fmt.Println(jsonKit.MarshalToString(b))
	fmt.Println(jsonKit.MarshalToString(b1))
}
