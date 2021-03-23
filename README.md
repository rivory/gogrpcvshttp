# Benchmark : GRPC vs HTTP protocols

We wanted to compare http protocol over the Grpc one.

This project contains a server which implements both http and grpc protocol handler, which operate the same business logic and return the same answer to clients. 
```
make server
```


This project contains two distinct clients, one to interact via HTTP and one via GRPC
```
make http
make grpc
```


It also includes a `benchmark_test.go` which benchmark the same usage case for both protocol and print metrics 
```
make bench
```


Few run from local env : 

```
grpcVsHttp/cmd/server via 🐹 v1.16.2 
❯ go test -bench=.
Starting http server at port 8080
goos: darwin
goarch: amd64
pkg: github.com/rivory/gogrpcvshttp/cmd/server
cpu: Intel(R) Core(TM) i7-8559U CPU @ 2.70GHz
BenchmarkHTTP-8             4558            267774 ns/op
BenchmarkGrpc-8            20925             48192 ns/op
PASS
ok      github.com/rivory/gogrpcvshttp/cmd/server       2.980s

❯ go test -bench=.
Starting http server at port 8080
goos: darwin
goarch: amd64
pkg: github.com/rivory/gogrpcvshttp/cmd/server
cpu: Intel(R) Core(TM) i7-8559U CPU @ 2.70GHz
BenchmarkHTTP-8             5167            263932 ns/op
BenchmarkGrpc-8            30001             39344 ns/op
PASS
ok      github.com/rivory/gogrpcvshttp/cmd/server       3.213s

❯ go test -bench=.
Starting http server at port 8080
goos: darwin
goarch: amd64
pkg: github.com/rivory/gogrpcvshttp/cmd/server
cpu: Intel(R) Core(TM) i7-8559U CPU @ 2.70GHz
BenchmarkHTTP-8             4916            314233 ns/op
BenchmarkGrpc-8            30205             39908 ns/op
PASS
ok      github.com/rivory/gogrpcvshttp/cmd/server       4.139s
```