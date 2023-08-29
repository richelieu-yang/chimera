//go:build !amd64 || !go1.16 || !avx

package jsonKit

import jsoniter "github.com/json-iterator/go"

func init() {
	library = "json-iterator/go"
	defaultApi = jsoniter.ConfigDefault
}
