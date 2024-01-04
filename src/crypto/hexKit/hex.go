package hexKit

import "encoding/hex"

// Encode
/*
参考: hex.EncodeToString
*/
func Encode(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}

// Decode
/*
参考: hex.DecodeString.
*/
func Decode(src []byte) ([]byte, error) {
	n, err := hex.Decode(src, src)
	return src[:n], err
}

func EncodeToString(src []byte) string {
	return hex.EncodeToString(src)
}

func DecodeString(s string) ([]byte, error) {
	return hex.DecodeString(s)
}
