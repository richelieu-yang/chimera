package randomKit

import "github.com/duke-git/lancet/v2/random"

// LetterString 生成给定长度的随机字符串，只包含字母(a-zA-Z)
var LetterString func(length int) string = random.RandString
