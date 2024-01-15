## 前端解码 
#### atob()解码（后端通过base64.StdEncoding 生成的base64字符串）
PS: js有一个内置的函数atob()，可以将Base64编码的字符串解码为ASCII编码的字符串。但是，如果Base64编码的数据包含非ASCII字符，比如中文，那么atob()函数就不能正确地解码。

总结：前端解码不推荐使用atob()，网上有别的解码方法. 

e.g. atob()无法准确解码（原文中包含中文）
"eyJkYXRhIjoic+WViuWViuWVii5kb2N4IiwiY29kZSI6NH0="

## GoFrame的 gbase64
Deprecated: 它只使用 base64.StdEncoding，太局限了.

BASE64编解码-gbase64
    https://goframe.org/pages/viewpage.action?pageId=1114301

## base64.StdEncoding 和 base64.URLEncoding 的区别
base64.StdEncoding和base64.URLEncoding的区别
    https://blog.csdn.net/harryhare/article/details/88086660
Go语言的JWT鉴权，怎么使用？
	https://mp.weixin.qq.com/s/Ygdqabh6ahu9ZUVsfZ7bzw

相较于 base64.StdEncoding，base64.URLEncoding 会将生成的base64字符串中的 
	'+'（替换为'-'）
	'/'（替换为'_'）

base64.StdEncoding 生成base64字符串可能包含的字符: 		ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/ 和 =（padding）
base64.URLEncoding 生成base64字符串可能包含的字符:		ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_ 和 =（padding）
base64.RawStdEncoding 生成base64字符串可能包含的字符: 	ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/
base64.RawURLEncoding 生成base64字符串可能包含的字符:	ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_

将生成的base64字符串作为 url参数（或文件名的一部分），必需使用 base64.RawURLEncoding，原因: 它生成的base64字符串不会包含以下3种字符: 
	'+'（被替换为'-'）
	'/'（被替换为'_'）
	'='（被省略，因为NoPadding）

e.g.
	input := []byte("\x00\x00\x3e\x00\x00\x3f\x00")
	fmt.Println(string(input))
	fmt.Println(base64.StdEncoding.EncodeToString(input))                               // AAA+AAA/AA==
	fmt.Println(base64.URLEncoding.EncodeToString(input))                               // AAA-AAA_AA==
	fmt.Println(base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(input)) // AAA+AAA/AA
	fmt.Println(base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(input)) // AAA-AAA_AA


