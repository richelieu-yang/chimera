## 参考
lancet 
    https://www.golancet.cn/api/packages/concurrency.html#Concurrency

## Channel的方法s
Bridge      将多个channel链接到一个channel，直到取消上下文。
FanIn       将多个channel合并为一个channel，直到取消上下文。
Generate    根据传入的值，生成channel.
Or          将一个或多个channel读取到一个channel中，当任何读取channel关闭时将结束读取。
OrDone      将一个channel读入另一个channel，直到取消上下文。
Repeat      返回一个channel，将参数`values`重复放入channel，直到取消上下文。
RepeatFn    返回一个channel，重复执行函数fn，并将结果放入返回的channel，直到取消上下文。
Take        返回一个channel，其值从另一个channel获取，直到取消上下文。
Tee         将一个channel分成两个channel，直到取消上下文。
