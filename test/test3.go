package main

import (
	"github.com/h2non/bimg"
)

func main() {
	if err := Convert("/Users/richelieu/Desktop/透明背景色.webp", "a.jpg", bimg.JPEG); err != nil {
		panic(err)
	}
}

// Convert
/*
!!!: 因为 h2non/bimg 基于C语言的libvip库，因此使用要满足"一些条件"，详见:
	「GoCN酷Go推荐」Go 语言高性能图像处理神器 h2non/bimg https://mp.weixin.qq.com/s/kAFZohzJo2DiKkxjnVti6A
*/
func Convert(src, dest string, imageType bimg.ImageType) error {
	buffer, err := bimg.Read(src)
	if err != nil {
		return err
	}
	newImage, err := bimg.NewImage(buffer).Convert(imageType)
	if err != nil {
		return err
	}
	//return bimg.Write("004-Convert."+bimg.ImageTypes[imageType], newImage)
	return bimg.Write(dest, newImage)
}
