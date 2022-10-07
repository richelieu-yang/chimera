package mapKit

import (
	"github.com/richelieu42/go-scales/src/core/strKit"
)

// JoinSS map[string]string => string
func JoinSS(m map[string]string, separator, keyValueSeparator string, keyCallback func(str string) string, valueCallback func(str string) string) (rst string) {
	if len(m) == 0 {
		return
	}

	for key, value := range m {
		if keyCallback != nil {
			key = keyCallback(key)
		}
		if valueCallback != nil {
			value = valueCallback(value)
		}

		var tmp string
		if strKit.IsEmpty(value) {
			tmp = key
		} else {
			tmp = key + keyValueSeparator + value
		}

		if strKit.IsEmpty(rst) {
			rst += tmp
		} else {
			rst += separator + tmp
		}
	}
	return
}

// CloneSS
/*
@return may be nil
*/
func CloneSS(m map[string]string) map[string]string {
	if m == nil {
		return nil
	}
	dolly := make(map[string]string)
	for k, v := range m {
		dolly[k] = v
	}
	return dolly
}

// MergeSS 合并，将后面的合并到前面的中
/*
@return may be nil
@return 新的map实例（不会修改传参map）
*/
func MergeSS(m map[string]string, m1 map[string]string) map[string]string {
	if m == nil {
		return CloneSS(m1)
	}
	if m1 == nil {
		return CloneSS(m)
	}

	rst := CloneSS(m)
	for k, v := range m1 {
		rst[k] = v
	}
	return rst
}

// MergeSSInQuantity 大批量地合并
/*
@return may be nil
@return 新的map实例（不会修改传参map）
*/
func MergeSSInQuantity(maps ...map[string]string) map[string]string {
	var rst map[string]string

	for _, m := range maps {
		if rst == nil {
			rst = CloneSS(m)
		} else {
			rst = MergeSS(rst, m)
		}
	}
	return rst
}
