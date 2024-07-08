#!/bin/bash

echo "building..."
################### build api
mkdir -p /output/apis/etc
mkdir -p /output/services/etc
pushd api || exit
for i in *;do
  echo "building $i"
  name="$i"
  capName="${name^}"
  cd "$i" || exit
  go build -ldflags="-s -w" -o "/output/apis/${capName}Api"
  cp ./etc/* /output/apis/etc/
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
  cp ./etc/* /output/services/etc/
  cd ..
done
echo "build services done"
echo "build done"