package imageKit

import (
	"github.com/h2non/bimg"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/mapKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"sync"
)

var bimgOnce = new(sync.Once)
var bimgMapper map[string]bimg.ImageType

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
	bimgOnce.Do(func() {
		bimgMapper = mapKit.Invert(bimg.ImageTypes)
		bimgMapper["jpg"] = bimg.JPEG
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
	imageType, ok := bimgMapper[extName]
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
