package main

import (
	"cmp"
	"fmt"
)

func main() {
	fmt.Println(Max(0, 100, 1, 2))
}

func Max[T cmp.Ordered](x T, y ...T) T {
	return max[T](x, y...)
}
