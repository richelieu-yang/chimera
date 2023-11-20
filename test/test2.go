package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/cronKit"
)

func main() {
	c := cronKit.NewCron()

	cronKit.StopCron(c)
	fmt.Println("-")
	cronKit.StopCron(c)
	fmt.Println("-")
}
