package hexKit

import "encoding/hex"

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

// DecodeStringToString hex string => string
func DecodeStringToString(s string) (string, error) {
	data, err := DecodeString(s)
	return string(data), err
}
