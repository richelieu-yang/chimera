package main

import (
	"github.com/richelieu42/chimera/v2/src/compressKit"
)

func main() {
	err := compressKit.ZipPath("/Users/richelieu/Downloads/a111", "/Users/richelieu/Downloads/a.zip")
	if err != nil {
		panic(err)
	}
}
