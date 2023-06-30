package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/runeKit"
)

func main() {
	fmt.Println(runeKit.IsSpace(' '))  // true
	fmt.Println(runeKit.IsSpace('\r')) // true
	fmt.Println(runeKit.IsSpace('\n')) // true
	fmt.Println(runeKit.IsSpace('\t')) // true
}
