generate:
	@protoc --go_out=. --go-grpc_out=. proto/*.proto

build:
	@go build -o server-bin cmd/server/main.go
	@go build -o client-bin cmd/client/main.go
	@protoc --go_out=plugins=grpc:. *.proto

run_server:
	@go run cmd/server/main.go

run_client:
	@go run cmd/client/main.go

test_cov:
	@go test ./...
	@go test -cover -coverprofile=c.out ./...
	@go tool cover -html=c.out -o coverage.html

test_benchmark:
	@go test -bench=. ./...

test_doc:
	@go get golang.org/x/tools/cmd/godoc
	@godoc -play -http=:6060

grpc_ui:
	@grpcui -plaintext localhost:${PORT}

request:
	@grpcurl -plaintext -d '' localhost:${PORT} CryptoCurrency.GetMostVotedCryptoCurrency