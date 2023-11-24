package excelKit

import (
	"github.com/xuri/excelize/v2"
	"io"
)

// OpenFile 打开本地文件.
var OpenFile func(filename string, opts ...excelize.Options) (*excelize.File, error) = excelize.OpenFile

// OpenReader 打开数据流.
/*
@params r 数据流（包括: 远程文件）
*/
var OpenReader func(r io.Reader, opts ...excelize.Options) (*excelize.File, error) = excelize.OpenReader
