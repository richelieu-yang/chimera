package main

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"os"
)

func main() {
	gfile.Ext()
	gfile.ExtName()

	err := os.MkdirAll(" ", os.ModePerm)
	fmt.Println(err)
	fmt.Println(err == nil)
}
