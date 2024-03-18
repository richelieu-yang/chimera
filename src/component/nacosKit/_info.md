## 参考
- notes/micro（微服务）/注册中心/Nacos（注册中心+配置中心）/Nacos2.x.wps
- notes/micro（微服务）/注册中心/Nacos（注册中心+配置中心）/Nacos2.x - Golang.wps

## naming
#### weight 字段
值范围: [0.0, 10000.0]

e.g. 注册实例时，weight值非法(12345.67)
    会返回error(retry 3 times request failed!: request return error code 500).

## config
#### 获取配置
* 某一namespace，DataId 和 Group 对应的配置不存在的情况，将返回 ("", nil)


