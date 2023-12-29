package gzipKit

import "github.com/zeromicro/go-zero/core/codec"

var (
	// Compress 压缩.
	Compress func(bs []byte) []byte = codec.Gzip

	// Uncompress 解压缩.
	/*
	   PS: 大小限制: 100MB.
	*/
	Uncompress func(bs []byte) ([]byte, error) = codec.Gunzip
)
