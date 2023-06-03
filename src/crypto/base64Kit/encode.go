package base64Kit

import "github.com/gogf/gf/v2/encoding/gbase64"

func Encode(src []byte) []byte {
	return gbase64.Encode(src)
}

func EncodeFile(path string) ([]byte, error) {
	return gbase64.EncodeFile(path)
}

func EncodeFileToString(path string) (string, error) {
	return gbase64.EncodeFileToString(path)
}

func EncodeString(src string) string {
	return gbase64.EncodeString(src)
}

func EncodeToString(src []byte) string {
	return gbase64.EncodeToString(src)
}
