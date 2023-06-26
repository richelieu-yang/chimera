package bytesKit

import "unsafe"

// StringToBytes string => []byte
/*
Go 如何让在强制转换类型时不发生内存拷贝
	https://mp.weixin.qq.com/s/wvJm0Q95PTB_YUkZq51C8Q
PS:
(1) 参考了 gin(v1.9.1) 中的 internal/bytesconv/bytesconv.go ，后续要及时同步更新此方法！
(2) 必须满足的条件：对数据仅仅只有读操作！！！
*/
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// BytesToString []byte => string
/*
Go 如何让在强制转换类型时不发生内存拷贝
	https://mp.weixin.qq.com/s/wvJm0Q95PTB_YUkZq51C8Q
PS:
(1) 参考了 gin(v1.9.1) 中的 internal/bytesconv/bytesconv.go ，后续要及时同步更新此方法！
(2) 必须满足的条件：对数据仅仅只有读操作！！！
*/
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
