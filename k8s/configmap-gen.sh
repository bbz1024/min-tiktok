#!/bin/bash

#kubectl create configmap my-config-from-file --from-file=config.properties
# 是否在min-tiktok目录下
if [ ! "$(basename "$PWD")" = "min-tiktok" ]; then
     echo "please run this script in min-tiktok directory"
  exit 1

fi

#gen-api-configmap
pushd ./api || exit
for i in *;do
  echo "gen-api-configmap $i"
  cd "$i" || exit
  kubectl create configmap "$i-api" --from-file=./etc -n tiktok
  cd ..
done

popd || exit
#gen-services-configmap
pushd ./services || exit
for i in *;do
  echo "gen-services-configmap $i"
  cd "$i" || exit
  kubectl create configmap "$i-service" --from-file=./etc -n tiktok
  cd ..
done
popd || exit