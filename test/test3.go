package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	test("a.png", "a.jpg")
}

func test(src, dest string) {
	pngImgFile, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer pngImgFile.Close()
	jpgImgFile, err := os.Create(dest)
	if err != nil {
		panic(err)
	}
	defer jpgImgFile.Close()

	imgSrc, err := png.Decode(pngImgFile)
	if err != nil {
		panic(err)
	}
	newImg := image.NewRGBA(imgSrc.Bounds())
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg, newImg.Bounds(), imgSrc, imgSrc.Bounds().Min, draw.Over)
	err = jpeg.Encode(jpgImgFile, newImg, &jpeg.Options{
		Quality: 100,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("=========")
}
