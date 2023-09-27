package vipsKit

import (
	"github.com/davidbyttow/govips/v2/vips"
	"os"
)

func ToWebp(src, target string, params *vips.WebpExportParams) error {
	image, err := vips.NewImageFromFile(src)
	if err != nil {
		return err
	}

	// Rotate the picture upright and reset EXIF orientation tag
	err = image.AutoRotate()
	if err != nil {
		return err
	}

	exportParams := vips.NewWebpExportParams()
	webpData, _, err := image.ExportWebp(exportParams)
	if err != nil {
		return err
	}
	return os.WriteFile(target, webpData, 0644)
}
