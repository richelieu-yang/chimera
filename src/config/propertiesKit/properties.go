package propertiesKit

import "github.com/magiconair/properties"

var (
	Load func(buf []byte, enc properties.Encoding) (*properties.Properties, error) = properties.Load

	LoadString func(s string) (*properties.Properties, error) = properties.LoadString

	LoadFile  func(filename string, enc properties.Encoding) (*properties.Properties, error)                        = properties.LoadFile
	LoadFiles func(filenames []string, enc properties.Encoding, ignoreMissing bool) (*properties.Properties, error) = properties.LoadFiles
)
