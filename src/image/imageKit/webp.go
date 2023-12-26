package imageKit

import (
	"github.com/richelieu-yang/chimera/v2/src/file/fileKit"
	"golang.org/x/image/webp"
	"image/png"
	"os"
)

func ToWebp(src, dest string) error {
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}

	srcImage, _, err := DecodeWithPath(src)
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	webp.Encode

	return png.Encode(destFile, srcImage)
}
