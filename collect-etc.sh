#!/bin/bash

# 收集配置
mkdir -p /data/etc  # 使用 -p 参数确保即使目录存在也不会出错
if [ ! -d "/data/etc" ]; then
    echo "Error: Failed to create directory /data/etc"
    exit 1
fi

pushd ./api || exit
for i in *; do
    echo "collecting $i"
    cd "$i" || exit
    cp -r ./etc/* /data/etc/ || { echo "Failed to copy files for $i"; exit 1; }
    sed -i 's/info/error/g' /data/etc/
    cd ..
done
popd || exit

pushd ./services || exit
for i in *; do
    echo "collecting $i"
    cd "$i" || exit
    cp -r ./etc/* /data/etc/ || { echo "Failed to copy files for $i"; exit 1; }
    cd ..
done
popd || exit

echo "collect done"

# scp
echo "scp"
ping -c 1 node2 || exit 1
scp -r /data/etc root@node2:/data/etc || { echo "Failed to scp files"; exit 1; }
echo "scp done"