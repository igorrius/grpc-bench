#!/usr/bin/env bash

export GOPATH="$(dirname $(dirname $(dirname $(realpath  $0))))"
echo "Setup GOPATH to $GOPATH"

# create project tree
echo "Creating project tree..."
mkdir ${GOPATH}/bin
mkdir ${GOPATH}/pkg

echo "Download and install Protoc"
TMP_DIR=__tmp
# PROTOC_DOWNLOAD_LINK is the link to protobuf binary.
# You can get it from https://github.com/google/protobuf/releases
PROTOC_DOWNLOAD_LINK=https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
mkdir ${TMP_DIR}
wget ${PROTOC_DOWNLOAD_LINK} -O ${TMP_DIR}/protoc.zip
(cd ${TMP_DIR} && unzip -o protoc.zip && cp bin/protoc ${GOPATH}/bin)
rm -rf ${TMP_DIR}

# install dependency packages
echo "Install dependency packages..."
go get -u github.com/golang/dep/cmd/dep
${GOPATH}/bin/dep ensure -v -vendor-only
go get -v -u github.com/golang/protobuf/protoc-gen-go

