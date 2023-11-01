package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := &Person{"Alice", 18}
	p2 := &Person{"Bob", 20}
	p3 := &Person{"Alice", 18}

	// 比较结构体实例指针的内存地址
	if p1 == p2 {
		fmt.Println("p1 equals to p2")
	}

	// 比较结构体实例指针是否指向同一个实例
	if p1 == p3 {
		fmt.Println("p1 equals to p3")
	}
}
