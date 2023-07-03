package imageKit

import (
	"github.com/h2non/bimg"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/mapKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"image/jpeg"
	"image/png"
	"os"
	"sync"
)

var formatOnce = new(sync.Once)
var formatMapper map[string]bimg.ImageType

// Convert 转换图片的格式.
/*
!!!:
(1) 因为 h2non/bimg 基于C语言的libvip库，因此使用要满足"一些条件"，详见: https://mp.weixin.qq.com/s/kAFZohzJo2DiKkxjnVti6A
(2) bug: 转换后，透明背景色 可能=> 黑色背景色（即使目标格式支持透明背景色）；
(3) bug: 图片转pdf.

@param dest 如果已经存在且是个文件，会覆盖

支持的格式:
	"jpg"
	"jpeg"
	"png"
	"webp"
	"tiff"
	"gif"
	"pdf"
	"svg"
	"magick"
	"heif"
	"avif"
*/
func Convert(src, dest string) error {
	formatOnce.Do(func() {
		formatMapper = mapKit.Invert(bimg.ImageTypes)
		formatMapper["jpg"] = bimg.JPEG
	})

	// src
	if err := fileKit.AssertExistAndIsFile(src); err != nil {
		return err
	}
	// dest
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}
	extName := fileKit.GetExtName(dest)
	if err := strKit.AssertNotEmpty(extName, "extName"); err != nil {
		return err
	}
	extName = strKit.ToLower(extName)
	imageType, ok := formatMapper[extName]
	if !ok {
		return errorKit.New("extName(%s) of dest is invalid", extName)
	}
	if !bimg.IsTypeSupportedSave(imageType) {
		return errorKit.New("imageType(%d, %s) isn't supported to save by current libvips compilation",
			imageType, mapKit.Get(bimg.ImageTypes, imageType))
	}

	data, err := bimg.Read(src)
	if err != nil {
		return err
	}
	img := bimg.NewImage(data)
	data1, err := img.Convert(imageType)
	if err != nil {
		return err
	}
	return bimg.Write(dest, data1)
}

//import (
//	"github.com/disintegration/imaging"
//	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
//)
//
//// ConvertImageType 图片格式转换（图片类型转换）
///*
//@param src	源图片路径（如果文件不存在，会报error）
//@param dest	目标图片路径（如果文件已存在，会覆盖，覆盖不了就报error）
//
//PS:
//(1) src和dest的图片格式可以是一样的，此种情况类似于复制；
//(2) 支持的图片格式："jpg"、"jpeg"、"png"、"gif"、"tif"、"tiff"、"bmp".（详见 imaging.Save()）
//*/
//func ConvertImageType(src, dest string) error {
//	if err := fileKit.AssertExistAndIsFile(src); err != nil {
//		return err
//	}
//	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
//		return err
//	}
//	if err := fileKit.MkParentDirs(dest); err != nil {
//		return err
//	}
//
//	image, err := imaging.Open(src)
//	if err != nil {
//		return err
//	}
//	return imaging.Save(image, dest)
//}

// ToJpeg 将图片格式转换为".jpg" || ".jpeg"
func ToJpeg(src, dest string) error {
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(dest); err != nil {
		return err
	}

	srcImage, _, err := DecodeWithPath(src)
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()
	return jpeg.Encode(destFile, srcImage, &jpeg.Options{Quality: 100})
}

// ToPng 将图片格式转换为".png"
func ToPng(src, dest string) error {
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(dest); err != nil {
		return err
	}

	srcImage, _, err := DecodeWithPath(src)
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()
	return png.Encode(destFile, srcImage)
}
