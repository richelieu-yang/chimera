package base64Kit

import (
	"github.com/richelieu-yang/chimera/v2/src/file/fileKit"
)

func Encode(src []byte, options ...Base64Option) []byte {
	opts := loadOptions(options...)
	return opts.Encode(src)
}

func EncodeToString(src []byte, options ...Base64Option) string {
	opts := loadOptions(options...)
	return opts.EncodeToString(src)
}

// EncodeFile file => []byte
func EncodeFile(path string, options ...Base64Option) ([]byte, error) {
	data, err := fileKit.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return Encode(data, options...), nil
}

// EncodeFileToString file => string
func EncodeFileToString(path string, options ...Base64Option) (string, error) {
	data, err := fileKit.ReadFile(path)
	if err != nil {
		return "", err
	}

	return EncodeToString(data, options...), nil
}
