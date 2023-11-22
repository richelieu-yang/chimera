package unicodeKit

import (
	"bytes"
	"strconv"
)

var (
	Decode func(s string) (string, error) = strconv.Unquote
)

func Encode(str string) string {
	buf := new(bytes.Buffer)

	s := []rune(str)
	for _, r := range s {
		buf.WriteString(strconv.QuoteRuneToASCII(r))
	}
	return buf.String()
}
