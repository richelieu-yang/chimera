## Carbon
github(4k Star):
    https://github.com/golang-module/carbon/blob/master/README.cn.md
「Go工具箱」推荐一个轻量级、语义化的时间处理库：carbon
    https://mp.weixin.qq.com/s/92O1SHs4tw1FMcDeQoSfkA

这是一个轻量级的、易于使用的、语义智能的日期时间库，适用于Go开发者.
包括: 
* 开始时间、结束时间
* 旅行时间（时间加减）
* 时间差
* 时间判断
* 时间设置
* 时间获取
* 时间输出（格式化）
* 星座
* 季节
* 农历（目前仅支持 公元1900年 至 2100年 的200年时间跨度）
* !!!: JSON支持 https://github.com/golang-module/carbon/blob/master/README.cn.md#json-%E6%94%AF%E6%8C%81
* i18n（国际化）

#### Carbon 和 time.Time 互转
// (1) 将 time.Time 转换成 Carbon
carbon.CreateFromStdTime(time.Now())
// (2) 将 Carbon 转换成 time.Time
carbon.Now().ToStdTime()

#### Carbon.ToDateTimeString()
输出 "2006-01-02 15:04:05" 格式字符串.
