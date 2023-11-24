package excelKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/xuri/excelize/v2"
	"io"
)

// OpenFile 打开本地文件.
/*
PS: 可能会返回 excelize.ErrWorkbookFileFormat，原因: 文件格式不支持.
*/
func OpenFile(path string, opts ...excelize.Options) (*excelize.File, error) {
	if err := fileKit.AssertExistAndIsFile(path); err != nil {
		return nil, err
	}
	return excelize.OpenFile(path, opts...)
}

// OpenReader 打开数据流.
/*
PS: 可能会返回 excelize.ErrWorkbookFileFormat，原因: 文件格式不支持.

@params r 数据流（包括: 远程文件）
*/
var OpenReader func(r io.Reader, opts ...excelize.Options) (*excelize.File, error) = excelize.OpenReader
