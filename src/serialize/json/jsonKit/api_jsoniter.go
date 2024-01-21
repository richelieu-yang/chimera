//go:build !(sonic && avx && go1.16 && amd64 && (linux || windows || darwin))

package jsonKit

import (
	jsoniter "github.com/json-iterator/go"
)

func init() {
	library = "json-iterator/go"
	defaultApi = jsoniter.ConfigDefault
	stdApi = jsoniter.ConfigCompatibleWithStandardLibrary

	testAPI()
}
