FROM golang:1.10.3-alpine3.7 AS build
LABEL pavle=523260513@qq.com

RUN apk add --no-cache git openssh

WORKDIR $GOPATH/src/

COPY .ssh /root/.ssh
RUN chmod 0600 /root/.ssh/id_rsa

RUN mkdir -p $GOPATH/src/github.com/Liv1020/ \
    && cd github.com/Liv1020/ \
    && git clone git@github.com:Liv1020/move-car.git

RUN cd $GOPATH/src/github.com/Liv1020/move-car/ \
    && git checkout master \
    && git pull origin master \
    && CGO_ENABLED=0 go build -a

FROM alpine:latest
LABEL pavle=523260513@qq.com

RUN apk add --no-cache ca-certificates

RUN mkdir -p /usr/local/server
WORKDIR /usr/local/server

VOLUME /usr/local/server/conf
COPY --from=build /go/src/github.com/Liv1020/move-car/move-car .

CMD ["/usr/local/server/move-car"]
