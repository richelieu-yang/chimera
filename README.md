# go-scales

#### 简介 
Golang的工具类.

#### 如何快速搭建一个项目？？？
(0) (optional)对jsonKit进行初始化
通过 jsonKit.SetMsgProcessor()对响应的message进行二次处理.

(1) 初始化日志框架logrus  
logrusKit.Initialize()
logrusKit.PrintBasicInfo()

(2) 读取配置文件，并设置go-scales是否为debug模式 
mainControl.SetDebug(config.Debug)

#### 安装此依赖
PS: 如果安装失败（因为部分依赖 i/o timeout等原因），可以考虑更改 环境变量GOPROXY，比如切换为阿里的代理（https://mirrors.aliyun.com/goproxy/）.
go get github.com/richelieu42/go-scales
