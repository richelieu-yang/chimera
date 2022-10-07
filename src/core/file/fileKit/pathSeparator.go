package fileKit

import "os"

const (
	// PathSeparator 路径分隔符，Mac("/")
	PathSeparator          = string(os.PathSeparator)
	PathSeparatorRune rune = os.PathSeparator

	// PathListSeparator 路径列表分隔符，Mac(":")
	PathListSeparator          = string(os.PathListSeparator)
	PathListSeparatorRune rune = os.PathListSeparator
)
