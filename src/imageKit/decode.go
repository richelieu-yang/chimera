package imageKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"image"
	"io"
	"os"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	// 兼容webp格式解析
	_ "golang.org/x/image/webp"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// Decode 解码图片
/*
@param r 类型可以是: *os.File（用完记得调用Close()）
*/
var Decode func(r io.Reader) (image.Image, string, error) = image.Decode

// DecodeWithPath 解码图片
func DecodeWithPath(src string) (image.Image, string, error) {
	if err := fileKit.AssertExistAndIsFile(src); err != nil {
		return nil, "", err
	}

	f, err := os.Open(src)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()
	return Decode(f)
}
