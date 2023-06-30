package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aaa0aaa0aaa"

	fmt.Println(strings.Trim(s, "a"))      // 0123
	fmt.Println(strings.TrimLeft(s, "a"))  // 0123aaa
	fmt.Println(strings.TrimRight(s, "a")) // aaa0123
}
