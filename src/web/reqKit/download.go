package reqKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/web/httpKit"
	"io"
)

func DownloadToFile(url, filePath string) error {
	// url
	if err := httpKit.AssertHttpUrl(url); err != nil {
		return err
	}
	// filePath
	if err := fileKit.AssertNotExistOrIsFile(filePath); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(filePath); err != nil {
		return err
	}

	client := NewClient()
	resp, err := client.R().SetOutputFile(filePath).Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func DownloadToWriter(path string, writer io.Writer) {

}
