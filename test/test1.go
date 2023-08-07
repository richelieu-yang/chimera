package main

import (
	"fmt"
	"os"
)

func main() {
	path := ""
	//path := "/aaa.log"
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	fmt.Println(info != nil)
}
