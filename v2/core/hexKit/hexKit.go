// Package hexKit 进制工具
/*
base: 进制
PS: 各种进制的前缀可以参考"Go.docx".
*/
package hexKit

import "strconv"

// ToStringWithBase 值（可以是各种进制，不限于十进制） => 目标进制的字符串
/*
@param i			要转换的值
@param targetBase	目标进制
@return				目标进制的字符串（不带进制的前缀和后缀）
e.g.
(5, 2)		=> 101
(011, 10)	=> 9（8进制转10进制）
(0x11, 10)	=> 17（16进制转10进制）
*/
func ToStringWithBase(i int64, targetBase int) string {
	return strconv.FormatInt(i, targetBase)
}

// ToDecimalism 指定进制的字符串 => 10进制数值
/*
@param s	字符串（不带进制的前缀和后缀）
@param base	字符串的进制
@return		10进制数值
e.g.
("11", 16)	=> 17, nil
("11", 8)	=> 9, nil
*/
func ToDecimalism(s string, base int) (int64, error) {
	return strconv.ParseInt(s, base, 64)
}
