package uniofficeKit

import (
	"github.com/unidoc/unioffice/spreadsheet"
)

var (
	NewWorkbook func() *spreadsheet.Workbook = spreadsheet.New
)
