package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	str := strings.TrimFunc("77GeeksForGeeks!!!11a1", func(r rune) bool {
		return unicode.IsDigit(r)
	})
	fmt.Print(str) // "GeeksForGeeks!!!11a"
}
