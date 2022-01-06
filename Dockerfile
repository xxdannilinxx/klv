FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main ./app/main.go

EXPOSE 8080

CMD ["/app/main"]