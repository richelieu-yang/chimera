package reqKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
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
	_, err := client.R().SetOutputFile(filePath).Get(url)
	if err != nil {
		return err
	}
	//defer resp.Body.Close()
	return nil
}

func DownloadToWriter(url string, writer io.Writer) error {
	// url
	if err := httpKit.AssertHttpUrl(url); err != nil {
		return err
	}
	// writer
	if err := interfaceKit.AssertNotNil(writer, "writer"); err != nil {
		return err
	}

	client := NewClient()
	_, err := client.R().SetOutput(writer).Get(url)
	if err != nil {
		return err
	}
	//defer resp.Body.Close()
	return nil
}
