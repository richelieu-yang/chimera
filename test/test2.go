package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/consts"
)

func main() {
	timePattern := "(%Y-%m-%d %H" + consts.ColonInFileName + "%M" + consts.ColonInFileName + "%S)"
	fmt.Println(timePattern)
}
