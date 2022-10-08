#!/bin/bash

# 启动编译环境
docker stop gtp
docker start go
# 编译 bulid/*.linux
docker exec  -w /go/src/gtp/ -i go sh build.sh amd64 linux 1.0
echo "1.编译完成."

# 编译docker文件
sh build.online
echo "2.docker编译完成"

# 运行
sh up.online
echo "启动"