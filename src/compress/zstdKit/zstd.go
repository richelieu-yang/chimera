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

//func CompressWriter(writer io.Writer) ([]byte, error) {
//	// 创建一个编码器
//	enc, err := zstd.NewWriter(w)
//	if err != nil {
//		// 处理错误
//	}
//	defer enc.Close()
//
//	// 从一个 io.Reader 中读取数据并压缩到编码器
//	_, err = io.Copy(enc, r)
//	if err != nil {
//		// 处理错误
//	}
//
//}
