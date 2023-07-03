package main

import (
	"github.com/h2non/bimg"
	"github.com/richelieu-yang/chimera/v2/src/imageKit"
)

func main() {
	if err := imageKit.Convert("c.webp", "c.png", bimg.JPEG); err != nil {
		panic(err)
	}
}
