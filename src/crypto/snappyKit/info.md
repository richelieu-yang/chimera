## 参考
「Go开源包」snappy：google开源的快速、无损压缩包
    https://mp.weixin.qq.com/s/ie7LqXZQOUX0Bfn4QhHzLA

## 简介
该包的目标并不是最大化的压缩比例，也不是和其他压缩库兼容；相反，snappy算法的目标是在合理的压缩率下尽可能的提高压缩速度。

### 优点
压缩速度.
e.g.
    与zlib的最快压缩模式相比，snappy依然比其快了一个数量级，但产生的压缩文件要比zip的大20%到100%。

