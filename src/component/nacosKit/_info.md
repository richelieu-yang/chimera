## 参考
"Nacos2.x - Golang.wps"

## 获取配置
* Group字段实际上可以不传，将采用默认值 "DEFAULT_GROUP"
* DataId 和 Group 对应的配置不存在的情况，将返回 ("", nil)
