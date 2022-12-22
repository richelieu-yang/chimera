package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/copyKit"
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
	b := &Bean{
		Id: 666,
		U:  u,
	}
	b1, err := copyKit.DeepCopy(b)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", b)
	fmt.Printf("%+v\n", b1)

	u.Name = "李四"
	fmt.Println(jsonKit.MarshalToString(b))
	fmt.Println(jsonKit.MarshalToString(b1))
}
