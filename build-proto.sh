#!/usr/bin/env bash

BUILD_ROOT=$(pwd)/proto

for file in $(find $BUILD_ROOT -name *.proto)
do
    echo "Build proto: $file"
    protoc -I${BUILD_ROOT} \
        --go_out=plugins=grpc:${BUILD_ROOT} \
        ${file}
done
