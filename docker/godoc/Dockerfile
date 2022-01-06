FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go get golang.org/x/tools/cmd/godoc

EXPOSE 6060
CMD ["godoc", "-http", ":6060"]