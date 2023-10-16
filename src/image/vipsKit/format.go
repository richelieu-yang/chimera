package vipsKit

import (
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"os"
)

func ToJpeg(src, dest string, exportParams *vips.JpegExportParams) error {
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}

	// 默认值
	if exportParams == nil {
		exportParams = vips.NewJpegExportParams()
		exportParams.Quality = 100
	}

	imageRef, err := Read(src, nil)
	if err != nil {
		return err
	}

	imageData, _, err := imageRef.ExportJpeg(exportParams)
	if err != nil {
		return err
	}
	return os.WriteFile(dest, imageData, 0644)
}

func ToPng(src, dest string, exportParams *vips.PngExportParams) error {
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}

	// 默认值
	if exportParams == nil {
		exportParams = vips.NewPngExportParams()
		exportParams.Quality = 100
	}

	imageRef, err := Read(src, nil)
	if err != nil {
		return err
	}

	imageData, _, err := imageRef.ExportPng(exportParams)
	if err != nil {
		return err
	}
	return os.WriteFile(dest, imageData, 0644)
}

func ToWebp(src, dest string, exportParams *vips.WebpExportParams) error {
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}

	// 默认值
	if exportParams == nil {
		exportParams = vips.NewWebpExportParams()
		exportParams.Quality = 100
	}

	imageRef, err := Read(src, nil)
	if err != nil {
		return err
	}

	imageData, _, err := imageRef.ExportWebp(exportParams)
	if err != nil {
		return err
	}
	return os.WriteFile(dest, imageData, 0644)
}

func ToTiff(src, dest string, exportParams *vips.TiffExportParams) error {
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}

	// 默认值
	if exportParams == nil {
		exportParams = vips.NewTiffExportParams()
		exportParams.Quality = 100
	}

	imageRef, err := Read(src, nil)
	if err != nil {
		return err
	}

	imageData, _, err := imageRef.ExportTiff(exportParams)
	if err != nil {
		return err
	}
	return os.WriteFile(dest, imageData, 0644)
}

func ToGif(src, dest string, exportParams *vips.GifExportParams) error {
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}

	// 默认值
	if exportParams == nil {
		exportParams = vips.NewGifExportParams()
		exportParams.Quality = 100
	}

	imageRef, err := Read(src, nil)
	if err != nil {
		return err
	}

	imageData, _, err := imageRef.ExportGIF(exportParams)
	if err != nil {
		return err
	}
	return os.WriteFile(dest, imageData, 0644)
}

func ToAvif(src, dest string, exportParams *vips.AvifExportParams) error {
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}

	// 默认值
	if exportParams == nil {
		exportParams = vips.NewAvifExportParams()
		exportParams.Quality = 100
	}

	imageRef, err := Read(src, nil)
	if err != nil {
		return err
	}

	imageData, _, err := imageRef.ExportAvif(exportParams)
	if err != nil {
		return err
	}
	return os.WriteFile(dest, imageData, 0644)
}

// ToHeif
/*
PS: iphone使用 heif格式 存储照片（占用内存比JPEG格式少）.
*/
func ToHeif(src, dest string, exportParams *vips.HeifExportParams) error {
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}

	// 默认值
	if exportParams == nil {
		exportParams = vips.NewHeifExportParams()
		exportParams.Quality = 100
	}

	imageRef, err := Read(src, nil)
	if err != nil {
		return err
	}

	imageData, _, err := imageRef.ExportHeif(exportParams)
	if err != nil {
		return err
	}
	return os.WriteFile(dest, imageData, 0644)
}
