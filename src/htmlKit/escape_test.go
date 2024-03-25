package htmlKit

import (
	"fmt"
	"testing"
)

func TestEscapeAndUnescape(t *testing.T) {
	htmlStr := "<html>Hello.</html>"
	escaped := EscapeString(htmlStr)
	fmt.Println(escaped) // &lt;html&gt;Hello.&lt;/html&gt;
	htmlStr1 := UnescapeString(escaped)
	fmt.Println(htmlStr1) // <html>Hello.</html>
}
