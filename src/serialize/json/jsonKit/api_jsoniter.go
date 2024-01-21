//go:build (linux || windows || darwin) && !(sonic && avx && go1.17 && amd64)

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
