package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
)

func main() {
	rw := ioKit.NewReadWriter(nil)
	//fmt.Println(rw.WriteString("0\n"))
	//fmt.Println(rw.WriteString("1\n"))
	str := rw.String()
	fmt.Println(str)
}
