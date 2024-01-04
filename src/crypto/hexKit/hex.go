package hexKit

import "encoding/hex"

// Encode []byte => []byte
/*
参考: hex.EncodeToString
*/
func Encode(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}

// EncodeToString []byte => string
func EncodeToString(src []byte) string {
	return hex.EncodeToString(src)
}

// Decode []byte => []byte
/*
参考: hex.DecodeString.
*/
func Decode(src []byte) ([]byte, error) {
	n, err := hex.Decode(src, src)
	return src[:n], err
}

// DecodeString string => []byte
func DecodeString(s string) ([]byte, error) {
	return hex.DecodeString(s)
}
