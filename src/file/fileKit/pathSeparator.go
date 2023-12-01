package fileKit

import "os"

const (
	PathSeparatorRune rune = os.PathSeparator

	// PathSeparator 路径分隔符，Mac("/")
	PathSeparator = string(os.PathSeparator)

	PathListSeparatorRune rune = os.PathListSeparator

	// PathListSeparator 路径列表分隔符，Mac(":")
	PathListSeparator = string(os.PathListSeparator)
)
