package idKit

/*
uuid
uuid的五个版本: https://www.cnblogs.com/xjnotxj/p/12162733.html

nanoid
NanoID 了解一下？比 UUID 更好用！ - https://mp.weixin.qq.com/s/4muEuUkk3tq6iJXLwspQyQ
gonanoid依赖（650 Star） - https://github.com/matoous/go-nanoid
*/

import (
	"gitee.com/richelieu042/go-scales/src/core/strKit"
	gonanoid "github.com/matoous/go-nanoid/v2"
	uuid "github.com/satori/go.uuid"
)

// NewUUID uuid v4
/*
@return e.g. "fd794d5a-4e7d-456d-a9e4-e377bf00f0a0"
*/
func NewUUID() string {
	return uuid.NewV4().String()
}

// NewSimpleUUID uuid v4
/**
@return e.g. "e28351058d0c446b85e3d7896c87078b"
*/
func NewSimpleUUID() string {
	return strKit.ReplaceAll(NewUUID(), "-", "")
}

// NewNanoId 生成NanoId
/*
NanoID 了解一下？比 UUID 更好用！ - https://mp.weixin.qq.com/s/4muEuUkk3tq6iJXLwspQyQ

@param l	生成NanoId字符串的长度，不传参则默认长度为21（！！！：传参个数只能为0或1，且值>=0）
@return		返回值的字母表为: "_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

e.g.
() => ("4spQ-UdpjbhSjE046w1Ij", nil)
*/
func NewNanoId(l ...int) (string, error) {
	return gonanoid.New(l...)
}

// NewCustomizedNanoId 生成定制化的NanoId
/*
@param alphabet	字母表（默认为"_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"）
@param size		生成NanoId字符串的长度
*/
func NewCustomizedNanoId(alphabet string, size int) (string, error) {
	return gonanoid.Generate(alphabet, size)
}
