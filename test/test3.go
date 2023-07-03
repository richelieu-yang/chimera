package main

import (
	"github.com/h2non/bimg"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/imageKit"
)

func main() {
	fileKit.GetExt
	fileKit.GetExtName()

	strKit.EndWith

	if err := imageKit.Convert("a.webp", "a.pdf", bimg.PDF); err != nil {
		panic(err)
	}
}
