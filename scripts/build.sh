#!/bin/bash
echo "building..."
################### build api
mkdir -p /output/apis
mkdir -p /output/services
mkdir -p /output/etc
pushd api || exit
for i in *;do
  echo "building $i"
  name="$i"
  capName="${name^}"
  cd "$i" || exit
  go build -ldflags="-s -w" -o "/output/apis/${capName}Api"
  /build/upx -9 "/output/apis/${capName}Api"
  cp ./etc/* /output/etc/
  cd ..
done
echo "build api done"

echo "building services"

popd || exit
pushd services || exit
for i in *;do
  echo "building $i"
  name="$i"
  capName="${name^}"
  cd "$i" || exit
  go build -ldflags="-s -w" -o "/output/services/${capName}Service"
  /build/upx -9 "/output/services/${capName}Service"
  cp ./etc/* /output/etc/
  cd ..
done
echo "build services done"
echo "build done"