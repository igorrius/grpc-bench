FROM golang:1-alpine AS build

# Install tools required to build the project
# We need to run `docker build --no-cache .` to update those dependencies
RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep

# Gopkg.toml and Gopkg.lock lists project dependencies
# These layers are only re-built when Gopkg files are updated
COPY Gopkg.lock Gopkg.toml /go/src/grpc-bench/
WORKDIR /go/src/grpc-bench/
# Install library dependencies
RUN dep ensure -vendor-only

# Copy all project and build it
# This layer is rebuilt when ever a file has changed in the project directory
COPY . /go/src/grpc-bench/
RUN go build -o /bin/server
