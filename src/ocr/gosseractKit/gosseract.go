package gosseractKit

import (
	"github.com/otiai10/gosseract/v2"
	"github.com/richelieu-yang/chimera/v2/src/file/fileKit"
)

// GertText
/*
!!!: 使用此函数，必须确保"CGO_ENABLED=1"，否则go run或go build会报错: undefined: gosseract.NewClient
*/
func GertText(imgPath string) (string, error) {
	if err := fileKit.AssertExistAndIsFile(imgPath); err != nil {
		return "", err
	}

	client := gosseract.NewClient()
	defer client.Close()
	if err := client.SetImage(imgPath); err != nil {
		return "", err
	}
	return client.Text()
}
