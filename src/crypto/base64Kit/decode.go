package base64Kit

import "github.com/gogf/gf/v2/encoding/gbase64"

var Decode func(data []byte) ([]byte, error) = gbase64.Decode

var DecodeString func(data string) ([]byte, error) = gbase64.DecodeString

var DecodeToString func(data string) (string, error) = gbase64.DecodeToString
