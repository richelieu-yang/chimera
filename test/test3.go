package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"os"
)

func main() {
	//fmt.Printf("%s %v\n", os.ModePerm, os.ModePerm)

	var perm os.FileMode = 0666
	//fmt.Printf("%s %v\n", perm, perm)

	err := errorKit.New(`fail with  perm(%s)`, perm)
	fmt.Printf(err.Error())
}
