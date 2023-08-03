## Linux服务器的推荐配置（要永久修改而非临时修改!!!）
ulimit -c: unlimited
ulimit -u: 服务器支持的最大值（通过 ulimit -Hu 命令查看）
ulimit -n: 服务器支持的最大值（通过 ulimit -Hn 命令查看）

kernel.pid_max:     (1) 32位系统: 32768
                    (2) 64位系统: 2000000（200W）
kernel.threads-max: 352656（参考yozo的服务器）
vm.max_map_count:   655360（参考yozo的服务器）
