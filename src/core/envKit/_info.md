## 参考
《Golang.wps》
github(6.9k Star):
    https://github.com/joho/godotenv
Go 每日一库之 godotenv
    https://mp.weixin.qq.com/s/595TIIlbhQlhSEkEkUILTw

## .env文件的文件名
文件名任意，甚至不必以 .env 为后缀（但是一般以.env为后缀，以便于识别）.

## .env文件的内容格式
* 普通格式
e.g.
a: 0
b: 1

* YAML格式（只是最基本的YAML格式，不支持嵌套）
e.g.
c = 2
d = 3
