package propertiesKit

import (
	"github.com/magiconair/properties"
	"github.com/mitchellh/mapstructure"
)

var (
	Load func(buf []byte, enc properties.Encoding) (*properties.Properties, error) = properties.Load

	LoadString func(s string) (*properties.Properties, error) = properties.LoadString

	LoadFile  func(filename string, enc properties.Encoding) (*properties.Properties, error)                        = properties.LoadFile
	LoadFiles func(filenames []string, enc properties.Encoding, ignoreMissing bool) (*properties.Properties, error) = properties.LoadFiles
)

func Unmarshal(content []byte, ptr interface{}) error {
	return unmarshal(content, properties.UTF8, ptr)
}

func UnmarshalWithIso88591(content []byte, ptr interface{}) error {
	return unmarshal(content, properties.ISO_8859_1, ptr)
}

// unmarshal
/*
@param ptr 可以是:	var obj interface{}
					unmarshal(content, enc, &obj)
*/
func unmarshal(content []byte, enc properties.Encoding, ptr interface{}) error {
	p, err := Load(content, enc)
	if err != nil {
		return err
	}

	return mapstructure.Decode(p.Map(), ptr)
}
