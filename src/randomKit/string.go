package randomKit

import "github.com/duke-git/lancet/v2/random"

var (
	// LetterString 生成给定长度的随机字符串，只包含字母(a-zA-Z).
	/*
		e.g.
			fmt.Println(randomKit.LetterString(20))	// lheIFnhIxSiBteKONNii
	*/
	LetterString func(length int) string = random.RandString

	// NumberString 生成给定长度的随机数字字符串（0-9）.
	NumberString func(length int) string = random.RandNumeral

	// LowerString 生成给定长度的随机小写字母字符串（a-z）.
	LowerString func(length int) string = random.RandLower

	// UpperString 生成给定长度的随机大写字母字符串（A-Z）.
	UpperString func(length int) string = random.RandUpper

	// NumberOrLetterString 生成给定长度的随机字符串（数字 + 字母; A-Z + a-z + 0-9).
	NumberOrLetterString func(length int) string = random.RandNumeralOrLetter
)
