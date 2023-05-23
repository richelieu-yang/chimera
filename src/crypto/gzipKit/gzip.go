package gzipKit

import "github.com/zeromicro/go-zero/core/codec"

func Gzip(bs []byte) []byte {
	return codec.Gzip(bs)
}

// Gunzip
/*
PS: 大小限制: 100MB.
*/
func Gunzip(bs []byte) ([]byte, error) {
	return codec.Gunzip(bs)
}
