FROM golang

WORKDIR $GOPATH/src/app/

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD ["go", "run", "server/main.go"]