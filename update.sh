#!/usr/bin/env bash

# （手动！！！）更新ip库
# src:  https://github.com/lionsoul2014/ip2region/blob/master/data/ip2region.xdb
# dest: resource/ip2region/ip2region.xdb

# （自动）打包内部资源
go-bindata -fs -o=src/resources/resources.go -pkg=resources resources/...

# （自动）更新本项目的第三方依赖
go-mod-upgrade
