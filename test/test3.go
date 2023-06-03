package main

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
)

func main() {
	fmt.Println(gfile.Pwd())
	fmt.Println(gfile.SelfDir())
	fmt.Println(gfile.MainPkgPath())
}
