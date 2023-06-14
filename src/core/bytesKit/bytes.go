package bytesKit

import "bytes"

// Compare 按照字典序比较两个字节切片的大小.
/*
@param a 可以为nil
@param b 可以为nil
@return -1:	a < b
		0:	a == b
		1:	a > b
*/
var Compare func(a, b []byte) int = bytes.Compare

// Equals
/*
@param a 可以为nil
@param b 可以为nil

e.g.
	var a = []byte("abcd")
	var b = []byte("abcd")
	println(bytesKit.Equals(a, b)) // true
	b = []byte("abcd1")
	println(bytesKit.Equals(a, b))     // false
	println(bytesKit.Equals(a, nil))   // false
	println(bytesKit.Equals(nil, nil)) // true
*/
func Equals(a, b []byte) bool {
	return bytes.Compare(a, b) == 0
}

// Contains 判断 subslice 子字节切片是否包含在 b 字节切片
var Contains = bytes.Contains

// Count 统计 sep 字节切片在 s 字节切片中非重叠实例数
var Count = bytes.Count

// HasPrefix 字节切片 s 是否以 prefix 开头
var HasPrefix = bytes.HasPrefix

// HasSuffix 字节切片 s 是否以 suffix 结尾
var HasSuffix = bytes.HasSuffix

// Index 查找 sep 在 s 中第一次出现的索引下标，如果没有则返回 -1
var Index = bytes.Index

// LastIndex 查找 sep 在 s 中最后一次出现的索引下标，如果没有则返回 -1
var LastIndex = bytes.LastIndex

// Split 将 sep 作为分割符，将 s 分割，返回拆分之后的字节切片
var Split = bytes.Split

// SplitN 将 sep 作为分割符，将 s 分割 n 份，返回拆分之后的字节切片
var SplitN = bytes.SplitN

// ToLower 将字节切片所有字节全部转换为小写字母，返回该字节切片的一个副本
var ToLower = bytes.ToLower

// ToUpper 将字节切片所有字节全部转换为大小字母，返回该字节切片的一个副本
var ToUpper = bytes.ToUpper

// Trim 返回清除 s 中开头和结尾存在的 cutset 之后的一个子切片
var Trim = bytes.Trim

// TrimLeft 返回清除 s 中开头存在的 cutset 之后的一个子切片
var TrimLeft = bytes.TrimLeft

// TrimRight 返回清除 s 中结尾存在的 cutset 之后的一个子切片
var TrimRight = bytes.TrimRight

// TrimSpace 返回清除 s 中开头和结尾存在的 \t\n\r 之后的一个子切片
var TrimSpace = bytes.TrimSpace
