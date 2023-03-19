package floatKit

import (
	"github.com/richelieu42/chimera/src/core/strKit"
	"github.com/spf13/cast"
)

func ToFloat64(obj interface{}) float64 {
	return ToFloat64WithDefault(obj, 0)
}

func ToFloat64WithDefault(obj interface{}, def float64) float64 {
	i, _ := ToFloat64WithDefaultE(obj, def)
	return i
}

func ToFloat64WithDefaultE(obj interface{}, def float64) (float64, error) {
	switch obj.(type) {
	case string:
		return stringToFloat64(obj.(string), def)
	default:
		if obj == nil {
			return def, nil
		}
		rst, err := cast.ToFloat64E(obj)
		if err != nil {
			return def, err
		}
		return rst, nil
	}
}

func stringToFloat64(str string, def float64) (float64, error) {
	str = strKit.Trim(str)

	if strKit.IsEmpty(str) {
		return def, nil
	}
	rst, err := cast.ToFloat64E(str)
	if err != nil {
		return def, err
	}
	return rst, nil
}
