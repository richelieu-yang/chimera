package main

import (
	"bufio"
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
)

func main() {
	i := 0

	fileKit.ReadFileByLine("a.txt", func(scan *bufio.Scanner) {
		i++
		fmt.Println(i, scan.Text())
	})
}
