## GoFrame的 gbase64
Deprecated: 它只使用 base64.StdEncoding，太局限了.

BASE64编解码-gbase64
    https://goframe.org/pages/viewpage.action?pageId=1114301


## base64.StdEncoding 和 base64.URLEncoding 的区别
base64.StdEncoding和base64.URLEncoding的区别
    https://blog.csdn.net/harryhare/article/details/88086660

base64.StdEncoding 生成base64字符串可能包含的字符: 
base64.URLEncoding 生成base64字符串可能包含的字符:
base64.RawStdEncoding 生成base64字符串可能包含的字符:
base64.RawURLEncoding 生成base64字符串可能包含的字符:

e.g.
	input := []byte("\x00\x00\x3e\x00\x00\x3f\x00")
	fmt.Println(string(input))
	fmt.Println(base64.StdEncoding.EncodeToString(input))                               // AAA+AAA/AA==
	fmt.Println(base64.URLEncoding.EncodeToString(input))                               // AAA-AAA_AA==
	fmt.Println(base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(input)) // AAA+AAA/AA
	fmt.Println(base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(input)) // AAA-AAA_AA

