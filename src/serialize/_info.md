## 参考
go语言序列化json/gob/msgp/protobuf性能对比
    https://zhuanlan.zhihu.com/p/409435090
go语言序列化json/gob/msgp/protobuf性能对比
    https://www.cnblogs.com/zhangchaoyang/p/15256978.html

## gob（Richelieu: 不涉及跨语言，推荐使用 gob 而非 json.）
json是一种通用的文本序列化格式，它支持字符串、数字、布尔值、数组、对象等基本类型。json的优点是支持跨语言，易于阅读和调试，广泛应用于web开发和API交互。
缺点: 速度慢，编码复杂，需要使用反射机制，而且不支持一些go语言的特殊类型，如接口、通道、复数等。

## json
gob是go语言特有的二进制序列化格式，它支持所有go语言的基本类型和复合类型，如结构体、切片、映射等。gob的优点是速度快，编码简单，不需要额外的标签或注释。
缺点: 不支持跨语言，只能在go语言之间使用，而且不方便人类阅读。


