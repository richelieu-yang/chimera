package vipsKit

import (
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"os"
)

func ToWebp(src, dest string, exportParams *vips.WebpExportParams) error {
	if err := fileKit.AssertNotExistOrIsFile(dest, true); err != nil {
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
