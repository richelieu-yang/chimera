package imageKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"image/jpeg"
	"os"
)

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
