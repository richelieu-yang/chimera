package i18nPackKit

import "github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"

func Seal(langs []string, code string, data interface{}, msgArgs ...interface{}) (string, error) {
	bean := Pack(langs, code, data, msgArgs...)
	return jsonKit.MarshalToString(bean)
}

func SealFully(langs []string, code, msg string, data interface{}, msgArgs ...interface{}) (string, error) {
	bean := PackFully(langs, code, msg, data, msgArgs...)
	return jsonKit.MarshalToString(bean)
}
