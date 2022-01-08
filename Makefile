generate:
	@protoc --go_out=. --go-grpc_out=. proto/*.proto

build:
	@go build -o server-bin cmd/server/main.go
	@go build -o client-bin cmd/client/main.go
	@protoc --go_out=plugins=grpc:. *.proto

run_server:
	@PORT=8090 POSTGRES_USER="admin" POSTGRES_PASSWORD="admin" POSTGRES_DB="klv" POSTGRES_HOST="localhost" POSTGRES_PORT="5432" go run cmd/server/main.go

run_client:
	@go run cmd/client/main.go

test_cov:
	@go test -cover -coverprofile=c.out ./...
	@go tool cover -html=c.out -o coverage.html

test_benchmark:
	@go test -bench=. ./...

test_doc:
	@go get golang.org/x/tools/cmd/godoc
	@godoc -play -http=:6060

request:
	@grpcurl -plaintext --msg-template -d '{}' localhost:8080 CryptoCurrency.ListCryptoCurrencys