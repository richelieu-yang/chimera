package jsoniterKit

import (
	jsoniter "github.com/json-iterator/go"
)

var (
	// Unmarshal 反序列化.
	/*
	   @param ptr 	(1) 不能为nil
	   			(2) 指针类型
	   @param data	必要条件: len(data) > 0（包含: 不能为nil）
	*/
	Unmarshal = jsoniter.Unmarshal

	// UnmarshalFromString 反序列化.
	/*
	   @param ptr 	(1) 不能为nil
	   			(2) 指针类型
	   @param str	不能为空字符串("")
	*/
	UnmarshalFromString = jsoniter.UnmarshalFromString
)
