## 参考
缓存管理-内存缓存
    https://goframe.org/pages/viewpage.action?pageId=1114311

## 简介
缓存组件默认提供了一个高速的内存缓存，操作效率非常高效，CPU性能损耗在ns纳秒级别.
* 支持设置过期时间（TTL）
* 不使用时，需要手动调用 Close()
