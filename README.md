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