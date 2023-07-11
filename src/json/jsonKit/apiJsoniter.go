//go:build !((linux || windows || darwin) && amd64)

package jsonKit

import jsoniter "github.com/json-iterator/go"

func init() {
	api = jsoniter.ConfigDefault
}
