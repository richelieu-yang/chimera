package imageKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"image/jpeg"
	"image/png"
	"os"
)

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
