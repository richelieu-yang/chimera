package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/runeKit"
)

func main() {
	fmt.Println(runeKit.IsDigit('1')) // true
	fmt.Println(runeKit.IsDigit('0')) // true
	fmt.Println(runeKit.IsDigit('-')) // false
}
