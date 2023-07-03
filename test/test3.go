package main

import (
	"github.com/richelieu-yang/chimera/v2/src/imageKit"
)

func main() {
	if err := imageKit.Convert("1.jpg", "1.pdf"); err != nil {
		panic(err)
	}
}
