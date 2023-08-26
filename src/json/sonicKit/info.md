## 限制!!!
(1) sonic only supports Go1.16~1.20 && CPU amd64
(2) 并非所有 CPU amd64 都支持（e.g.yozo某台内网机）

## Sonic简介
一个速度奇快的 JSON 序列化/反序列化库，由 JIT （即时编译）和 SIMD （单指令流多数据流）加速。

## 文档s
官方文档
    https://github.com/bytedance/sonic/blob/main/README_ZH_CN.md
Go语言JSON解析届顶流：Sonic
    https://mp.weixin.qq.com/s/Ij5wNjNZ6rRbQqTYIvP_aw

## 对map的键排序（性能损失约10%，默认不启用）
(1) 最简单的一种方法: 直接使用 sonic.ConfigStd .
(2) 自行创建 sonic.API实例(SortMapKeys).

## HTML 转义（性能损失约15%，默认不启用）
(1) 最简单的一种方法: 直接使用 sonic.ConfigStd .
(2) 自行创建 sonic.API实例(EscapeHTML).

## problems
#### 多次序列化同一map（无序的），希望结果一致
sonic.ConfigDefault:    多次实验（1000000），发现生成的json字符串一致.
sonic.ConfigStd:        生成的json字符串必定一致.

结论：稳一点，还是使用 sonic.ConfigStd 作为API吧.
