package main

import (
	"fmt"
	"github.com/samber/lo"
)

func main() {
	m1 := lo.Invert(map[string]int{"a": 1, "b": 2})
	fmt.Println(m1) // map[1:a 2:b]
	m2 := lo.Invert(map[string]int{"a": 1, "b": 2, "c": 1})
	fmt.Println(m2) // map[1:c 2:b] 或 map[1:a 2:b]（因为map是无序的）
}
