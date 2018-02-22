#!/usr/bin/env bash

export GOPATH="$(dirname $(dirname $(dirname $(realpath  $0))))"
echo "Setup GOPATH to $GOPATH"

# create project tree
echo "Creating project tree..."
mkdir ${GOPATH}/bin
mkdir ${GOPATH}/pkg

# install dependency packages
echo "Install dependency packages..."
go get -u github.com/golang/dep/cmd/dep
${GOPATH}/bin/dep ensure -v -vendor-only
