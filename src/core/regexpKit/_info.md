## 参考
Golang - 1.wps

Go语言的一些正则表达式例子，开发中经常用到
    https://mp.weixin.qq.com/s/H8DcvXyoV6mRqfEsCypn8A

## demos
e.g.    中文姓名的正则表达式
regex := "^[\u4e00-\u9fa5]{2,4}$"

e.g.1   匹配连续重复的字符
re := regexp.MustCompile(`(.)\1+`)

e.g.2   定义手机号码运营商的正则表达式
// 中国移动
cmccRegex := `^1(34[0-8]|3[5-9]\d|4[7-8]\d|5[0-27-9]\d|7[28]\d|8[2-478]\d|9[8-9]\d)\d{7}$`
// 中国联通
cuccRegex := `^1(3[0-2]\d|4[56]\d|5[56]\d|6[67]\d|7[1-9]\d|8[56]\d)\d{7}$`
// 中国电信
ctccRegex := `^1(34[9]\d|3[3]\d|53\d|7[037]\d|8[019]\d|9[019]\d)\d{7}$`
