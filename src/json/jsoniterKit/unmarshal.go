package jsoniterKit

import (
	jsoniter "github.com/json-iterator/go"
)

// Unmarshal 反序列化.
/*
Description: 建议使用sonicKit.

@param ptr 	(1) 不能为nil
			(2) 指针类型
@param data	必要条件: len(data) > 0（包含: 不能为nil）
*/
func Unmarshal(data []byte, ptr interface{}) error {
	///* 传参检查 */
	//if err := ptrKit.AssertNotNilAndIsPointer(ptr); err != nil {
	//	return err
	//}
	//if err := sliceKit.AssertNotEmpty(data, "data"); err != nil {
	//	return err
	//}

	return jsoniter.Unmarshal(data, ptr)
}

// UnmarshalFromString 反序列化.
/*
Description: 建议使用sonicKit.

@param ptr 	(1) 不能为nil
			(2) 指针类型
@param str	不能为空字符串("")
*/
func UnmarshalFromString(str string, ptr interface{}) error {
	///* 传参检查 */
	//if err := ptrKit.AssertNotNilAndIsPointer(ptr); err != nil {
	//	return err
	//}
	//if err := strKit.AssertNotBlank(str, "str"); err != nil {
	//	return err
	//}

	return jsoniter.UnmarshalFromString(str, ptr)
}
