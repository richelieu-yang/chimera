package hexKit

import "encoding/hex"

func Encode(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}

func EncodeToString(src []byte) string {
	return hex.EncodeToString(src)
}

func Decode(src []byte) ([]byte, error) {
	n, err := hex.Decode(src, src)
	return src[:n], err
}

func DecodeString(s string) ([]byte, error) {
	return hex.DecodeString(s)
}
