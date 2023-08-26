//go:build !amd64 || !go1.16

package jsonKit

func init() {
	library = "json-iterator/go"
	defaultAPI = jsoniter.ConfigDefault
}
