package uniofficeKit

import (
	"github.com/unidoc/unioffice/presentation"
)

var (
	NewPresentation func() *presentation.Presentation = presentation.New
)
