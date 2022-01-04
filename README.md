# klv
app test for klv;

# run project
``$ go run main.go``
# run tests coverage
``$ go test -cover -coverprofile=tests/coverage/c.out ./...``

``$ go tool cover -html=tests/coverage/c.out -o tests/coverage/coverage.html``

# run tests benchmark
``$ go test -bench=. ./...``
# run doc
``$ go get golang.org/x/tools/cmd/godoc``
``$ godoc -play -http=:6060``