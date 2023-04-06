package idKit

import gonanoid "github.com/matoous/go-nanoid/v2"

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
