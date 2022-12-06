package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Getenv("user.home")可能会返回""，比如在Mac环境下
	tmp := os.Getenv("user.home")
	if tmp == "" {
		tmp = os.Getenv("HOME")
	}
	fmt.Println(tmp) // "/Users/richelieu"
}

//fmt.Println(userKit.GetName())
//fmt.Println(userKit.GetUserName())
//fmt.Println(userKit.GetUserHomeDir())
