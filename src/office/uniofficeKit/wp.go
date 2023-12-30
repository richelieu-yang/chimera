package uniofficeKit

import "github.com/unidoc/unioffice/document"

var (
	NewDocument func() *document.Document = document.New
)
