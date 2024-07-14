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