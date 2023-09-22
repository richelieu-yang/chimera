package excelKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/xuri/excelize/v2"
	"io"
)

var NewFile func(opts ...excelize.Options) *excelize.File = excelize.NewFile

var OpenFile func(filename string, opts ...excelize.Options) (*excelize.File, error) = excelize.OpenFile

var OpenReader func(r io.Reader, opts ...excelize.Options) (*excelize.File, error) = excelize.OpenReader

// NewEmptyFile 新建1个空白的Excel文件（其内只有1个空白的工作表，表名为"Sheet1"）
/*
@param path 文件的路径（如果文件已经存在，会覆盖它）
*/
func NewEmptyFile(path string, opts ...excelize.Options) (*excelize.File, error) {
	if err := fileKit.MkParentDirs(path); err != nil {
		return nil, err
	}

	f := excelize.NewFile()
	err := f.SaveAs(path, opts...)
	if err != nil {
		defer func(f *excelize.File) {
			_ = f.Close()
		}(f)
		return nil, err
	}
	return f, nil
}
