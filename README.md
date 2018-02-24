How to using bench?
===================
1. Create `folder/src` and clone project from GIT to `folder/src`
2. Run `./init-dev.sh`

P.S. for build proto use `./build-proto.sh` command

How to run benchmark?
=====================
1. Start server `docker-compose up`
2. Run command `go test -run=10 -bench=.`

or

2. Run `client/client.go`

Enjoy!   

Collected stats on my home PC
=============================
```
goos: linux
goarch: amd64
pkg: grpc-bench
BenchmarkClientHaProxy-4           10000            147564 ns/op
BenchmarkManyClientHaProxy-4        2000            504657 ns/op
BenchmarkClientNginx-4             10000            142201 ns/op
BenchmarkManyClientNginx-4          3000            485117 ns/op
```