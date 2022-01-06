# klv
app test for klv;

# run project
``$ go mod vendor``
``$ go run cmd/app/main.go --PORT 8080``
# run tests coverage
``$ go test -cover -coverprofile=c.out ./...``

``$ go tool cover -html=c.out -o coverage.html``

# run tests benchmark
``$ go test -bench=. ./...``

# run doc
``$ go get golang.org/x/tools/cmd/godoc``
``$ godoc -play -http=:6060``

# run in heroku
https://db-postgre-dev.herokuapp.com/