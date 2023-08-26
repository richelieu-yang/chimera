//go:build !amd64 || !go1.16

package jsonKit

import jsoniter "github.com/json-iterator/go"

func init() {
	library = "json-iterator/go"
	defaultAPI = jsoniter.ConfigDefault
}
