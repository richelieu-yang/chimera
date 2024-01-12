package zstdKit

import (
	"github.com/klauspost/compress/zstd"
)

func Compress(data []byte) ([]byte, error) {
	writer, err := zstd.NewWriter(nil)
	if err != nil {
		return nil, err
	}
	defer writer.Close()

	compressed := writer.EncodeAll(data, nil)
	return compressed, nil
}

func Decompress(compressed []byte) ([]byte, error) {
	reader, err := zstd.NewReader(nil)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	return reader.DecodeAll(compressed, nil)
}
