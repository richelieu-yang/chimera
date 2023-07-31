## 依赖
github.com/klauspost/compress 而非 github.com/golang/snappy（库太老了）

## 参考
klauspost/compress
    https://github.com/klauspost/compress
「Go开源包」snappy：google开源的快速、无损压缩包
    https://mp.weixin.qq.com/s/ie7LqXZQOUX0Bfn4QhHzLA

## 简介
snappy: Google开源的快速、无损的压缩包.
该包的目标并不是最大化的压缩比例，也不是和其他压缩库兼容；相反，snappy算法的目标是在合理的压缩率下尽可能的提高压缩速度。

### 优点
压缩速度.
e.g.
    与zlib的最快压缩模式相比，snappy依然比其快了一个数量级，但产生的压缩文件要比zip的大20%到100%。

## problem: 压缩后的数据比压缩前大
有时候你会发现，压缩后会比压缩前字节数变大。这是和原字符串有关系，
如果原字符串中重复的字符越少，那么压缩后的长度就有可能会比之前变长；
如果原字符串中重复的字符比较多，那么压缩比率就会很高。这也是压缩的基本原理.
