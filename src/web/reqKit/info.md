## imroc/req 参考（本包使用的是 imroc/req）
GO语言最好用的HTTP请求库
    https://www.bilibili.com/video/BV1Mu4y197dA/
github(3.5k Star)
    https://github.com/imroc/req
快速开始（中文）
    https://req.cool/zh/docs/prologue/quickstart/

#### Close
不需要手动调用 resp.Body.Close()（无论是 普通GET、POST请求 还是 下载请求）.

#### 请求超时（2选1; 推荐使用第1种）
(1) 调用 Client.SetTimeout() 设置超时时间（imroc/req默认: 2 * time.Minute）
(2) 调用 Request.Do() 时传参（context.Context类型）

## go-resty/resty 参考
github(8.2k Star)
    https://github.com/go-resty/resty
