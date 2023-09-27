package main

import (
	"github.com/richelieu-yang/chimera/v2/src/image/vipsKit"
)

func main() {
	vipsKit.SetUp(nil)

	if err := vipsKit.ToWebp("iShot_2023-09-27_14.57.56.png", "1.webp", nil); err != nil {
		panic(err)
	}
	if err := vipsKit.ToWebp("iShot_2023-09-27_14.57.56.png", "2.webp", nil); err != nil {
		panic(err)
	}
}
