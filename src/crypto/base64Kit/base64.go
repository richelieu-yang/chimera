package base64Kit

import (
	"encoding/base64"
	"github.com/richelieu42/chimera/src/core/sliceKit"
)

// Encode
/*
@param args 可以是 base64.StdEncoding、base64.URLEncoding、base64.RawStdEncoding、base64.RawURLEncoding 等（默认: base64.StdEncoding）
*/
func Encode(data []byte, args ...*base64.Encoding) []byte {
	encoding := sliceKit.GetFirstItemWithDefault(base64.StdEncoding, args...)

	encodedLen := encoding.EncodedLen(len(data))
	out := make([]byte, encodedLen)
	encoding.Encode(out, data)
	return out
}

func EncodeToString(data []byte, args ...*base64.Encoding) string {
	encoding := sliceKit.GetFirstItemWithDefault(base64.StdEncoding, args...)

	return encoding.EncodeToString(data)
}

func Decode(data []byte, args ...*base64.Encoding) ([]byte, error) {
	encoding := sliceKit.GetFirstItemWithDefault(base64.StdEncoding, args...)

	decodedLen := encoding.DecodedLen(len(data))
	out := make([]byte, decodedLen)
	n, err := encoding.Decode(out, data)
	if err != nil {
		return nil, err
	}
	return out[:n], nil
}

func DecodeString(str string, args ...*base64.Encoding) ([]byte, error) {
	encoding := sliceKit.GetFirstItemWithDefault(base64.StdEncoding, args...)

	return encoding.DecodeString(str)
}

func DecodeToString(data []byte, args ...*base64.Encoding) (string, error) {
	out, err := Decode(data, args...)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
