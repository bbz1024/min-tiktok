#!/usr/bin/env bash
# start build
echo  "start build . . ."

# build app
echo "build app . . ."
#docker build -t min-tiktok .
docker save -o min-tiktok.tar min-tiktok

scp min-tiktok.tar root@node2:/root/min-tiktok.tar
ssh root@node2 "docker load -i /root/min-tiktok.tar"

# collect conf
