package main

import (
	"fmt"
	"reflect"
)

type MyInt int

type Bean struct {
}

func main() {
	var i int
	var j MyInt
	i = int(j) // 必须强转

	typeI := reflect.TypeOf(i)
	fmt.Println("type of i:", typeI.String()) // type of i: int
	fmt.Println("kind of i:", typeI.Kind())   // kind of i: int

	typeJ := reflect.TypeOf(j)
	fmt.Println("type of j:", typeJ.String()) // type of j: main.MyInt
	fmt.Println("kind of j:", typeJ.Kind())   // kind of j: int

	b := Bean{}
	tmp := reflect.TypeOf(b)
	fmt.Println("type of b:", tmp.String()) // type of b: main.Bean
	fmt.Println("kind of b:", tmp.Kind())   // kind of b: struct

	b1 := &Bean{}
	tmp1 := reflect.TypeOf(b1)
	fmt.Println("type of b1:", tmp1.String()) // type of b1: *main.Bean
	fmt.Println("kind of b1:", tmp1.Kind())   // kind of b1: ptr
}
