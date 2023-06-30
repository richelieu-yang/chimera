package main

import "github.com/richelieu-yang/chimera/v2/src/imageKit"

func main() {
	if err := imageKit.ToJpeg("a.png", "b.jpg"); err != nil {
		panic(err)
	}
	if err := imageKit.ToJpeg("a.png", "b.jpeg"); err != nil {
		panic(err)
	}
	if err := imageKit.ToPng("b.jpg", "c.png"); err != nil {
		panic(err)
	}
}
