package base64Kit

import "github.com/gogf/gf/v2/encoding/gbase64"

func Decode(data []byte) ([]byte, error) {
	return gbase64.Decode(data)
}

func DecodeString(data string) ([]byte, error) {
	return gbase64.DecodeString(data)
}

func DecodeToString(data string) (string, error) {
	return gbase64.DecodeToString(data)
}
