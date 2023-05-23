package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/crypto/md5Kit"
)

func main() {
	fmt.Println(md5Kit.EncryptFile("/Users/richelieu/Desktop/tasks.wps"))

	fmt.Println("d9bc2286000a70c293180252479d8075" == "d9bc2286000a70c293180252479d8075")
}
