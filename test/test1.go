package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type (
	User struct {
		Id   int
		Name string
	}
)

func main() {
	u := &User{
		Id:   111,
		Name: "张三",
	}
	m := make(map[string]interface{})

	if err := mapstructure.Decode(u, m); err != nil {
		panic(err)
	}
	fmt.Println(m)
}
