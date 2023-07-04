package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/imageKit"
)

func main() {
	fmt.Println(imageKit.ToJpeg("/Users/richelieu/Desktop/a.png", "/Users/richelieu/Desktop/c1.jpg", 0))
}

//func test(src, dest string) {
//	pngImgFile, err := os.Open(src)
//	if err != nil {
//		panic(err)
//	}
//	defer pngImgFile.Close()
//	jpgImgFile, err := os.Create(dest)
//	if err != nil {
//		panic(err)
//	}
//	defer jpgImgFile.Close()
//
//	imgSrc, err := png.Decode(pngImgFile)
//	if err != nil {
//		panic(err)
//	}
//	newImg := image.NewRGBA(imgSrc.Bounds())
//	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
//	draw.Draw(newImg, newImg.Bounds(), imgSrc, imgSrc.Bounds().Min, draw.Over)
//	err = jpeg.Encode(jpgImgFile, newImg, &jpeg.Options{
//		Quality: 100,
//	})
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("=========")
//}
