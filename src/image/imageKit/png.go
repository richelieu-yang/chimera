package imageKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"image/png"
	"os"
)

// ToPng 将图片格式转换为".png"
func ToPng(src, dest string) error {
	if err := fileKit.AssertNotExistOrIsFile(dest, true); err != nil {
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
