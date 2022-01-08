generate:
	@protoc --go_out=. --go-grpc_out=. proto/*.proto

build:
	@go build -o server-bin cmd/server/main.go
	@go build -o client-bin cmd/client/main.go
	@protoc --go_out=plugins=grpc:. *.proto

run_server:
	@PORT=8080 DATABASE_URL="#" go run cmd/server/main.go

run_client:
	@go run cmd/client/main.go

request:
	@grpcurl -plaintext --msg-template -d '{}' localhost:8080 CryptoCurrency.ListCryptoCurrencys