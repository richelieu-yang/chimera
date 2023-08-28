package reqKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"io"
)

func DownloadToFile(url, filePath string) error {
	if err := fileKit.AssertNotExistOrIsFile(filePath); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(filePath); err != nil {
		return err
	}

	client := NewClient()
	_, err := client.R().SetOutputFile(filePath).Get(url)
	if err != nil {
		return err
	}
	return nil
}

func DownloadToWriter(path string, writer io.Writer) {

}
