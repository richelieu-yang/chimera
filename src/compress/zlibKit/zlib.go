package zlibKit

import "github.com/gogf/gf/v2/encoding/gcompress"

func Zlib(data []byte) ([]byte, error) {
	return gcompress.Zlib(data)
}

func UnZlib(data []byte) ([]byte, error) {
	return gcompress.UnZlib(data)
}
