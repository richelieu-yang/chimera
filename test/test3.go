package main

import (
	"github.com/gogf/gf/v2/os/gfile"
	"os"
)

func main() {
	var f *os.File
	f.Truncate()

	gfile.Truncate()
}
