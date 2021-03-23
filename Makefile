bench:
	go test -bench=. -v ./...

server:
	go run cmd/server/main.go

http:
	go run cmd/client/http/client.go

grpc:
	go run cmd/client/grpc/client.go