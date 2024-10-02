#!/usr/bin/env bash

if [ ! "$(basename "$PWD")" ] = "min-tiktok"; then
  echo "please run this script in min-tiktok"
  exit 1
fi

echo "starting..."
# 获取Jenkins构建号（如果在Jenkins环境中）
BUILD_NUMBER=${BUILD_NUMBER:-"unknown"}
BUILD_HASH=${BUILD_NUMBER}-${GIT_COMMIT}
IMAGE_NAME=min-tiktok:${BUILD_HASH}
PUSH_IMAGE_NAME=swr.cn-south-1.myhuaweicloud.com/lzb/${IMAGE_NAME}
PUSH_IMAGE_NAME_LATEST=swr.cn-south-1.myhuaweicloud.com/lzb/min-tiktok:latest


# 构建镜像并使用唯一的标签
docker build -t  "${IMAGE_NAME}" .
echo "build success"


# 推送带有唯一标签的镜像
docker tag "${IMAGE_NAME}"  "${PUSH_IMAGE_NAME}" # 提交带有唯一标签的镜像
docker tag "${IMAGE_NAME}"  "${PUSH_IMAGE_NAME_LATEST}"
docker push "${PUSH_IMAGE_NAME}"
docker push "${PUSH_IMAGE_NAME_LATEST}"

# 清理本地镜像
docker rmi -f "${IMAGE_NAME}"
docker rmi -f "${PUSH_IMAGE_NAME_LATEST}"
docker rmi -f "${PUSH_IMAGE_NAME}"
echo "clear success"


# k8 replace
echo "k8 replace"
# 遍历目录中的所有 YAML 文件
find "./k8s" -name "*.yaml" | while read FILE; do
  sed -i "s|BUILD_HASH|${BUILD_HASH}|" "$FILE"
done
echo "k8 replace done"

echo "collecting etc"
# 收集配置文件
sh ./scripts/collect-etc.sh
# 传输
scp -r /data/etc root@k8s-worker1:/data
scp -r /data/etc root@k8s-worker2:/data
scp -r ./k8s root@k8s-master:/project/min-tiktok
echo "collect done"