package main

import (
	"fmt"
	"github.com/samber/lo"
)

func main() {
	lo.ToPtr()

	m := map[string]int{}
	m["boy"] = 0
	// 就地更新
	m["boy"]++
	fmt.Println(m["boy"]) // 1
}
