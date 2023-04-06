package floatKit

import (
	"github.com/richelieu42/chimera/v2/src/core/strKit"
	"github.com/spf13/cast"
)

func ToFloat32(obj interface{}) float32 {
	return ToFloat32WithDefault(obj, 0)
}

func ToFloat32WithDefault(obj interface{}, def float32) float32 {
	i, _ := ToFloat32WithDefaultE(obj, def)
	return i
}

func ToFloat32WithDefaultE(obj interface{}, def float32) (float32, error) {
	switch obj.(type) {
	case string:
		return stringToFloat32(obj.(string), def)
	default:
		if obj == nil {
			return def, nil
		}
		rst, err := cast.ToFloat32E(obj)
		if err != nil {
			return def, err
		}
		return rst, nil
	}
}

func stringToFloat32(str string, def float32) (float32, error) {
	str = strKit.Trim(str)

	if strKit.IsEmpty(str) {
		return def, nil
	}
	rst, err := cast.ToFloat32E(str)
	if err != nil {
		return def, err
	}
	return rst, nil
}
