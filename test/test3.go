package main

import (
	"fmt"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Uid)
	fmt.Println(u.Gid)
	fmt.Println(u.Name)
	fmt.Println(u.Username)
}
