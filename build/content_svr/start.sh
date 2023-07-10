#!/bin/bash

#chmod -R +x ./

TARGET_BIN='content_svr'
#grep命令的参数 -v 即反向查找，awk '{print $2}'  打印出第二列参数即所有过滤后进程的pid；
for N in `ps -ef | grep "./bin/$TARGET_BIN" | grep -v grep | awk '{print \$2}'`
do
        STOP="kill -9 $N"
        eval $STOP
done

sleep 1

echo "starting ..."

START="nohup ./bin/$TARGET_BIN > /dev/null 2>&1 &"

eval $START
echo "done."