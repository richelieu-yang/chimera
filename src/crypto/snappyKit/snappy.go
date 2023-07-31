package snappyKit

import "github.com/golang/snappy"

// Encode 编码.
func Encode(src []byte) []byte {
	var dst []byte
	dst = snappy.Encode(dst, src)
	return dst
}

// Decode 解码.
func Decode(src []byte) ([]byte, error) {
	//length, err := snappy.DecodedLen(src)
	//if err != nil {
	//	return nil, err
	//}

	var dst []byte
	dst, err := snappy.Decode(dst, src)
	if err != nil {
		return nil, err
	}
	return dst, nil
}
