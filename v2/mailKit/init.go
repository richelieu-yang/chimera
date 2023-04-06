package mailKit

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-message/charset"
)

func init() {
	/*
		处理邮件正文，包里已经封装处理好了，包括多字节字符的处理，只需要调用就行了。
		这里需要用到 emersion/go-message 包 设置 imap.CharsetReader 以支持除了 UTF-8 和 ASCII 以外的字符编码，如果不设置则支持 UTF-8 和 ASCII ，像 gb2312、gb18030 这些是无法处理的。
	*/
	imap.CharsetReader = charset.Reader
}
