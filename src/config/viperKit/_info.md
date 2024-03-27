## viper(v1.17.0) 支持读取的配置文件格式

#### !!!: 反序列化为 map 实例
可能需要加上 mapstructure tag.
e.g. MiddlewareConfig.ResponseHeaders

#### 反序列化，key中的大写会被转换为小写
Viper库在处理配置文件时，会默认将所有的键转换为小写。
这是因为在Viper的设计中，键名是不区分大小写的。这意味着，无论你的配置文件中的键是大写还是小写，Viper在处理时都会将其视为小写。

目前，Viper库并没有提供直接的选项或方法来改变这一行为。也就是说，你无法直接通过某个设置让Viper在处理配置文件时保留键名的原始大小写。

#### 支持的配置文件类型 
viper.go
var SupportedExts = []string{"json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "tfvars", "dotenv", "env", "ini"}

#### TODO: 默认值
Viper不支持通过结构体标签（tag）来设置默认值.  
![_img.png](_img.png)  

e.g. 反例，设置默认值失败
type config struct {  
    A int `json:"a,default=1"`  
    B int `json:"b,default=2"`  
    C int `json:"c"`  
}
