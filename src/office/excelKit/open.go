package excelKit

import (
	"github.com/xuri/excelize/v2"
	"io"
)

var OpenFile func(filename string, opts ...excelize.Options) (*excelize.File, error) = excelize.OpenFile

var OpenReader func(r io.Reader, opts ...excelize.Options) (*excelize.File, error) = excelize.OpenReader
