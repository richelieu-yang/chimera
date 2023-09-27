package main

import (
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	vips.Startup(nil)
	defer vips.Shutdown()

	logrus.Info("666")

	image1, err := vips.NewImageFromFile("iShot_2023-09-27_14.57.56.png")
	if err != nil {
		panic(err)
	}

	// Rotate the picture upright and reset EXIF orientation tag
	err = image1.AutoRotate()
	if err != nil {
		panic(err)
	}

	exportParams := vips.NewWebpExportParams()
	image1bytes, _, err := image1.ExportWebp(exportParams)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("output.webp", image1bytes, 0644)
	if err != nil {
		panic(err)
	}
}
