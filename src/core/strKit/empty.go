package strKit

import "strings"

// IsEmpty strings.Compare比自建方法“==”的速度要快
/*
PS: 由于此方法被调用的频率较高，因此方法体内禁止调用自己写的方法，全用原生的.
*/
func IsEmpty(str string) bool {
	return strings.Compare(str, "") == 0
}

func IsNotEmpty(str string) bool {
	return strings.Compare(str, "") != 0
}

func IsEmptyString(obj interface{}) bool {
	if str, ok := obj.(string); ok {
		return IsEmpty(str)
	}
	return false
}

func HasEmpty(strings ...string) bool {
	for _, str := range strings {
		if IsEmpty(str) {
			return true
		}
	}
	return false
}

func IsAllEmpty(strings ...string) bool {
	for _, str := range strings {
		if IsNotEmpty(str) {
			return false
		}
	}
	return true
}

func IsAllNotEmpty(strings ...string) bool {
	return !HasEmpty(strings...)
}

func EmptyToDefault(str, def string) string {
	if IsEmpty(str) {
		return def
	}
	return str
}

func TrimAndEmptyToDefault(str, def string) string {
	str = Trim(str)
	return EmptyToDefault(str, def)
}
