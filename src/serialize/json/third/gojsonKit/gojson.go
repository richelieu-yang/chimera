package gojsonKit

import "github.com/goccy/go-json"

var (
	Marshal func(v interface{}) ([]byte, error) = json.Marshal

	MarshalIndent func(v interface{}, prefix, indent string) ([]byte, error) = json.MarshalIndent

	Unmarshal func(data []byte, v interface{}) error = json.Unmarshal
)
