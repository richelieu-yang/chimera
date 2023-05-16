package main

import (
	"fmt"
	"sort"
)

func main() {
	var s sort.Float64Slice = nil
	sort.Stable(s)
	fmt.Println(s == nil)
}
