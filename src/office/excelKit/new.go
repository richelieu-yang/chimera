package excelKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/xuri/excelize/v2"
)

var NewFile func(opts ...excelize.Options) *excelize.File = excelize.NewFile

// NewFileWithPath 新建1个空白的Excel文件（其内只有1个空白的工作表，表名为"Sheet1"）
/*
@param path 文件的路径（如果文件已经存在，会覆盖它）
*/
func NewFileWithPath(path string, opts ...excelize.Options) (f *excelize.File, err error) {
	if err = fileKit.MkParentDirs(path); err != nil {
		return
	}

	f = NewFile()
	defer func(f *excelize.File) {
		if err != nil {
			_ = f.Close()
		}
	}(f)
	if err = f.SaveAs(path, opts...); err != nil {
		return nil, err
	}
	return
}
