#!/usr/bin/env bash
# start build
echo  "start build . . ."

# build app
echo "build app . . ."
#docker build -t min-tiktok .
#docker save -o min-tiktok.tar min-tiktok
#scp min-tiktok.tar root@node2:/root/min-tiktok.tar
#ssh root@node2 "docker load -i /root/min-tiktok.tar"

# collect conf
echo "collect conf . . ."
bash ./scripts/collect-etc.sh
# nginx conf
cp ./conf/nginx-k8s.conf /data/etc/nginx-k8s.conf
echo "collect conf done"
# scp
echo "scp ..."
#ping -c 1 node2 || exit 1
scp -r /data/etc/* root@node1:/data/etc || { echo "Failed to scp files"; exit 1; }
scp -r /data/etc/* root@node2:/data/etc || { echo "Failed to scp files"; exit 1; }
mkdir -p /project/min-tiktok/k8s
cp ./k8s/* /project/min-tiktok/k8s
scp -r /project/min-tiktok/k8s root@node1:/ || { echo "Failed to scp files"; exit 1; }
scp -r /project/min-tiktok/k8s root@node2:/ || { echo "Failed to scp files"; exit 1; }
echo "scp done ..."

# k8s

#kubectl apply -f ./k8s/nginx.yaml
#kubectl apply -f ./k8s/auths.yaml
#kubectl apply -f ./k8s/user.yaml
