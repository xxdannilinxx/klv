FROM registry.semaphoreci.com/golang:1.15 as builder
ENV APP_USER app
ENV APP_HOME /go/src/klv
RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME
WORKDIR $APP_HOME
USER $APP_USER
COPY src/ .
RUN go mod download
RUN go mod verify
RUN go build -o klv

FROM debian:buster
FROM registry.semaphoreci.com/golang:1.15
ENV APP_USER app
ENV APP_HOME /go/src/klv
RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME
COPY --chown=0:0 --from=builder $APP_HOME/klv $APP_HOME
EXPOSE 8080
USER $APP_USER
CMD ["./klv"]